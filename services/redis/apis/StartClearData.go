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
    redis "github.com/jdcloud-api/jdcloud-sdk-go/services/redis/models"
)

type StartClearDataRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 数据清理任务类型  */
    ClearType string `json:"clearType"`

    /* 匹配模式, 如: test*、*test、ab*cc*, 当节点为AllData、ExpiredData时可以忽略此参数 (Optional) */
    KeyPattern *string `json:"keyPattern"`

    /* key的过滤条件, 当节点为AllData、ExpiredData时可以忽略此参数 (Optional) */
    KeyFilter []redis.KeyFilter `json:"keyFilter"`

    /* 数据遍历的速率 (Optional) */
    QpsLimit *int `json:"qpsLimit"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param clearType: 数据清理任务类型 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStartClearDataRequest(
    regionId string,
    cacheInstanceId string,
    clearType string,
) *StartClearDataRequest {

	return &StartClearDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/startClearData",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        ClearType: clearType,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param clearType: 数据清理任务类型 (Required)
 * param keyPattern: 匹配模式, 如: test*、*test、ab*cc*, 当节点为AllData、ExpiredData时可以忽略此参数 (Optional)
 * param keyFilter: key的过滤条件, 当节点为AllData、ExpiredData时可以忽略此参数 (Optional)
 * param qpsLimit: 数据遍历的速率 (Optional)
 */
func NewStartClearDataRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    clearType string,
    keyPattern *string,
    keyFilter []redis.KeyFilter,
    qpsLimit *int,
) *StartClearDataRequest {

    return &StartClearDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/startClearData",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        ClearType: clearType,
        KeyPattern: keyPattern,
        KeyFilter: keyFilter,
        QpsLimit: qpsLimit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStartClearDataRequestWithoutParam() *StartClearDataRequest {

    return &StartClearDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/startClearData",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *StartClearDataRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *StartClearDataRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}
/* param clearType: 数据清理任务类型(Required) */
func (r *StartClearDataRequest) SetClearType(clearType string) {
    r.ClearType = clearType
}
/* param keyPattern: 匹配模式, 如: test*、*test、ab*cc*, 当节点为AllData、ExpiredData时可以忽略此参数(Optional) */
func (r *StartClearDataRequest) SetKeyPattern(keyPattern string) {
    r.KeyPattern = &keyPattern
}
/* param keyFilter: key的过滤条件, 当节点为AllData、ExpiredData时可以忽略此参数(Optional) */
func (r *StartClearDataRequest) SetKeyFilter(keyFilter []redis.KeyFilter) {
    r.KeyFilter = keyFilter
}
/* param qpsLimit: 数据遍历的速率(Optional) */
func (r *StartClearDataRequest) SetQpsLimit(qpsLimit int) {
    r.QpsLimit = &qpsLimit
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StartClearDataRequest) GetRegionId() string {
    return r.RegionId
}

type StartClearDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StartClearDataResult `json:"result"`
}

type StartClearDataResult struct {
}