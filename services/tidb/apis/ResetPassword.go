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

type ResetPasswordRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 账号名，在同一个实例中账号名不能重复  */
    AccountName string `json:"accountName"`

    /* 新密码  */
    AccountPassword string `json:"accountPassword"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param accountName: 账号名，在同一个实例中账号名不能重复 (Required)
 * param accountPassword: 新密码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewResetPasswordRequest(
    regionId string,
    instanceId string,
    accountName string,
    accountPassword string,
) *ResetPasswordRequest {

	return &ResetPasswordRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/accounts/{accountName}:resetPassword",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        AccountName: accountName,
        AccountPassword: accountPassword,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param accountName: 账号名，在同一个实例中账号名不能重复 (Required)
 * param accountPassword: 新密码 (Required)
 */
func NewResetPasswordRequestWithAllParams(
    regionId string,
    instanceId string,
    accountName string,
    accountPassword string,
) *ResetPasswordRequest {

    return &ResetPasswordRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/accounts/{accountName}:resetPassword",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        AccountName: accountName,
        AccountPassword: accountPassword,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewResetPasswordRequestWithoutParam() *ResetPasswordRequest {

    return &ResetPasswordRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/accounts/{accountName}:resetPassword",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *ResetPasswordRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *ResetPasswordRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param accountName: 账号名，在同一个实例中账号名不能重复(Required) */
func (r *ResetPasswordRequest) SetAccountName(accountName string) {
    r.AccountName = accountName
}

/* param accountPassword: 新密码(Required) */
func (r *ResetPasswordRequest) SetAccountPassword(accountPassword string) {
    r.AccountPassword = accountPassword
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ResetPasswordRequest) GetRegionId() string {
    return r.RegionId
}

type ResetPasswordResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ResetPasswordResult `json:"result"`
}

type ResetPasswordResult struct {
}