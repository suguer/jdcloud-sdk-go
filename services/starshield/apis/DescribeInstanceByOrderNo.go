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
    starshield "github.com/jdcloud-api/jdcloud-sdk-go/services/starshield/models"
)

type DescribeInstanceByOrderNoRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /*   */
    OrderNumber string `json:"orderNumber"`
}

/*
 * param regionId: 地域ID (Required)
 * param orderNumber:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceByOrderNoRequest(
    regionId string,
    orderNumber string,
) *DescribeInstanceByOrderNoRequest {

	return &DescribeInstanceByOrderNoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{orderNumber}/describeInstance",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        OrderNumber: orderNumber,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param orderNumber:  (Required)
 */
func NewDescribeInstanceByOrderNoRequestWithAllParams(
    regionId string,
    orderNumber string,
) *DescribeInstanceByOrderNoRequest {

    return &DescribeInstanceByOrderNoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{orderNumber}/describeInstance",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        OrderNumber: orderNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceByOrderNoRequestWithoutParam() *DescribeInstanceByOrderNoRequest {

    return &DescribeInstanceByOrderNoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{orderNumber}/describeInstance",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeInstanceByOrderNoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param orderNumber: (Required) */
func (r *DescribeInstanceByOrderNoRequest) SetOrderNumber(orderNumber string) {
    r.OrderNumber = orderNumber
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceByOrderNoRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceByOrderNoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceByOrderNoResult `json:"result"`
}

type DescribeInstanceByOrderNoResult struct {
    DescribeInstancesRes starshield.DescribeInstancesRes `json:"describeInstancesRes"`
}