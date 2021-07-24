// Copyright (c) 2020.
// ALL Rights reserved.
// @Description ars.go
// @Author moxiao
// @Date 2020/11/21 18:19

package tencent

import (
	v20180522 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aai/v20180522"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

//一句话识别语音识别
//voiceData：语音数据，当SourceType 值为1时必须填写，base64编码
//voiceDataLen：SourceType 值为1时必须填写，未进行base64编码时的数据长度
//voiceURL：语音URL，当SourceType值为0时必须填写
//sourceType：语音数据来源0：语音 URL；1：语音数据
//voiceFormat：识别音频的音频格式（支持mp3,wav）
//usrAudioKey：用户端对此任务的唯一标识，用户自助生成
func (mx *MXAi) Sentence(voiceData string, voiceDataLen int64, voiceURL string, sourceType uint64,
	voiceFormat string, usrAudioKey string) (content string, err error) {
	if (SourceTypeUrl == sourceType && len(voiceURL) == 0) ||
		(SourceTypeData == sourceType && (len(voiceData) == 0 || voiceDataLen < 0)) ||
		len(voiceFormat) == 0 || len(usrAudioKey) == 0 {
		err = MXAI_PARAM_ERROR
		return
	}
	//实例化请求client对象
	client, _ := common.NewClientWithSecretId(mx.SecretId, mx.SecretKey, regions.Guangzhou)
	// 实例化一个请求对象
	request := v20180522.NewSentenceRecognitionRequest()
	//设置请求参数
	request.ProjectId = common.Uint64Ptr(0)                           // 腾讯云项目 ID
	request.SubServiceType = common.Uint64Ptr(SubServiceTypeSentence) // 子服务类型
	request.EngSerViceType = common.StringPtr(EngServiceType16k)      // 引擎类型
	request.SourceType = common.Uint64Ptr(sourceType)                 // 语音数据来源。0：语音 URL；1：语音数据（post body）
	request.VoiceFormat = common.StringPtr(voiceFormat)               // 识别音频的音频格式（支持mp3,wav）
	request.UsrAudioKey = common.StringPtr(usrAudioKey)               // 用户端对此任务的唯一标识
	if SourceTypeUrl == sourceType {
		request.Url = common.StringPtr(voiceURL) // 语音 URL，公网可下载
	} else {
		request.Data = common.StringPtr(voiceData)      // 语音数据，base64编码
		request.DataLen = common.Int64Ptr(voiceDataLen) // 数据长度未进行base64编码时的数据长度
	}

	response := v20180522.NewSentenceRecognitionResponse()
	err = client.Send(request, response)
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		return
	}
	content = *response.Response.Result
	return
}
