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
    vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
)

type DescribeImagesRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 镜像来源，如果没有指定 `ids` 参数，此参数必传。取值范围：
`public`：官方镜像。
`thirdparty`：镜像市场镜像。
`private`：用户自己的私有镜像。
`shared`：其他用户分享的镜像。
`community`：社区镜像。
 (Optional) */
    ImageSource *string `json:"imageSource"`

    /* 查询已经下线的镜像时使用。
只有查询 `官方镜像` 或者 `镜像市场镜像` 时，此参数才有意义，其它情况下此参数无效。
指定 `ids` 查询时，此参数无效。
 (Optional) */
    Offline *bool `json:"offline"`

    /* 根据镜像的操作系统发行版查询。
取值范围：`Ubuntu、CentOS、Windows Server`。
 (Optional) */
    Platform *string `json:"platform"`

    /* 指定镜像ID查询，如果指定了此参数，其它参数可以不传。
 (Optional) */
    Ids []string `json:"ids"`

    /* 根据镜像名称模糊查询。 (Optional) */
    ImageName *string `json:"imageName"`

    /* 根据镜像支持的系统盘类型查询。支持范围：`localDisk` 本地系统盘镜像；`cloudDisk` 云盘系统盘镜像。 (Optional) */
    RootDeviceType *string `json:"rootDeviceType"`

    /* 根据镜像的使用权限查询，可选参数，仅当 `imageSource` 为 `private` 时有效。取值范围：
`all`：没有限制，所有人均可以使用。
`specifiedUsers`：只有共享用户可以使用。
`ownerOnly`：镜像拥有者自己可以使用。
 (Optional) */
    LaunchPermission *string `json:"launchPermission"`

    /* 根据镜像状态查询。参考 [镜像状态](https://docs.jdcloud.com/virtual-machines/api/image_status) (Optional) */
    Status *string `json:"status"`

    /* 已废弃。 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 页码；默认为1。 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；<br>默认为20；取值范围[10, 100]。 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域ID。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeImagesRequest(
    regionId string,
) *DescribeImagesRequest {

	return &DescribeImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param imageSource: 镜像来源，如果没有指定 `ids` 参数，此参数必传。取值范围：
`public`：官方镜像。
`thirdparty`：镜像市场镜像。
`private`：用户自己的私有镜像。
`shared`：其他用户分享的镜像。
`community`：社区镜像。
 (Optional)
 * param offline: 查询已经下线的镜像时使用。
只有查询 `官方镜像` 或者 `镜像市场镜像` 时，此参数才有意义，其它情况下此参数无效。
指定 `ids` 查询时，此参数无效。
 (Optional)
 * param platform: 根据镜像的操作系统发行版查询。
取值范围：`Ubuntu、CentOS、Windows Server`。
 (Optional)
 * param ids: 指定镜像ID查询，如果指定了此参数，其它参数可以不传。
 (Optional)
 * param imageName: 根据镜像名称模糊查询。 (Optional)
 * param rootDeviceType: 根据镜像支持的系统盘类型查询。支持范围：`localDisk` 本地系统盘镜像；`cloudDisk` 云盘系统盘镜像。 (Optional)
 * param launchPermission: 根据镜像的使用权限查询，可选参数，仅当 `imageSource` 为 `private` 时有效。取值范围：
`all`：没有限制，所有人均可以使用。
`specifiedUsers`：只有共享用户可以使用。
`ownerOnly`：镜像拥有者自己可以使用。
 (Optional)
 * param status: 根据镜像状态查询。参考 [镜像状态](https://docs.jdcloud.com/virtual-machines/api/image_status) (Optional)
 * param serviceCode: 已废弃。 (Optional)
 * param pageNumber: 页码；默认为1。 (Optional)
 * param pageSize: 分页大小；<br>默认为20；取值范围[10, 100]。 (Optional)
 */
func NewDescribeImagesRequestWithAllParams(
    regionId string,
    imageSource *string,
    offline *bool,
    platform *string,
    ids []string,
    imageName *string,
    rootDeviceType *string,
    launchPermission *string,
    status *string,
    serviceCode *string,
    pageNumber *int,
    pageSize *int,
) *DescribeImagesRequest {

    return &DescribeImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ImageSource: imageSource,
        Offline: offline,
        Platform: platform,
        Ids: ids,
        ImageName: imageName,
        RootDeviceType: rootDeviceType,
        LaunchPermission: launchPermission,
        Status: status,
        ServiceCode: serviceCode,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeImagesRequestWithoutParam() *DescribeImagesRequest {

    return &DescribeImagesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *DescribeImagesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param imageSource: 镜像来源，如果没有指定 `ids` 参数，此参数必传。取值范围：
`public`：官方镜像。
`thirdparty`：镜像市场镜像。
`private`：用户自己的私有镜像。
`shared`：其他用户分享的镜像。
`community`：社区镜像。
(Optional) */
func (r *DescribeImagesRequest) SetImageSource(imageSource string) {
    r.ImageSource = &imageSource
}

/* param offline: 查询已经下线的镜像时使用。
只有查询 `官方镜像` 或者 `镜像市场镜像` 时，此参数才有意义，其它情况下此参数无效。
指定 `ids` 查询时，此参数无效。
(Optional) */
func (r *DescribeImagesRequest) SetOffline(offline bool) {
    r.Offline = &offline
}

/* param platform: 根据镜像的操作系统发行版查询。
取值范围：`Ubuntu、CentOS、Windows Server`。
(Optional) */
func (r *DescribeImagesRequest) SetPlatform(platform string) {
    r.Platform = &platform
}

/* param ids: 指定镜像ID查询，如果指定了此参数，其它参数可以不传。
(Optional) */
func (r *DescribeImagesRequest) SetIds(ids []string) {
    r.Ids = ids
}

/* param imageName: 根据镜像名称模糊查询。(Optional) */
func (r *DescribeImagesRequest) SetImageName(imageName string) {
    r.ImageName = &imageName
}

/* param rootDeviceType: 根据镜像支持的系统盘类型查询。支持范围：`localDisk` 本地系统盘镜像；`cloudDisk` 云盘系统盘镜像。(Optional) */
func (r *DescribeImagesRequest) SetRootDeviceType(rootDeviceType string) {
    r.RootDeviceType = &rootDeviceType
}

/* param launchPermission: 根据镜像的使用权限查询，可选参数，仅当 `imageSource` 为 `private` 时有效。取值范围：
`all`：没有限制，所有人均可以使用。
`specifiedUsers`：只有共享用户可以使用。
`ownerOnly`：镜像拥有者自己可以使用。
(Optional) */
func (r *DescribeImagesRequest) SetLaunchPermission(launchPermission string) {
    r.LaunchPermission = &launchPermission
}

/* param status: 根据镜像状态查询。参考 [镜像状态](https://docs.jdcloud.com/virtual-machines/api/image_status)(Optional) */
func (r *DescribeImagesRequest) SetStatus(status string) {
    r.Status = &status
}

/* param serviceCode: 已废弃。(Optional) */
func (r *DescribeImagesRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param pageNumber: 页码；默认为1。(Optional) */
func (r *DescribeImagesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；<br>默认为20；取值范围[10, 100]。(Optional) */
func (r *DescribeImagesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeImagesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeImagesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeImagesResult `json:"result"`
}

type DescribeImagesResult struct {
    Images []vm.Image `json:"images"`
    TotalCount int `json:"totalCount"`
}