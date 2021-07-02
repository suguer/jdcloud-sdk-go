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

type VideoResultsV2Request struct {

    core.JDCloudRequest

    /* 业务bizType，请联系客户经理获取 (Optional) */
    BizType *string `json:"bizType"`

    /* 接口版本v1 (Optional) */
    Version *string `json:"version"`

    /* 要查询的taskId (Optional) */
    TaskIds []string `json:"taskIds"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewVideoResultsV2Request(
) *VideoResultsV2Request {

	return &VideoResultsV2Request{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/video:resultsv2",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param bizType: 业务bizType，请联系客户经理获取 (Optional)
 * param version: 接口版本v1 (Optional)
 * param taskIds: 要查询的taskId (Optional)
 */
func NewVideoResultsV2RequestWithAllParams(
    bizType *string,
    version *string,
    taskIds []string,
) *VideoResultsV2Request {

    return &VideoResultsV2Request{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/video:resultsv2",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        BizType: bizType,
        Version: version,
        TaskIds: taskIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewVideoResultsV2RequestWithoutParam() *VideoResultsV2Request {

    return &VideoResultsV2Request{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/video:resultsv2",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param bizType: 业务bizType，请联系客户经理获取(Optional) */
func (r *VideoResultsV2Request) SetBizType(bizType string) {
    r.BizType = &bizType
}

/* param version: 接口版本v1(Optional) */
func (r *VideoResultsV2Request) SetVersion(version string) {
    r.Version = &version
}

/* param taskIds: 要查询的taskId(Optional) */
func (r *VideoResultsV2Request) SetTaskIds(taskIds []string) {
    r.TaskIds = taskIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r VideoResultsV2Request) GetRegionId() string {
    return ""
}

type VideoResultsV2Response struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result VideoResultsV2Result `json:"result"`
}

type VideoResultsV2Result struct {
    Result []censor.VideoResultV2 `json:"result"`
}