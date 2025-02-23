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

type DeletePageRuleRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /*   */
    Identifier string `json:"identifier"`
}

/*
 * param zone_identifier:  (Required)
 * param identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeletePageRuleRequest(
    zone_identifier string,
    identifier string,
) *DeletePageRuleRequest {

	return &DeletePageRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/pagerules/{identifier}",
			Method:  "DELETE",
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
 */
func NewDeletePageRuleRequestWithAllParams(
    zone_identifier string,
    identifier string,
) *DeletePageRuleRequest {

    return &DeletePageRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/pagerules/{identifier}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        Identifier: identifier,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeletePageRuleRequestWithoutParam() *DeletePageRuleRequest {

    return &DeletePageRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/pagerules/{identifier}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *DeletePageRuleRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}

/* param identifier: (Required) */
func (r *DeletePageRuleRequest) SetIdentifier(identifier string) {
    r.Identifier = identifier
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeletePageRuleRequest) GetRegionId() string {
    return ""
}

type DeletePageRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeletePageRuleResult `json:"result"`
}

type DeletePageRuleResult struct {
    Data string `json:"data"`
}