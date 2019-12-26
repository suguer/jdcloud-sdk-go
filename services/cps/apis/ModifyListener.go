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
    cps "github.com/jdcloud-api/jdcloud-sdk-go/services/cps/models"
)

type ModifyListenerRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 监听器ID  */
    ListenerId string `json:"listenerId"`

    /* 调度算法 (Optional) */
    Algorithm *string `json:"algorithm"`

    /* 会话保持 (Optional) */
    StickySession *string `json:"stickySession"`

    /* 是否获取真实ip，取值范围on|off (Optional) */
    RealIp *string `json:"realIp"`

    /* 名称 (Optional) */
    Name *string `json:"name"`

    /* 描述 (Optional) */
    Description *string `json:"description"`

    /* 健康检查 (Optional) */
    HealthCheck *string `json:"healthCheck"`

    /* 健康检查响应的最大超时时间 (Optional) */
    HealthCheckTimeout *int `json:"healthCheckTimeout"`

    /* 健康检查响应的最大间隔时间 (Optional) */
    HealthCheckInterval *int `json:"healthCheckInterval"`

    /* 健康检查结果为success的阈值 (Optional) */
    HealthyThreshold *int `json:"healthyThreshold"`

    /* 健康检查结果为fail的阈值 (Optional) */
    UnhealthyThreshold *int `json:"unhealthyThreshold"`

    /* 服务器组id (Optional) */
    ServerGroupId *string `json:"serverGroupId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param listenerId: 监听器ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyListenerRequest(
    regionId string,
    listenerId string,
) *ModifyListenerRequest {

	return &ModifyListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/listeners/{listenerId}:modifyListenerAttributes",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ListenerId: listenerId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param listenerId: 监听器ID (Required)
 * param algorithm: 调度算法 (Optional)
 * param stickySession: 会话保持 (Optional)
 * param realIp: 是否获取真实ip，取值范围on|off (Optional)
 * param name: 名称 (Optional)
 * param description: 描述 (Optional)
 * param healthCheck: 健康检查 (Optional)
 * param healthCheckTimeout: 健康检查响应的最大超时时间 (Optional)
 * param healthCheckInterval: 健康检查响应的最大间隔时间 (Optional)
 * param healthyThreshold: 健康检查结果为success的阈值 (Optional)
 * param unhealthyThreshold: 健康检查结果为fail的阈值 (Optional)
 * param serverGroupId: 服务器组id (Optional)
 */
func NewModifyListenerRequestWithAllParams(
    regionId string,
    listenerId string,
    algorithm *string,
    stickySession *string,
    realIp *string,
    name *string,
    description *string,
    healthCheck *string,
    healthCheckTimeout *int,
    healthCheckInterval *int,
    healthyThreshold *int,
    unhealthyThreshold *int,
    serverGroupId *string,
) *ModifyListenerRequest {

    return &ModifyListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/{listenerId}:modifyListenerAttributes",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ListenerId: listenerId,
        Algorithm: algorithm,
        StickySession: stickySession,
        RealIp: realIp,
        Name: name,
        Description: description,
        HealthCheck: healthCheck,
        HealthCheckTimeout: healthCheckTimeout,
        HealthCheckInterval: healthCheckInterval,
        HealthyThreshold: healthyThreshold,
        UnhealthyThreshold: unhealthyThreshold,
        ServerGroupId: serverGroupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyListenerRequestWithoutParam() *ModifyListenerRequest {

    return &ModifyListenerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/{listenerId}:modifyListenerAttributes",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域(Required) */
func (r *ModifyListenerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param listenerId: 监听器ID(Required) */
func (r *ModifyListenerRequest) SetListenerId(listenerId string) {
    r.ListenerId = listenerId
}

/* param algorithm: 调度算法(Optional) */
func (r *ModifyListenerRequest) SetAlgorithm(algorithm string) {
    r.Algorithm = &algorithm
}

/* param stickySession: 会话保持(Optional) */
func (r *ModifyListenerRequest) SetStickySession(stickySession string) {
    r.StickySession = &stickySession
}

/* param realIp: 是否获取真实ip，取值范围on|off(Optional) */
func (r *ModifyListenerRequest) SetRealIp(realIp string) {
    r.RealIp = &realIp
}

/* param name: 名称(Optional) */
func (r *ModifyListenerRequest) SetName(name string) {
    r.Name = &name
}

/* param description: 描述(Optional) */
func (r *ModifyListenerRequest) SetDescription(description string) {
    r.Description = &description
}

/* param healthCheck: 健康检查(Optional) */
func (r *ModifyListenerRequest) SetHealthCheck(healthCheck string) {
    r.HealthCheck = &healthCheck
}

/* param healthCheckTimeout: 健康检查响应的最大超时时间(Optional) */
func (r *ModifyListenerRequest) SetHealthCheckTimeout(healthCheckTimeout int) {
    r.HealthCheckTimeout = &healthCheckTimeout
}

/* param healthCheckInterval: 健康检查响应的最大间隔时间(Optional) */
func (r *ModifyListenerRequest) SetHealthCheckInterval(healthCheckInterval int) {
    r.HealthCheckInterval = &healthCheckInterval
}

/* param healthyThreshold: 健康检查结果为success的阈值(Optional) */
func (r *ModifyListenerRequest) SetHealthyThreshold(healthyThreshold int) {
    r.HealthyThreshold = &healthyThreshold
}

/* param unhealthyThreshold: 健康检查结果为fail的阈值(Optional) */
func (r *ModifyListenerRequest) SetUnhealthyThreshold(unhealthyThreshold int) {
    r.UnhealthyThreshold = &unhealthyThreshold
}

/* param serverGroupId: 服务器组id(Optional) */
func (r *ModifyListenerRequest) SetServerGroupId(serverGroupId string) {
    r.ServerGroupId = &serverGroupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyListenerRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyListenerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyListenerResult `json:"result"`
}

type ModifyListenerResult struct {
    Listener cps.Listener `json:"listener"`
}