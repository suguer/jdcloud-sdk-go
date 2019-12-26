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


type DistributorPolicy struct {

    /* ID (Optional) */
    Id int `json:"id"`

    /* 服务商ID (Optional) */
    DistributorId string `json:"distributorId"`

    /* 服务商名称 (Optional) */
    DistributorName string `json:"distributorName"`

    /* 返还政策ID (Optional) */
    ReturnPolicyId int `json:"returnPolicyId"`

    /* 返还政策名称 (Optional) */
    ReturnPolicyName string `json:"returnPolicyName"`

    /* 状态 (Optional) */
    Status int `json:"status"`

    /* 说明 (Optional) */
    Remark string `json:"remark"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 创建人 (Optional) */
    CreateUser string `json:"createUser"`

    /* 修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 修改人 (Optional) */
    UpdateUser string `json:"updateUser"`

    /* 是否删除0未删除,1已删除 (Optional) */
    Yn int `json:"yn"`
}
