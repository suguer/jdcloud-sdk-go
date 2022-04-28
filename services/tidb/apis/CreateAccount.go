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

type CreateAccountRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 账号名，在同一个实例中，账号名不能重复  */
    AccountName string `json:"accountName"`

    /* 密码  */
    AccountPassword string `json:"accountPassword"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param accountName: 账号名，在同一个实例中，账号名不能重复 (Required)
 * param accountPassword: 密码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateAccountRequest(
    regionId string,
    instanceId string,
    accountName string,
    accountPassword string,
) *CreateAccountRequest {

	return &CreateAccountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/accounts",
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
 * param accountName: 账号名，在同一个实例中，账号名不能重复 (Required)
 * param accountPassword: 密码 (Required)
 */
func NewCreateAccountRequestWithAllParams(
    regionId string,
    instanceId string,
    accountName string,
    accountPassword string,
) *CreateAccountRequest {

    return &CreateAccountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/accounts",
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
func NewCreateAccountRequestWithoutParam() *CreateAccountRequest {

    return &CreateAccountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/accounts",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *CreateAccountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *CreateAccountRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param accountName: 账号名，在同一个实例中，账号名不能重复(Required) */
func (r *CreateAccountRequest) SetAccountName(accountName string) {
    r.AccountName = accountName
}

/* param accountPassword: 密码(Required) */
func (r *CreateAccountRequest) SetAccountPassword(accountPassword string) {
    r.AccountPassword = accountPassword
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateAccountRequest) GetRegionId() string {
    return r.RegionId
}

type CreateAccountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateAccountResult `json:"result"`
}

type CreateAccountResult struct {
}