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

type StopInstanceRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 云主机ID。  */
    InstanceId string `json:"instanceId"`

    /* 停机不计费模式。
该参数仅对按配置计费且系统盘为云硬盘的实例生效，并且不是专有宿主机中的实例。
配置停机不计费且停机后，实例部分将停止计费，且释放实例自身包含的资源（CPU/内存/GPU/本地数据盘）。
可选值：
`keepCharging`：停机后保持计费，不释放资源。
`stopCharging`：停机后停止计费，释放实例资源。默认值为空。
 (Optional) */
    ChargeOnStopped *string `json:"chargeOnStopped"`

}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceId: 云主机ID。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStopInstanceRequest(
    regionId string,
    instanceId string,
) *StopInstanceRequest {

	return &StopInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:stopInstance",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceId: 云主机ID。 (Required)
 * param chargeOnStopped: 停机不计费模式。
该参数仅对按配置计费且系统盘为云硬盘的实例生效，并且不是专有宿主机中的实例。
配置停机不计费且停机后，实例部分将停止计费，且释放实例自身包含的资源（CPU/内存/GPU/本地数据盘）。
可选值：
`keepCharging`：停机后保持计费，不释放资源。
`stopCharging`：停机后停止计费，释放实例资源。默认值为空。
 (Optional)
 */
func NewStopInstanceRequestWithAllParams(
    regionId string,
    instanceId string,
    chargeOnStopped *string,
) *StopInstanceRequest {

    return &StopInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:stopInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ChargeOnStopped: chargeOnStopped,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStopInstanceRequestWithoutParam() *StopInstanceRequest {

    return &StopInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:stopInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *StopInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 云主机ID。(Required) */
func (r *StopInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param chargeOnStopped: 停机不计费模式。
该参数仅对按配置计费且系统盘为云硬盘的实例生效，并且不是专有宿主机中的实例。
配置停机不计费且停机后，实例部分将停止计费，且释放实例自身包含的资源（CPU/内存/GPU/本地数据盘）。
可选值：
`keepCharging`：停机后保持计费，不释放资源。
`stopCharging`：停机后停止计费，释放实例资源。默认值为空。
(Optional) */
func (r *StopInstanceRequest) SetChargeOnStopped(chargeOnStopped string) {
    r.ChargeOnStopped = &chargeOnStopped
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StopInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type StopInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StopInstanceResult `json:"result"`
}

type StopInstanceResult struct {
}