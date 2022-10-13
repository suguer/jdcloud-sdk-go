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


type InstanceVoucherDTO struct {

    /* 实例券ID (Optional) */
    InstanceVoucherId string `json:"instanceVoucherId"`

    /* 实例券类型(对应iaas侧resourceType) (Optional) */
    ReservedInstanceType string `json:"reservedInstanceType"`

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 产品 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 是否有预留(1 有预留、2 无预留,对应iaas侧reservedType 预留:reserved, 不预留:unReserved) (Optional) */
    IsReserved int `json:"isReserved"`

    /* 状态(1：正常: 2：已退订,-1:失效) (Optional) */
    Status int `json:"status"`

    /* 生效时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 失效时间 (Optional) */
    EndTime string `json:"endTime"`
}
