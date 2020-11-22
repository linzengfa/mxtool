/**********************************************
** @Des: WXPay.go
** @Author: MoXiao
** @Date:   2018/10/16 9:43
** @Last Modified by:  MoXiao
** @Last Modified time: 2018/10/16 9:43
***********************************************/
package wxpay

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//创建微信支付对象
func New(appId, subAppId, mchId, subMchId, partnerKey, notifyUrl, spBillCreateIp, tradeType, signType, deviceInfo string, useSandbox bool, machCert []byte) (wx *WxPay) {
	wx = &WxPay{}
	wx.AppId = appId                   //小程序appId,是
	wx.SubAppId = subAppId             //子商户公众账号ID,否
	wx.MchId = mchId                   //商户号,是
	wx.SubMchId = subMchId             //子商户号
	wx.PartnerKey = partnerKey         //商户密钥,是
	wx.NotifyUrl = notifyUrl           //通知地址,是
	wx.SpBillCreateIp = spBillCreateIp //终端IP,是
	wx.DeviceInfo = deviceInfo         //设备号,否
	wx.SignType = signType             //签名类型,是sss
	wx.FeeType = DEFAULT_FEETYPE       //货币类型
	wx.TradeType = tradeType           //交易类型,是
	wx.UseSandbox = useSandbox         //是否测试环境
	wx.MachCert = machCert             // 商户证书,否
	wx.HttpClient = http.DefaultClient
	if machCert != nil { //带认证证书
		wx.HttpClientWithCert = http.DefaultClient
	}
	return
}

//统一下单
//场景：公共号支付、扫码支付、APP支付
func (wx *WxPay) UnifiedOrder(totalFee int64, outTradeNo, body, openId, detail, attach, goodsTag, productId,
limitPay, sceneInfo string) (orderResult *UnifiedOrderResult, err error) {
	if len(outTradeNo) == 0 || totalFee == 0 || (wx.TradeType == TRADE_TYPE_JSAPI && len(openId) == 0) || !checkWXPay(wx) {
		err = WXPAY_REQ_PARAMT_ERROR
		return
	}
	orderRequest := createUnifiedOrder(wx)
	if len(body) == 0 {
		orderRequest.Body = DEFAULT_ORDER_BODY
	} else {
		orderRequest.Body = body //商品描述,必填
	}
	orderRequest.OutTradeNo = outTradeNo //商户订单号,必填
	orderRequest.TotalFee = totalFee     //订单总金额，单位为分,必填
	orderRequest.OpenId = openId         //用户标识,必填
	orderRequest.Detail = detail         //商品详情,非必填
	orderRequest.Attach = attach         //附加数据,非必填
	orderRequest.GoodsTag = goodsTag     //订单优惠标记,非必填
	orderRequest.ProductId = productId   //商品ID,非必填
	orderRequest.LimitPay = limitPay     //指定支付方式,非必填
	orderRequest.SceneInfo = sceneInfo   //场景信息,非必填

	sign, err := signStruct(orderRequest, wx.PartnerKey) //计算签名
	if err != nil {
		return
	}

	orderRequest.Sign = sign
	orderXml, err := buildXML(orderRequest)
	if err != nil {
		err = WXPAY_XML_MARSHA_ERROR
		return
	}
	apiUrl := WXPAY_UNIFIEDORDER_URL
	if wx.UseSandbox {
		apiUrl = WXPAY_SANDBOX_UNIFIEDORDER_URL
	}

	respBody, err := requestOnce(wx.HttpClient, apiUrl, K_CONTENT_TYPE_FORM, orderXml)
	if err != nil {
		return
	}
	fmt.Println("统一下单结果", string(respBody))
	//xml转结构体
	orderRespond := unifiedOrderRespond{}
	err = parseXML(respBody, &orderRespond)
	if err != nil {
		err = WXPAY_XML_UNMARSHA_ERROR
		return
	}

	orderResult = &UnifiedOrderResult{}
	if SUCCESS == orderRespond.ReturnCode && SUCCESS == orderRespond.ResultCode { //成功的时候验证签名
		fmt.Println("统一下单成功", orderRespond)
		signResult, err := signStruct(orderRespond, wx.PartnerKey)
		if err != nil {
			err = WXPAY_SIGN_CAL_ERROR
			return nil, err
		}
		if signResult != orderRespond.Sign {
			err = WXPAY_SIGN_VERIFY_ERROR
			return nil, err
		}
		if TRADE_TYPE_JSAPI == wx.TradeType {
			rp, err := createRequestPayment(orderRespond, wx.PartnerKey, wx.SignType)
			if err != nil {
				return nil, err
			}
			orderResult.RequestPayment = rp
		}
		orderResult.ReturnCode = orderRespond.ReturnCode
		orderResult.ResultCode = orderRespond.ResultCode
		orderResult.PrepayId = orderRespond.PrepayId
	} else {
		fmt.Println("统一下单失败", orderRespond)
		orderResult.ReturnCode = orderRespond.ReturnCode
		orderResult.ReturnMsg = orderRespond.ReturnMsg
		orderResult.ResultCode = orderRespond.ResultCode
		orderResult.ErrCode = orderRespond.ErrCode
		orderResult.ErrCodeDes = orderRespond.ErrCodeDes
	}
	return orderResult, nil
}

//统一下单-支付结果通知
func (wx *WxPay) UnifiedOrderNotify(result []byte) (unifiedOrderNotifyResult *UnifiedOrderNotifyResult, err error) {
	if len(result) == 0 {
		err = WXPAY_REQ_PARAMT_ERROR
		return
	}
	fmt.Println("统一下单支付结果通知结果", string(result))
	//xml转结构体
	uResult := UnifiedOrderNotifyResult{}
	err = parseXML(result, &uResult)
	if err != nil {
		err = WXPAY_XML_UNMARSHA_ERROR
		return
	}
	if SUCCESS == uResult.ReturnCode { //成功的时候验证签名
		fmt.Println("统一下单成功", uResult)
		signResult, err := signStruct(uResult, wx.PartnerKey)
		if err != nil {
			err = WXPAY_SIGN_CAL_ERROR
			return nil, err
		}
		fmt.Println("签名计算结果：", signResult)
		fmt.Println("返回签名结果：", uResult.Sign)
		if signResult != uResult.Sign {
			err = WXPAY_SIGN_VERIFY_ERROR
			return nil, err
		}
		return &uResult, nil
	} else {
		err = errors.New(uResult.ReturnMsg)
		return nil, err
	}
}

//提交刷卡支付
//场景：刷卡支付
func (wx *WxPay) MicroPay(totalFee int64, outTradeNo, body, authCode, detail, attach, goodsTag,
limitPay, sceneInfo string) (micropayResult *MicropayResult, err error) {
	start:=time.Now().UnixNano()/1e6
	if len(outTradeNo) == 0 || totalFee == 0 || len(authCode) == 0 || !checkWXPay(wx) {
		err = WXPAY_REQ_PARAMT_ERROR
		return
	}
	orderRequest := createMicropayOrder(wx)
	if len(body) == 0 {
		orderRequest.Body = DEFAULT_ORDER_BODY
	} else {
		orderRequest.Body = body //商品描述,必填
	}
	orderRequest.OutTradeNo = outTradeNo                 //商户订单号,必填
	orderRequest.TotalFee = totalFee                     //订单总金额，单位为分,必填
	orderRequest.AuthCode = authCode                     //授权码,必填
	orderRequest.Detail = detail                         //商品详情,非必填
	orderRequest.Attach = attach                         //附加数据,非必填
	orderRequest.GoodsTag = goodsTag                     //订单优惠标记,非必填
	orderRequest.LimitPay = limitPay                     //指定支付方式,非必填
	orderRequest.SceneInfo = sceneInfo                   //场景信息,非必填
	sign, err := signStruct(orderRequest, wx.PartnerKey) //计算签名
	if err != nil {
		return
	}

	orderRequest.Sign = sign
	orderXml, err := buildXML(orderRequest)
	if err != nil {
		err = WXPAY_XML_MARSHA_ERROR
		return
	}
	apiUrl := WXPAY_MICROPAY_URL
	if wx.UseSandbox {
		apiUrl = WXPAY_SANDBOX_MICROPAY_URL
	}

	respBody, err := requestOnce(wx.HttpClient, apiUrl, K_CONTENT_TYPE_FORM, orderXml)
	if err != nil {
		return
	}
	fmt.Println("提交刷卡支付结果", string(respBody))
	//xml转结构体
	result := MicropayResult{}
	err = parseXML(respBody, &result)
	if err != nil {
		err = WXPAY_XML_UNMARSHA_ERROR
		return
	}

	if SUCCESS == result.ReturnCode { //成功的时候验证签名
		fmt.Println("提交刷卡支付成功", result)
		signResult, err := signStruct(result, wx.PartnerKey)
		if err != nil {
			err = WXPAY_SIGN_CAL_ERROR
			return nil, err
		}
		if signResult != result.Sign {
			err = WXPAY_SIGN_VERIFY_ERROR
			return nil, err
		}
		end:=time.Now().UnixNano()/1e6
		fmt.Println("刷卡时间",(end-start))
		return &result, nil
	} else {
		fmt.Println("提交刷卡支付失败", result)
		return &result, nil
	}
	return
}

//提交刷卡支付
//内置重试机制，最多60s
func (wx *WxPay) MicroPayWithPos(totalFee int64, outTradeNo, body, authCode, detail, attach, goodsTag,
limitPay, sceneInfo string) (micropayResult *MicropayResult, err error) {
	remainingTimeMs := 60 * 1000
	for {
		startTimestampMs := time.Now().Unix()
		readTimeoutMs := remainingTimeMs - 0
		if readTimeoutMs > 1000 {
			micropayResult, err = wx.MicroPay(totalFee, outTradeNo, body, authCode, detail, attach, goodsTag, limitPay, sceneInfo)
			if err != nil || micropayResult == nil {
				break
			}
			if SUCCESS == micropayResult.ReturnCode {
				if SUCCESS == micropayResult.ResultCode {
					break
				} else {
					if SYSTEMERROR == micropayResult.ErrCode || BANKERROR == micropayResult.ErrCode { //立即调用被扫订单结果查询API【查询订单API】
						fmt.Println("立即调用被扫订单结果查询API【查询订单API】")
						queryResult, err := wx.OrderQuery("", outTradeNo)
						if err == nil && queryResult != nil {
							fmt.Println("查询支付结果成功:queryResult.ResultCode=", queryResult.ResultCode)
							if SUCCESS == queryResult.ReturnCode && SUCCESS == queryResult.ResultCode { //支付成功
								micropayResult.ResultCode = queryResult.ResultCode
								micropayResult.ErrCode = queryResult.ErrCode
								micropayResult.ErrCodeDes = queryResult.ErrCodeDes
								micropayResult.TimeEnd = queryResult.TimeEnd
								break
							}
						}
						remainingTimeMs = remainingTimeMs - (int)(time.Now().Unix()-startTimestampMs)
						if remainingTimeMs <= 100 {
							break
						} else {
							if remainingTimeMs > 5*1000 {
								time.Sleep(5 * time.Second)
							} else {
								time.Sleep(1 * time.Second)
							}
							continue;
						}
					} else if USERPAYING == micropayResult.ErrCode { //商户系统可设置间隔时间(建议10秒)重新查询支付结果，直到支付成功或超时(建议30秒)
						fmt.Println("商户系统可设置间隔时间(建议10秒)重新查询支付结果，直到支付成功或超时(建议30秒)")
						remainingTimeMs = remainingTimeMs - (int)(time.Now().Unix()-startTimestampMs)
						if remainingTimeMs <= 100 {
							break
						} else {
							if remainingTimeMs > 10*1000 {
								time.Sleep(10 * time.Second)
							} else {
								time.Sleep(1 * time.Second)
							}
							queryResult, err := wx.OrderQuery("", outTradeNo)
							if err == nil && queryResult != nil {
								fmt.Println("查询支付结果成功:queryResult.ResultCode=", queryResult.ResultCode)
								if SUCCESS == queryResult.ReturnCode && SUCCESS == queryResult.ResultCode { //支付成功
									micropayResult.ResultCode = queryResult.ResultCode
									micropayResult.ErrCode = queryResult.ErrCode
									micropayResult.ErrCodeDes = queryResult.ErrCodeDes
									micropayResult.TimeEnd = queryResult.TimeEnd
									break
								}
							}
							continue
						}
					} else {
						break
					}
				}
			} else {
				break
			}
		} else {
			break
		}
	}

	return
}

//查询订单
//场景：刷卡支付、公共号支付、扫码支付、APP支付
func (wx *WxPay) OrderQuery(transactionId, outTradeNo string) (queryResult *OrderqueryResult, err error) {
	if len(transactionId) == 0 && len(outTradeNo) == 0 {
		err = WXPAY_REQ_PARAMT_ERROR
		return
	}
	queryRequest := orderQueryRequest{}
	queryRequest.AppId = wx.AppId
	queryRequest.MchId = wx.MchId
	queryRequest.NonceStr = GetRandomString(32)
	queryRequest.SignType = wx.SignType

	if len(transactionId) != 0 {
		queryRequest.TransactionId = transactionId
	} else {
		queryRequest.OutTradeNo = outTradeNo
	}

	sign, err := signStruct(queryRequest, wx.PartnerKey)
	if err != nil {
		return
	}
	queryRequest.Sign = sign
	orderXml, err := buildXML(queryRequest)
	if err != nil {
		err = WXPAY_XML_MARSHA_ERROR
		return
	}
	apiUrl := WXPAY_ORDERQUER_URL
	if wx.UseSandbox {
		apiUrl = WXPAY_SANDBOX_ORDERQUER_URL
	}
	respBody, err := requestOnce(wx.HttpClient, apiUrl, K_CONTENT_TYPE_FORM, orderXml)
	if err != nil {
		return
	}
	fmt.Println("查询订单结果", string(respBody))
	queryRespond := OrderqueryResult{}
	err = parseXML(respBody, &queryRespond)
	//err = json.Unmarshal(respBody, u)
	if err != nil {
		err = WXPAY_XML_UNMARSHA_ERROR
		return
	}

	if SUCCESS == queryRespond.ReturnCode && SUCCESS == queryRespond.ResultCode { //成功的时候验证签名
		fmt.Println("查询订单成功", queryRespond)
		signResult, err := signStruct(queryRespond, wx.PartnerKey)
		if err != nil {
			err = WXPAY_SIGN_CAL_ERROR
			return nil, err
		}
		if signResult != queryRespond.Sign {
			err = WXPAY_SIGN_VERIFY_ERROR
			return nil, err
		}

		return &queryRespond, nil
	} else {
		fmt.Println("查询订单失败", queryRespond)
		if (FAIL == queryRespond.ReturnCode) {
			err = errors.New(queryRespond.ReturnMsg)
		} else {
			err = errors.New(fmt.Sprintf("%v-%v", queryRespond.ErrCode, queryRespond.ErrCodeDes))
		}
		return nil, err
	}
	return nil, nil
}

//关闭订单
//场景：公共号支付、扫码支付、APP支付
func (wx *WxPay) CloseOrder() (queryResult *OrderqueryResult, err error) {

	return
}

//创建统一下单信息
func createUnifiedOrder(wx *WxPay) unifiedOrderRequest {
	requestData := unifiedOrderRequest{}
	requestData.AppId = wx.AppId
	requestData.MchId = wx.MchId
	requestData.DeviceInfo = wx.DeviceInfo
	requestData.SignType = wx.SignType
	requestData.FeeType = wx.FeeType
	requestData.SpBillCreateIp = wx.SpBillCreateIp
	requestData.NotifyUrl = wx.NotifyUrl
	requestData.TradeType = wx.TradeType
	requestData.NonceStr = GetRandomString(32)
	now := time.Now()
	mm, _ := time.ParseDuration(DEFAULT_EXPIRE)
	mm1 := now.Add(mm)
	requestData.TimeStart = now.Format("20060102150405")
	requestData.TimeExpire = mm1.Format("20060102150405")
	return requestData
}

//创建统一下单信息
func createMicropayOrder(wx *WxPay) micropayRequest {
	requestData := micropayRequest{}
	requestData.AppId = wx.AppId
	requestData.SubAppId = wx.SubAppId
	requestData.MchId = wx.MchId
	requestData.SubMchId = wx.SubMchId
	requestData.DeviceInfo = wx.DeviceInfo
	requestData.SignType = wx.SignType
	requestData.FeeType = wx.FeeType
	requestData.SpBillCreateIp = wx.SpBillCreateIp
	requestData.NonceStr = GetRandomString(32)
	now := time.Now()
	mm, _ := time.ParseDuration(DEFAULT_EXPIRE)
	mm1 := now.Add(mm)
	requestData.TimeStart = now.Format("20060102150405")
	requestData.TimeExpire = mm1.Format("20060102150405")
	return requestData
}

//检查参数是否完整
func checkWXPay(wx *WxPay) bool {
	if wx == nil {
		return false
	}
	if len(wx.AppId) == 0 || len(wx.MchId) == 0 || len(wx.NotifyUrl) == 0 ||
		len(wx.PartnerKey) == 0 || len(wx.SpBillCreateIp) == 0 || len(wx.SignType) == 0 ||
		len(wx.TradeType) == 0 {
		return false
	}
	return true
}

//生成微信小程序支付信息
func createRequestPayment(orderRespond unifiedOrderRespond, partnerKey string, signType string) (rp *WXRequestPayment, err error) {
	//组织返回数据
	requestPayment := WXRequestPayment{}
	requestPayment.AppId = orderRespond.AppId
	requestPayment.NonceStr = GetRandomString(32)
	requestPayment.TimeStamp = strconv.FormatInt(time.Now().Unix(), 10)
	requestPayment.Package = "prepay_id=" + orderRespond.PrepayId
	requestPayment.SignType = signType
	sign, err := signStruct(requestPayment, partnerKey)
	if err != nil {
		return nil, err
	}
	requestPayment.Sign = sign
	return &requestPayment, nil
}
