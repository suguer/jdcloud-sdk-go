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


type BackInstanceVoucherVo struct {

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 账户名 (Optional) */
    Account string `json:"account"`

    /* 实例券id (Optional) */
    InstanceVoucherId string `json:"instanceVoucherId"`

    /* 实例券类型 (Optional) */
    InstanceVoucherType string `json:"instanceVoucherType"`

    /* 资源预留(1 有预留、2 无预留) (Optional) */
    IsReserved int `json:"isReserved"`

    /* 地域 (Optional) */
    Region string `json:"region"`

    /* 展示说明 (Optional) */
    Tips string `json:"tips"`

    /* 展示实例族/规格 (Optional) */
    Label string `json:"label"`

    /* 数量 (Optional) */
    Count int `json:"count"`

    /* 显示单位 (Optional) */
    Unit string `json:"unit"`

    /* 生效时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 失效时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 状态(1：正常: 2：已退订 -1:失效) (Optional) */
    Status int `json:"status"`
}
