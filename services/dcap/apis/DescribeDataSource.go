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

type DescribeDataSourceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 数据源 ID  */
    DataSourceId string `json:"dataSourceId"`
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDataSourceRequest(
    regionId string,
    dataSourceId string,
) *DescribeDataSourceRequest {

	return &DescribeDataSourceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dataSources/{dataSourceId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DataSourceId: dataSourceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 */
func NewDescribeDataSourceRequestWithAllParams(
    regionId string,
    dataSourceId string,
) *DescribeDataSourceRequest {

    return &DescribeDataSourceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDataSourceRequestWithoutParam() *DescribeDataSourceRequest {

    return &DescribeDataSourceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeDataSourceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源 ID(Required) */
func (r *DescribeDataSourceRequest) SetDataSourceId(dataSourceId string) {
    r.DataSourceId = dataSourceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDataSourceRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDataSourceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDataSourceResult `json:"result"`
}

type DescribeDataSourceResult struct {
    DataSourceId string `json:"dataSourceId"`
    DataSourceType int `json:"dataSourceType"`
    DataSourceName string `json:"dataSourceName"`
    DataSourceAddr string `json:"dataSourceAddr"`
    DataSourcePort int `json:"dataSourcePort"`
    DataSourceDbName string `json:"dataSourceDbName"`
    Region string `json:"region"`
    VpcId string `json:"vpcId"`
    SubnetId string `json:"subnetId"`
    ProtectStatus bool `json:"protectStatus"`
    KmsKeyId string `json:"kmsKeyId"`
    KeyCipher string `json:"keyCipher"`
    EncryptAlgo string `json:"encryptAlgo"`
    IndexAlgo string `json:"indexAlgo"`
}