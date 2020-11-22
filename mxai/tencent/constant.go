// Copyright (c) 2020.
// ALL Rights reserved.
// @Description constant.go
// @Author moxiao
// @Date 2020/11/22 10:19

package tencent

const (
	METHOD_POST string = "POST"
	METHOD_GET  string = "GET"
)

const (
	ASR_URL                   = "https://aai.tencentcloudapi.com"
	SR_ACTION                 = "SentenceRecognition"
	SR_API_VERSION            = "2018-05-22"
	SR_SERVICE_NAME           = "aai"
	SIGNATURE_METHOD_SHA256   = "HmacSHA256"
	SIGNATURE_METHOD_SHA1     = "HmacSHA1"
	SUB_SERVICE_TYPE_SENTENCE = 2
	ENGSERVICETYPE_16K        = "16k"
)

const (
	SOURCE_TYPE_URL  = 0 //0：语音 URL
	SOURCE_TYPE_DATA = 1 //1：语音数据（post body）
)

const (
	VOICE_FILE_TYPE_RAW = 1 //语音文件类型 1:raw
	VOICE_FILE_TYPE_WAV = 2 //语音文件类型2:wav
	VOICE_FILE_TYPE_MP3 = 3 //语音文件类型 3:mp3
)

const (
	VOICE_FILE_FORMAT_PCM = "pcm"
	VOICE_FILE_FORMAT_WAV = "wav"
	VOICE_FILE_FORMAT_AMR = "amr"
	VOICE_FILE_FORMAT_MP3 = "mp3"
)

const (
	WORKMODE_STREAM     = 0 //0：流式分片
	WORKMODE_NON_STREAM = 1 //1：非流式一次性评估
)

const (
	EVALMODE_WORD     = 0 //评估模式，0：词模式，,1：:句子模式，2：段落模式，3：自由说模式
	EVALMODE_SENTENCE = 1 //评估模式，0：词模式，,1：:句子模式，2：段落模式，3：自由说模式
	EVALMODE_PART     = 2 //评估模式，0：词模式，,1：:句子模式，2：段落模式，3：自由说模式
	EVALMODE_FREE     = 3 //评估模式，0：词模式，,1：:句子模式，2：段落模式，3：自由说模式
)

const (
	SERVERTYPE_EN = 0 //评估语言，0：英文，1：中文。
	SERVERTYPE_CN = 1 //评估语言，0：英文，1：中文。
)
