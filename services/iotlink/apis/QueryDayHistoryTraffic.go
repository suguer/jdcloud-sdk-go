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
)

type QueryDayHistoryTrafficRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 物联网卡日历史流量查询请求参数  */
    RequestParam string `json:"requestParam"`
}

/*
 * param regionId: Region ID (Required)
 * param requestParam: 物联网卡日历史流量查询请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryDayHistoryTrafficRequest(
    regionId string,
    requestParam string,
) *QueryDayHistoryTrafficRequest {

	return &QueryDayHistoryTrafficRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/queryDayHistoryTraffic",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RequestParam: requestParam,
	}
}

/*
 * param regionId: Region ID (Required)
 * param requestParam: 物联网卡日历史流量查询请求参数 (Required)
 */
func NewQueryDayHistoryTrafficRequestWithAllParams(
    regionId string,
    requestParam string,
) *QueryDayHistoryTrafficRequest {

    return &QueryDayHistoryTrafficRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/queryDayHistoryTraffic",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RequestParam: requestParam,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryDayHistoryTrafficRequestWithoutParam() *QueryDayHistoryTrafficRequest {

    return &QueryDayHistoryTrafficRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/queryDayHistoryTraffic",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *QueryDayHistoryTrafficRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param requestParam: 物联网卡日历史流量查询请求参数(Required) */
func (r *QueryDayHistoryTrafficRequest) SetRequestParam(requestParam string) {
    r.RequestParam = requestParam
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryDayHistoryTrafficRequest) GetRegionId() string {
    return r.RegionId
}

type QueryDayHistoryTrafficResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryDayHistoryTrafficResult `json:"result"`
}

type QueryDayHistoryTrafficResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Result []interface{} `json:"result"`
}