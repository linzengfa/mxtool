// Copyright (c) 2020.
// ALL Rights reserved.
// @Description soe.go
// @Author moxiao
// @Date 2020/11/22 10:19

package tencent

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	v20180724 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/soe/v20180724"
)

const SoeAPIVersion = "2018-07-24"

func (mx *MXAi) OralProcess(userVoiceData string, voiceFileType int64, refText string, workMode int64, evalMode int64,
	scoreCoeff float64, sessionId string, serverType int64) (content *OralProcessResult, err error) {
	if len(userVoiceData) == 0 || len(refText) == 0 || len(sessionId) == 0 ||
		(serverType != 0 && serverType != 1) {
		err = MXAI_PARAM_ERROR
		return
	}

	//实例化请求client对象
	client, _ := common.NewClientWithSecretId(mx.SecretId, mx.SecretKey, regions.Guangzhou)

	request := v20180724.NewTransmitOralProcessWithInitRequest()
	request.WorkMode = common.Int64Ptr(workMode)
	request.EvalMode = common.Int64Ptr(evalMode)
	request.ScoreCoeff = common.Float64Ptr(scoreCoeff)
	request.SessionId = common.StringPtr(sessionId)
	request.RefText = common.StringPtr(refText)
	request.SeqId = common.Int64Ptr(1)
	request.IsEnd = common.Int64Ptr(1)
	request.VoiceFileType = common.Int64Ptr(voiceFileType)
	request.VoiceEncodeType = common.Int64Ptr(1)
	request.ServerType = common.Int64Ptr(serverType)
	request.StorageMode = common.Int64Ptr(1)
	request.IsAsync = common.Int64Ptr(0)
	// 设置base64加密后的语音数据
	request.UserVoiceData = common.StringPtr(userVoiceData)
	request.SetHttpMethod(METHOD_POST)

	response := v20180724.NewTransmitOralProcessWithInitResponse()
	err = client.Send(request, response)

	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Printf(response.ToJsonString())

	content = &OralProcessResult{
		response.Response.PronAccuracy,
		response.Response.PronFluency,
		response.Response.PronCompletion,
		response.Response.Words,
		response.Response.SessionId,
		response.Response.AudioUrl,
		response.Response.SentenceInfoSet,
		response.Response.Status,
		response.Response.RequestId,
	}

	return
}
