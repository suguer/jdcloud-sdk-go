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


type RouteTable struct {

    /* 路由表ID (Optional) */
    RouteTableId string `json:"routeTableId"`

    /* 地域 (Optional) */
    Region string `json:"region"`

    /* 私有网络ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 名称 (Optional) */
    Name string `json:"name"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 路由规则 (Optional) */
    Routes []Route `json:"routes"`
}
