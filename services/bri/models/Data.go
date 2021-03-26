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


type Data struct {

    /* 累计付费用户 (Optional) */
    PaidUsersNumber int `json:"paidUsersNumber"`

    /* 累计体验用户 (Optional) */
    FreeUsersNumber int `json:"freeUsersNumber"`

    /* 累计无效用户 (Optional) */
    InvalidUsersNumber int `json:"invalidUsersNumber"`

    /* 资源包总数 (Optional) */
    TotalResourcePack int `json:"totalResourcePack"`

    /* 有效资源包 (Optional) */
    ValidResourcePack int `json:"validResourcePack"`

    /* 过期资源包 (Optional) */
    ExpireResourcePack int `json:"expireResourcePack"`

    /* 订单总额 (Optional) */
    TotalOrder int `json:"totalOrder"`

    /* 收入总额 (Optional) */
    TotalIncome int `json:"totalIncome"`
}
