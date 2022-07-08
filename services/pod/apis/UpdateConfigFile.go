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
    pod "github.com/jdcloud-api/jdcloud-sdk-go/services/pod/models"
)

type UpdateConfigFileRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Name  */
    Name string `json:"name"`

    /* key 的有效字符包括字母、数字、-、_和.; <br>
value 每个value长度上限为32KB，整个data的长度不能超过1M; <br>
  */
    Data []pod.FileToPath `json:"data"`
}

/*
 * param regionId: Region ID (Required)
 * param name: Name (Required)
 * param data: key 的有效字符包括字母、数字、-、_和.; <br>
value 每个value长度上限为32KB，整个data的长度不能超过1M; <br>
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateConfigFileRequest(
    regionId string,
    name string,
    data []pod.FileToPath,
) *UpdateConfigFileRequest {

	return &UpdateConfigFileRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/configFiles/{name}:update",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
        Data: data,
	}
}

/*
 * param regionId: Region ID (Required)
 * param name: Name (Required)
 * param data: key 的有效字符包括字母、数字、-、_和.; <br>
value 每个value长度上限为32KB，整个data的长度不能超过1M; <br>
 (Required)
 */
func NewUpdateConfigFileRequestWithAllParams(
    regionId string,
    name string,
    data []pod.FileToPath,
) *UpdateConfigFileRequest {

    return &UpdateConfigFileRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/configFiles/{name}:update",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        Data: data,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateConfigFileRequestWithoutParam() *UpdateConfigFileRequest {

    return &UpdateConfigFileRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/configFiles/{name}:update",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UpdateConfigFileRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: Name(Required) */
func (r *UpdateConfigFileRequest) SetName(name string) {
    r.Name = name
}

/* param data: key 的有效字符包括字母、数字、-、_和.; <br>
value 每个value长度上限为32KB，整个data的长度不能超过1M; <br>
(Required) */
func (r *UpdateConfigFileRequest) SetData(data []pod.FileToPath) {
    r.Data = data
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateConfigFileRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateConfigFileResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateConfigFileResult `json:"result"`
}

type UpdateConfigFileResult struct {
    Name string `json:"name"`
}