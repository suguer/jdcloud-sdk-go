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

type DescribeImageRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 镜像ID。  */
    ImageId string `json:"imageId"`

}

/*
 * param regionId: 地域ID。 (Required)
 * param imageId: 镜像ID。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeImageRequest(
    regionId string,
    imageId string,
) *DescribeImageRequest {

	return &DescribeImageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images/{imageId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ImageId: imageId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param imageId: 镜像ID。 (Required)
 */
func NewDescribeImageRequestWithAllParams(
    regionId string,
    imageId string,
) *DescribeImageRequest {

    return &DescribeImageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images/{imageId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ImageId: imageId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeImageRequestWithoutParam() *DescribeImageRequest {

    return &DescribeImageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images/{imageId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *DescribeImageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param imageId: 镜像ID。(Required) */
func (r *DescribeImageRequest) SetImageId(imageId string) {
    r.ImageId = imageId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeImageRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeImageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeImageResult `json:"result"`
}

type DescribeImageResult struct {
    Image vm.Image `json:"image"`
}