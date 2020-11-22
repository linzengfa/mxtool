// Copyright (c) 2020.
// ALL Rights reserved.
// @Description ars.go
// @Author moxiao
// @Date 2020/11/21 18:19

package aliyun

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

	"time"
)

func (mg *MGAi) FileASR(fileURL string) (content string, err error) {
	if len(fileURL) == 0 {
		err = MXAI_PARAM_ERROR
		return
	}
	client, err := sdk.NewClientWithAccessKey(REGION_ID, mg.AccessKeyId, mg.AccessKeySecret)
	if err != nil {
		return
	}
	/**
	 * 创建录音文件识别请求，设置请求参数
	 */
	postRequest := requests.NewCommonRequest()
	postRequest.Domain = DOMAIN
	postRequest.Version = API_VERSION
	postRequest.Product = PRODUCT
	postRequest.ApiName = POST_REQUEST_ACTION
	postRequest.Method = METHOD_POST
	// 设置task，以JSON字符串的格式设置到请求中
	mapTask := make(map[string]string)
	mapTask[KEY_APP_KEY] = mg.AppKey
	mapTask[KEY_FILE_LINK] = fileURL
	task, err := json.Marshal(mapTask)
	if err != nil {
		return
	}
	postRequest.QueryParams[KEY_TASK] = string(task)

	/**
	* 提交录音文件识别请求，处理服务端返回的响应
	 */
	postResponse, err := client.ProcessCommonRequest(postRequest)
	if err != nil {
		return
	}

	postResponseContent := postResponse.GetHttpContentString()
	fmt.Println(postResponseContent)
	if postResponse.GetHttpStatus() != 200 {
		err = MXAI_HTTP_REQUEST_ERROR
		return
	}
	var postMapResult map[string]interface{}
	err = json.Unmarshal([]byte(postResponseContent), &postMapResult)
	if err != nil {
		return
	}

	// 获取录音文件识别请求任务的ID，以供识别结果查询使用
	var taskId string = ""
	var statusText string = ""
	statusText = postMapResult[KEY_STATUS_TEXT].(string)
	if statusText == "SUCCESS" {
		fmt.Println("录音文件识别请求成功响应!")
		taskId = postMapResult[KEY_TASK_ID].(string)
	} else {
		fmt.Println("录音文件识别请求失败!")
		err = MXAI_ASR_ERROR
		return
	}

	/**
	 * 创建识别结果查询请求，并设置taskId作为查询参数
	 */
	getRequest := requests.NewCommonRequest()
	getRequest.Domain = DOMAIN
	getRequest.Version = API_VERSION
	getRequest.Product = PRODUCT
	getRequest.ApiName = GET_REQUEST_ACTION
	getRequest.Method = METHOD_GET
	getRequest.QueryParams[KEY_TASK_ID] = taskId

	/**
	 * 提交识别结果查询请求
	 * 以轮询的方式进行识别结果的查询，直到服务端返回的状态描述为“SUCCESS”、“SUCCESS_WITH_NO_VALID_FRAGMENT”，或者为错误描述，则结束轮询。
	 */
	statusText = ""
	for true {
		getResponse, err := client.ProcessCommonRequest(getRequest)
		if err != nil {
			break
		}
		getResponseContent := getResponse.GetHttpContentString()
		fmt.Println("识别查询结果：", getResponseContent)
		if getResponse.GetHttpStatus() != 200 {
			fmt.Println("识别结果查询请求失败，Http错误码：", getResponse.GetHttpStatus())
			err = MXAI_HTTP_REQUEST_ERROR
			break
		}
		var getMapResult map[string]interface{}
		err = json.Unmarshal([]byte(getResponseContent), &getMapResult)
		if err != nil {
			break
		}
		statusText = getMapResult[KEY_STATUS_TEXT].(string)
		if statusText == "RUNNING" || statusText == "QUEUEING" {
			// 继续轮询
			time.Sleep(3 * time.Second)
		} else {
			// 退出轮询
			break
		}
	}

	if statusText == "SUCCESS" || statusText == "SUCCESS_WITH_NO_VALID_FRAGMENT" {
		fmt.Println("录音文件识别成功！")
		content = statusText
		return
	} else {
		fmt.Println("录音文件识别失败！")
		err = MXAI_ASR_FAILURE
		return
	}
}
