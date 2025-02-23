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

type EditUniversalSSLSettingsRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /* 禁用通用SSL将从边缘上删除域的所有当前激活的通用SSL证书并且防止将来订购任何通用SSL证书。如果没有为域上载专用证书或自定义证书，访问者将无法通过HTTPS访问域。
通过禁用通用SSL，您知道以下星盾设置和首选项将导致访问者无法访问您的域，除非您上载了自定义证书或购买了专用证书。
  * HSTS
  * Always Use HTTPS
  * Opportunistic Encryption
  * Onion Routing
  * Any Page Rules redirecting traffic to HTTPS
类似地，在启用星盾代理时，在源站将任何HTTP重定向到HTTPS将导致用户在星盾的边缘没有有效证书的情况下无法访问您的站点。
如果您在星盾的边缘没有有效的自定义或专用证书，并且不确定是否启用了上述任何星盾设置，或者如果您的源站存在任何HTTP重定向，我们建议您的域保持启用通用SSL。
 (Optional) */
    Enabled *bool `json:"enabled"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEditUniversalSSLSettingsRequest(
    zone_identifier string,
) *EditUniversalSSLSettingsRequest {

	return &EditUniversalSSLSettingsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/ssl$$universal$$settings",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param enabled: 禁用通用SSL将从边缘上删除域的所有当前激活的通用SSL证书并且防止将来订购任何通用SSL证书。如果没有为域上载专用证书或自定义证书，访问者将无法通过HTTPS访问域。
通过禁用通用SSL，您知道以下星盾设置和首选项将导致访问者无法访问您的域，除非您上载了自定义证书或购买了专用证书。
  * HSTS
  * Always Use HTTPS
  * Opportunistic Encryption
  * Onion Routing
  * Any Page Rules redirecting traffic to HTTPS
类似地，在启用星盾代理时，在源站将任何HTTP重定向到HTTPS将导致用户在星盾的边缘没有有效证书的情况下无法访问您的站点。
如果您在星盾的边缘没有有效的自定义或专用证书，并且不确定是否启用了上述任何星盾设置，或者如果您的源站存在任何HTTP重定向，我们建议您的域保持启用通用SSL。
 (Optional)
 */
func NewEditUniversalSSLSettingsRequestWithAllParams(
    zone_identifier string,
    enabled *bool,
) *EditUniversalSSLSettingsRequest {

    return &EditUniversalSSLSettingsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/ssl$$universal$$settings",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        Enabled: enabled,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEditUniversalSSLSettingsRequestWithoutParam() *EditUniversalSSLSettingsRequest {

    return &EditUniversalSSLSettingsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/ssl$$universal$$settings",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *EditUniversalSSLSettingsRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}

/* param enabled: 禁用通用SSL将从边缘上删除域的所有当前激活的通用SSL证书并且防止将来订购任何通用SSL证书。如果没有为域上载专用证书或自定义证书，访问者将无法通过HTTPS访问域。
通过禁用通用SSL，您知道以下星盾设置和首选项将导致访问者无法访问您的域，除非您上载了自定义证书或购买了专用证书。
  * HSTS
  * Always Use HTTPS
  * Opportunistic Encryption
  * Onion Routing
  * Any Page Rules redirecting traffic to HTTPS
类似地，在启用星盾代理时，在源站将任何HTTP重定向到HTTPS将导致用户在星盾的边缘没有有效证书的情况下无法访问您的站点。
如果您在星盾的边缘没有有效的自定义或专用证书，并且不确定是否启用了上述任何星盾设置，或者如果您的源站存在任何HTTP重定向，我们建议您的域保持启用通用SSL。
(Optional) */
func (r *EditUniversalSSLSettingsRequest) SetEnabled(enabled bool) {
    r.Enabled = &enabled
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EditUniversalSSLSettingsRequest) GetRegionId() string {
    return ""
}

type EditUniversalSSLSettingsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EditUniversalSSLSettingsResult `json:"result"`
}

type EditUniversalSSLSettingsResult struct {
    Data starshield.UniversalSSLSetting `json:"data"`
}