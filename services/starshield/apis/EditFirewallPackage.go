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

type EditFirewallPackageRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /*   */
    Identifier string `json:"identifier"`

    /* 防火墙包的敏感度。 (Optional) */
    Sensitivity *string `json:"sensitivity"`

    /* 将对防火墙包下的规则执行的默认操作。 (Optional) */
    Action_mode *string `json:"action_mode"`
}

/*
 * param zone_identifier:  (Required)
 * param identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEditFirewallPackageRequest(
    zone_identifier string,
    identifier string,
) *EditFirewallPackageRequest {

	return &EditFirewallPackageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/firewall$$waf$$packages/{identifier}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
        Identifier: identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param identifier:  (Required)
 * param sensitivity: 防火墙包的敏感度。 (Optional)
 * param action_mode: 将对防火墙包下的规则执行的默认操作。 (Optional)
 */
func NewEditFirewallPackageRequestWithAllParams(
    zone_identifier string,
    identifier string,
    sensitivity *string,
    action_mode *string,
) *EditFirewallPackageRequest {

    return &EditFirewallPackageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/firewall$$waf$$packages/{identifier}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        Identifier: identifier,
        Sensitivity: sensitivity,
        Action_mode: action_mode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEditFirewallPackageRequestWithoutParam() *EditFirewallPackageRequest {

    return &EditFirewallPackageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/firewall$$waf$$packages/{identifier}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *EditFirewallPackageRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}

/* param identifier: (Required) */
func (r *EditFirewallPackageRequest) SetIdentifier(identifier string) {
    r.Identifier = identifier
}

/* param sensitivity: 防火墙包的敏感度。(Optional) */
func (r *EditFirewallPackageRequest) SetSensitivity(sensitivity string) {
    r.Sensitivity = &sensitivity
}

/* param action_mode: 将对防火墙包下的规则执行的默认操作。(Optional) */
func (r *EditFirewallPackageRequest) SetAction_mode(action_mode string) {
    r.Action_mode = &action_mode
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EditFirewallPackageRequest) GetRegionId() string {
    return ""
}

type EditFirewallPackageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EditFirewallPackageResult `json:"result"`
}

type EditFirewallPackageResult struct {
    Data starshield.WAFRulePackage `json:"data"`
}