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


type QueryOption struct {

    /* 聚合方式，默认等于downSampleType(若downSampleType为last，AggrType取max)或avg，可选值参考:sum、avg、min、max (Optional) */
    AggrType string `json:"aggrType"`

    /* 监控指标数据的维度信息,根据tags来指定查询的监控数据维度，至少指定一个查询维度  */
    Dimensions []DimensionsParam `json:"dimensions"`

    /* 采样方式，默认等于aggrType或avg，可选值参考：sum、avg、last、min、max (Optional) */
    DownSampleType string `json:"downSampleType"`

    /* metric  */
    Metric string `json:"metric"`
}
