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
    bastion "github.com/jdcloud-api/jdcloud-sdk-go/services/bastion/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribePinsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* valid - 有效用户 1:0
upgrade - 升配标志 1:0
 (Optional) */
    Filters []common.Filter `json:"filters"`

    /* createTime - 创建时间,asc（正序），desc（倒序）
 (Optional) */
    Sorts []common.Sort `json:"sorts"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePinsRequest(
    regionId string,
) *DescribePinsRequest {

	return &DescribePinsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pins",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 * param filters: valid - 有效用户 1:0
upgrade - 升配标志 1:0
 (Optional)
 * param sorts: createTime - 创建时间,asc（正序），desc（倒序）
 (Optional)
 */
func NewDescribePinsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
    sorts []common.Sort,
) *DescribePinsRequest {

    return &DescribePinsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pins",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
        Sorts: sorts,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePinsRequestWithoutParam() *DescribePinsRequest {

    return &DescribePinsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pins",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribePinsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribePinsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小；默认为20；取值范围[10, 100](Optional) */
func (r *DescribePinsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param filters: valid - 有效用户 1:0
upgrade - 升配标志 1:0
(Optional) */
func (r *DescribePinsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}
/* param sorts: createTime - 创建时间,asc（正序），desc（倒序）
(Optional) */
func (r *DescribePinsRequest) SetSorts(sorts []common.Sort) {
    r.Sorts = sorts
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePinsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePinsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePinsResult `json:"result"`
}

type DescribePinsResult struct {
    Pins []bastion.UserPin `json:"pins"`
}