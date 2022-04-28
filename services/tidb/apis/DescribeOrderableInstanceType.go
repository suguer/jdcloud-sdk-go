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
    tidb "github.com/jdcloud-api/jdcloud-sdk-go/services/tidb/models"
)

type DescribeOrderableInstanceTypeRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: 地域代码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeOrderableInstanceTypeRequest(
    regionId string,
) *DescribeOrderableInstanceTypeRequest {

	return &DescribeOrderableInstanceTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/orderableInstanceType",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 */
func NewDescribeOrderableInstanceTypeRequestWithAllParams(
    regionId string,
) *DescribeOrderableInstanceTypeRequest {

    return &DescribeOrderableInstanceTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/orderableInstanceType",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeOrderableInstanceTypeRequestWithoutParam() *DescribeOrderableInstanceTypeRequest {

    return &DescribeOrderableInstanceTypeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/orderableInstanceType",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeOrderableInstanceTypeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeOrderableInstanceTypeRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeOrderableInstanceTypeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeOrderableInstanceTypeResult `json:"result"`
}

type DescribeOrderableInstanceTypeResult struct {
    EngineStatus int `json:"engineStatus"`
    OrderableAZs []tidb.Az `json:"orderableAZs"`
}