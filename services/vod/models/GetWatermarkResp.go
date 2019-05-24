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

package models


type GetWatermarkResp struct {

    /* 水印ID (Optional) */
    Id int64 `json:"id"`

    /* 水印名称 (Optional) */
    Name string `json:"name"`

    /* 图片地址 (Optional) */
    ImgUrl string `json:"imgUrl"`

    /* 宽度 (Optional) */
    Width int `json:"width"`

    /* 高度 (Optional) */
    Height int `json:"height"`

    /* 水印位置 (Optional) */
    Position string `json:"position"`

    /* 偏移单位 (Optional) */
    Unit string `json:"unit"`

    /* 水平偏移 (Optional) */
    OffsetX int `json:"offsetX"`

    /* 竖直偏移 (Optional) */
    OffsetY int `json:"offsetY"`

    /*  (Optional) */
    CreateTime string `json:"createTime"`

    /*  (Optional) */
    UpdateTime string `json:"updateTime"`
}
