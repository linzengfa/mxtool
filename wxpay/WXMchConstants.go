/**********************************************
** @Des: WXMchConstants.go
** @Author: moxiao
** @Date:   2019/9/22 17:52
** @Last Modified by:  moxiao
** @Last Modified time: 2019/9/22 17:52
***********************************************/
package wxpay

const (
	//企业付款-接口URL地址
	WXMXCHPAY_TRANSFERS_URL = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers" //企业付款,用于企业向微信用户个人付款
	//企业付款-接口URL地址（沙箱）
	WXMXCHPAY_SANDBOX_TRANSFERS_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder" //企业付款,用于企业向微信用户个人付款
)

const (
	CHECK_NAME_NO_CHECK    = "NO_CHECK"    //不校验真实姓名
	CHECK_NAME_FORCE_CHECK = "FORCE_CHECK" //强校验真实姓名
)