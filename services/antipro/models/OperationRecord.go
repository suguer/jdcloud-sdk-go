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


type OperationRecord struct {

    /* 操作时间 (Optional) */
    Time string `json:"time"`

    /* 防护包名称 (Optional) */
    Name string `json:"name"`

    /* 操作类型.<br>- 1: 套餐变更<br>- 2: 防护规则变更<br>- 3: 防护对象变更<br>- 4: IP 地址变更<br>- 5: 防护包名称变更<br>- 6: IP地址库变更<br>- 7: 端口库变更<br>- 8: 访问控制规则变更 (Optional) */
    Action int `json:"action"`

    /* 操作详情 (Optional) */
    Info string `json:"info"`

    /* 操作人 (Optional) */
    Operator string `json:"operator"`
}
