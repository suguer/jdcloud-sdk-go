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


type ReturnBillDetailResult struct {

    /* ID (Optional) */
    Id int `json:"id"`

    /* 服务商ID (Optional) */
    DistributorId string `json:"distributorId"`

    /* 渠道商类型 (Optional) */
    DistributorType int `json:"distributorType"`

    /* 生成时间 (Optional) */
    GenerateTime string `json:"generateTime"`

    /* 返还类型 (Optional) */
    ReturnType int `json:"returnType"`

    /* 返还项目ID (Optional) */
    ItemId int `json:"itemId"`

    /* 返还项目名称 (Optional) */
    ItemName string `json:"itemName"`

    /* 返还依据类型 (Optional) */
    ReturnRuleType int `json:"returnRuleType"`

    /* 产品类型 (Optional) */
    ProductType int `json:"productType"`

    /* 产品ID (Optional) */
    ProductId string `json:"productId"`

    /* 产品名称 (Optional) */
    ProductName string `json:"productName"`

    /*  (Optional) */
    ReturnBillDetailProductList []ReturnBillDetailProduct `json:"returnBillDetailProductList"`

    /* 账单ID (Optional) */
    AccountingBillId int `json:"accountingBillId"`

    /* 订单号 (Optional) */
    MainTransactionNo string `json:"mainTransactionNo"`

    /* 周期类型 (Optional) */
    CircleType int `json:"circleType"`

    /* 周期名称 (Optional) */
    CircleName int `json:"circleName"`

    /* 周期值 (Optional) */
    CircleValue int `json:"circleValue"`

    /* 周期值名称 (Optional) */
    CircleValueName string `json:"circleValueName"`

    /* 开始时间 (Optional) */
    CircleBegin string `json:"circleBegin"`

    /* 结束时间 (Optional) */
    CircleEnd string `json:"circleEnd"`

    /* 返还金额 (Optional) */
    ReturnMount int `json:"returnMount"`

    /* 业绩金额 (Optional) */
    Amount int `json:"amount"`

    /* 条件值 (Optional) */
    ConditionValue string `json:"conditionValue"`

    /* 条件值名称 (Optional) */
    ConditionValueName string `json:"conditionValueName"`

    /* 是否返还标识 (Optional) */
    ReturnFlag int `json:"returnFlag"`

    /* 返还比例 (Optional) */
    ReturnRatio int `json:"returnRatio"`

    /* 返还比例名称 (Optional) */
    ReturnRatioName string `json:"returnRatioName"`

    /* 返还单号 (Optional) */
    ReturnOrderId string `json:"returnOrderId"`

    /* 部门ID (Optional) */
    DeptId int `json:"deptId"`

    /* 部门名称 (Optional) */
    DeptName string `json:"deptName"`

    /* 状态 (Optional) */
    Status int `json:"status"`

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
