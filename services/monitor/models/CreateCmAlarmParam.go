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


type CreateCmAlarmParam struct {

    /* 统计方法，必须与定义的metric一致，可选值列表：avg,max,sum,min  */
    Calculation string `json:"calculation"`

    /* 报警规则通知的联系组，必须在控制台上已创建，例如" ['联系组1','联系组2']" (Optional) */
    ContactGroups []string `json:"contactGroups"`

    /* 报警规则通知的联系人，必须在控制台上已创建，例如 [“联系人1”,”联系人2”] (Optional) */
    ContactPersons []string `json:"contactPersons"`

    /* 取样频次 (Optional) */
    DownSample string `json:"downSample"`

    /* 根据产品线查询可用监控项列表 接口 返回的Metric字段  */
    MetricUID string `json:"metricUID"`

    /* 命名空间  */
    NamespaceUID string `json:"namespaceUID"`

    /* 通知周期 单位：小时 (Optional) */
    NoticePeriod int64 `json:"noticePeriod"`

    /* 报警规则对应实例列表，每次最多100个，例如"['resourceId1','resourceId2']"  */
    ObjUIDs []string `json:"objUIDs"`

    /* 报警比较符，只能为以下几种<=,<,>,>=,==,!=  */
    Operation string `json:"operation"`

    /* 查询指标的周期，单位为分钟,目前支持的取值：2，5，15，30，60  */
    Period int64 `json:"period"`

    /* 规则名称，最大长度42个字符，只允许中英文、数字、''-''和"_" (Optional) */
    RuleName string `json:"ruleName"`

    /* 报警阈值，目前只开放数值类型功能  */
    Threshold float64 `json:"threshold"`

    /* 连续探测几次都满足阈值条件时报警，可选值:1,2,3,5  */
    Times int64 `json:"times"`
}
