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

type UniversalSSLSettingsDetailsRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUniversalSSLSettingsDetailsRequest(
    zone_identifier string,
) *UniversalSSLSettingsDetailsRequest {

	return &UniversalSSLSettingsDetailsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/ssl$$universal$$settings",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 */
func NewUniversalSSLSettingsDetailsRequestWithAllParams(
    zone_identifier string,
) *UniversalSSLSettingsDetailsRequest {

    return &UniversalSSLSettingsDetailsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/ssl$$universal$$settings",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUniversalSSLSettingsDetailsRequestWithoutParam() *UniversalSSLSettingsDetailsRequest {

    return &UniversalSSLSettingsDetailsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/ssl$$universal$$settings",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *UniversalSSLSettingsDetailsRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UniversalSSLSettingsDetailsRequest) GetRegionId() string {
    return ""
}

type UniversalSSLSettingsDetailsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UniversalSSLSettingsDetailsResult `json:"result"`
}

type UniversalSSLSettingsDetailsResult struct {
    Data starshield.UniversalSSLSetting `json:"data"`
}