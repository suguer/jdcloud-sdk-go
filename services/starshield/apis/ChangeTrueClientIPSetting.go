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

type ChangeTrueClientIPSettingRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /* on - 开启；off - 关闭 (Optional) */
    Value *string `json:"value"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewChangeTrueClientIPSettingRequest(
    zone_identifier string,
) *ChangeTrueClientIPSettingRequest {

	return &ChangeTrueClientIPSettingRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/settings$$true_client_ip_header",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param value: on - 开启；off - 关闭 (Optional)
 */
func NewChangeTrueClientIPSettingRequestWithAllParams(
    zone_identifier string,
    value *string,
) *ChangeTrueClientIPSettingRequest {

    return &ChangeTrueClientIPSettingRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/settings$$true_client_ip_header",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        Value: value,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewChangeTrueClientIPSettingRequestWithoutParam() *ChangeTrueClientIPSettingRequest {

    return &ChangeTrueClientIPSettingRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/settings$$true_client_ip_header",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *ChangeTrueClientIPSettingRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}

/* param value: on - 开启；off - 关闭(Optional) */
func (r *ChangeTrueClientIPSettingRequest) SetValue(value string) {
    r.Value = &value
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ChangeTrueClientIPSettingRequest) GetRegionId() string {
    return ""
}

type ChangeTrueClientIPSettingResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ChangeTrueClientIPSettingResult `json:"result"`
}

type ChangeTrueClientIPSettingResult struct {
    Data starshield.ZoneSetting `json:"data"`
}