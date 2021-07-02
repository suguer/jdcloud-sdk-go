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


type SceneResult struct {

    /* 图片名称(或图片标识) (Optional) */
    Name string `json:"name"`

    /* 本次请求数据标识，可以根据该标识查询数据最新结果 (Optional) */
    TaskId string `json:"taskId"`

    /* 场景检测详细信息 (Optional) */
    Details []SceneResultDetail `json:"details"`
}
