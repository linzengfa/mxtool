// Copyright (c) 2020.
// ALL Rights reserved.
// @Description def.go
// @Author moxiao
// @Date 2020/11/22 10:19

package tencent

import (
	"errors"

	v20180724 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/soe/v20180724"
)

type MXAi struct {
	SecretId  string
	SecretKey string
}

type OralProcessResult struct {
	// 发音精准度，取值范围[-1, 100]，当取-1时指完全不匹配，当为句子模式时，是所有已识别单词准确度的加权平均值。当为流式模式且请求中IsEnd未置1时，取值无意义
	PronAccuracy *float64 `json:"PronAccuracy,omitempty" name:"PronAccuracy"`

	// 发音流利度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
	PronFluency *float64 `json:"PronFluency,omitempty" name:"PronFluency"`

	// 发音完整度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
	PronCompletion *float64 `json:"PronCompletion,omitempty" name:"PronCompletion"`

	// 详细发音评估结果
	Words []*v20180724.WordRsp `json:"Words,omitempty" name:"Words" list`

	// 语音段唯一标识，一段语音一个SessionId
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 保存语音音频文件下载地址
	AudioUrl *string `json:"AudioUrl,omitempty" name:"AudioUrl"`

	// 断句中间结果，中间结果是局部最优而非全局最优的结果，所以中间结果有可能和最终整体结果对应部分不一致；中间结果的输出便于客户端UI更新；待用户发音完全结束后，系统会给出一个综合所有句子的整体结果。
	SentenceInfoSet []*v20180724.SentenceInfo `json:"SentenceInfoSet,omitempty" name:"SentenceInfoSet" list`

	// 评估 session 状态，“Evaluating"：评估中、"Failed"：评估失败、"Finished"：评估完成
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

var (
	MXAI_PARAM_ERROR        = errors.New("request param errors")
	MXAI_ASR_ERROR          = errors.New("asr errors")
	MXAI_ASR_FAILURE        = errors.New("asr failure")
	MXAI_HTTP_REQUEST_ERROR = errors.New("http request errors")
)
