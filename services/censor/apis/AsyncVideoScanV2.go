// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    censor "github.com/jdcloud-api/jdcloud-sdk-go/services/censor/models"
)

type AsyncVideoScanV2Request struct {

    core.JDCloudRequest

    /* 业务bizType，请联系客户经理获取 (Optional) */
    BizType *string `json:"bizType"`

    /* 最大长度512, 点播视频地址 (Optional) */
    Url *string `json:"url"`

    /* 最大长度128，点播视频唯一标识 (Optional) */
    DataId *string `json:"dataId"`

    /* 接口版本号，可选值 v3.2 (Optional) */
    Version *string `json:"version"`

    /* 最大长度512，视频名称 (Optional) */
    Title *string `json:"title"`

    /* 最大长度512，数据回调参数，产品根据业务情况自行设计，当获取离线检测结果时，内容安全服务会返回该字段 (Optional) */
    Callback *string `json:"callback"`

    /* 最大长度256，离线结果回调通知到客户的URL。主动回调数据接口超时时间设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性。 (Optional) */
    CallbackUrl *string `json:"callbackUrl"`

    /* 最大长度64，客户个性化视频唯一性标识，传入后，将以此值作为重复检测依据，若不传，默认以URL作为查重依据,如果重复提交会被拒绝，返回报错信息请求重复，以及原提交taskID值，具体返回请查看响应示例 (Optional) */
    UniqueKey *string `json:"uniqueKey"`

    /* 最大长度64，客户个性化视频唯一性标识，传入后，将以此值作为重复检测依据，若不传，默认以URL作为查重依据,如果重复提交会被拒绝，返回报错信息请求重复，以及原提交taskID值，具体返回请查看响应示例 (Optional) */
    ScFrequency *int `json:"scFrequency"`

    /* 高级截帧设置，此项填写，默认截帧策略失效 (Optional) */
    AdvancedFrequency *censor.AdvancedFrequency `json:"advancedFrequency"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAsyncVideoScanV2Request(
) *AsyncVideoScanV2Request {

	return &AsyncVideoScanV2Request{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/video:asyncscanv2",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param bizType: 业务bizType，请联系客户经理获取 (Optional)
 * param url: 最大长度512, 点播视频地址 (Optional)
 * param dataId: 最大长度128，点播视频唯一标识 (Optional)
 * param version: 接口版本号，可选值 v3.2 (Optional)
 * param title: 最大长度512，视频名称 (Optional)
 * param callback: 最大长度512，数据回调参数，产品根据业务情况自行设计，当获取离线检测结果时，内容安全服务会返回该字段 (Optional)
 * param callbackUrl: 最大长度256，离线结果回调通知到客户的URL。主动回调数据接口超时时间设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性。 (Optional)
 * param uniqueKey: 最大长度64，客户个性化视频唯一性标识，传入后，将以此值作为重复检测依据，若不传，默认以URL作为查重依据,如果重复提交会被拒绝，返回报错信息请求重复，以及原提交taskID值，具体返回请查看响应示例 (Optional)
 * param scFrequency: 最大长度64，客户个性化视频唯一性标识，传入后，将以此值作为重复检测依据，若不传，默认以URL作为查重依据,如果重复提交会被拒绝，返回报错信息请求重复，以及原提交taskID值，具体返回请查看响应示例 (Optional)
 * param advancedFrequency: 高级截帧设置，此项填写，默认截帧策略失效 (Optional)
 */
func NewAsyncVideoScanV2RequestWithAllParams(
    bizType *string,
    url *string,
    dataId *string,
    version *string,
    title *string,
    callback *string,
    callbackUrl *string,
    uniqueKey *string,
    scFrequency *int,
    advancedFrequency *censor.AdvancedFrequency,
) *AsyncVideoScanV2Request {

    return &AsyncVideoScanV2Request{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/video:asyncscanv2",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        BizType: bizType,
        Url: url,
        DataId: dataId,
        Version: version,
        Title: title,
        Callback: callback,
        CallbackUrl: callbackUrl,
        UniqueKey: uniqueKey,
        ScFrequency: scFrequency,
        AdvancedFrequency: advancedFrequency,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAsyncVideoScanV2RequestWithoutParam() *AsyncVideoScanV2Request {

    return &AsyncVideoScanV2Request{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/video:asyncscanv2",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param bizType: 业务bizType，请联系客户经理获取(Optional) */
func (r *AsyncVideoScanV2Request) SetBizType(bizType string) {
    r.BizType = &bizType
}

/* param url: 最大长度512, 点播视频地址(Optional) */
func (r *AsyncVideoScanV2Request) SetUrl(url string) {
    r.Url = &url
}

/* param dataId: 最大长度128，点播视频唯一标识(Optional) */
func (r *AsyncVideoScanV2Request) SetDataId(dataId string) {
    r.DataId = &dataId
}

/* param version: 接口版本号，可选值 v3.2(Optional) */
func (r *AsyncVideoScanV2Request) SetVersion(version string) {
    r.Version = &version
}

/* param title: 最大长度512，视频名称(Optional) */
func (r *AsyncVideoScanV2Request) SetTitle(title string) {
    r.Title = &title
}

/* param callback: 最大长度512，数据回调参数，产品根据业务情况自行设计，当获取离线检测结果时，内容安全服务会返回该字段(Optional) */
func (r *AsyncVideoScanV2Request) SetCallback(callback string) {
    r.Callback = &callback
}

/* param callbackUrl: 最大长度256，离线结果回调通知到客户的URL。主动回调数据接口超时时间设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性。(Optional) */
func (r *AsyncVideoScanV2Request) SetCallbackUrl(callbackUrl string) {
    r.CallbackUrl = &callbackUrl
}

/* param uniqueKey: 最大长度64，客户个性化视频唯一性标识，传入后，将以此值作为重复检测依据，若不传，默认以URL作为查重依据,如果重复提交会被拒绝，返回报错信息请求重复，以及原提交taskID值，具体返回请查看响应示例(Optional) */
func (r *AsyncVideoScanV2Request) SetUniqueKey(uniqueKey string) {
    r.UniqueKey = &uniqueKey
}

/* param scFrequency: 最大长度64，客户个性化视频唯一性标识，传入后，将以此值作为重复检测依据，若不传，默认以URL作为查重依据,如果重复提交会被拒绝，返回报错信息请求重复，以及原提交taskID值，具体返回请查看响应示例(Optional) */
func (r *AsyncVideoScanV2Request) SetScFrequency(scFrequency int) {
    r.ScFrequency = &scFrequency
}

/* param advancedFrequency: 高级截帧设置，此项填写，默认截帧策略失效(Optional) */
func (r *AsyncVideoScanV2Request) SetAdvancedFrequency(advancedFrequency *censor.AdvancedFrequency) {
    r.AdvancedFrequency = advancedFrequency
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AsyncVideoScanV2Request) GetRegionId() string {
    return ""
}

type AsyncVideoScanV2Response struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AsyncVideoScanV2Result `json:"result"`
}

type AsyncVideoScanV2Result struct {
    Result censor.VideoTaskData `json:"result"`
}