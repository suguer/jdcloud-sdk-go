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

type UpgradeEngineVersionRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 升级后的版本号 (Optional) */
    Version *string `json:"version"`

    /* 版本升级的时间点,时间格式yyyy-mm-dd hh:mm:ss。不传或者传入空表示取消升级 (Optional) */
    Timing *string `json:"timing"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpgradeEngineVersionRequest(
    regionId string,
    instanceId string,
) *UpgradeEngineVersionRequest {

	return &UpgradeEngineVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:upgradeEngineVersion",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param version: 升级后的版本号 (Optional)
 * param timing: 版本升级的时间点,时间格式yyyy-mm-dd hh:mm:ss。不传或者传入空表示取消升级 (Optional)
 */
func NewUpgradeEngineVersionRequestWithAllParams(
    regionId string,
    instanceId string,
    version *string,
    timing *string,
) *UpgradeEngineVersionRequest {

    return &UpgradeEngineVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:upgradeEngineVersion",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Version: version,
        Timing: timing,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpgradeEngineVersionRequestWithoutParam() *UpgradeEngineVersionRequest {

    return &UpgradeEngineVersionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:upgradeEngineVersion",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *UpgradeEngineVersionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *UpgradeEngineVersionRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param version: 升级后的版本号(Optional) */
func (r *UpgradeEngineVersionRequest) SetVersion(version string) {
    r.Version = &version
}

/* param timing: 版本升级的时间点,时间格式yyyy-mm-dd hh:mm:ss。不传或者传入空表示取消升级(Optional) */
func (r *UpgradeEngineVersionRequest) SetTiming(timing string) {
    r.Timing = &timing
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpgradeEngineVersionRequest) GetRegionId() string {
    return r.RegionId
}

type UpgradeEngineVersionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpgradeEngineVersionResult `json:"result"`
}

type UpgradeEngineVersionResult struct {
}