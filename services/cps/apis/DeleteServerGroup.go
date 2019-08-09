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

type DeleteServerGroupRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 服务器组ID  */
    ServerGroupId string `json:"serverGroupId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param serverGroupId: 服务器组ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteServerGroupRequest(
    regionId string,
    serverGroupId string,
) *DeleteServerGroupRequest {

	return &DeleteServerGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/serverGroups/{serverGroupId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServerGroupId: serverGroupId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param serverGroupId: 服务器组ID (Required)
 */
func NewDeleteServerGroupRequestWithAllParams(
    regionId string,
    serverGroupId string,
) *DeleteServerGroupRequest {

    return &DeleteServerGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serverGroups/{serverGroupId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServerGroupId: serverGroupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteServerGroupRequestWithoutParam() *DeleteServerGroupRequest {

    return &DeleteServerGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serverGroups/{serverGroupId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DeleteServerGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param serverGroupId: 服务器组ID(Required) */
func (r *DeleteServerGroupRequest) SetServerGroupId(serverGroupId string) {
    r.ServerGroupId = serverGroupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteServerGroupRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteServerGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteServerGroupResult `json:"result"`
}

type DeleteServerGroupResult struct {
    Success bool `json:"success"`
}