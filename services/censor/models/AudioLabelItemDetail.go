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


type AudioLabelItemDetail struct {

    /* 命中检测器类型：10：用户名单 11:IP名单 12:设备名单 30:敏感词 110:IP地区限制 130:声纹检测 (Optional) */
    HitType int `json:"hitType"`

    /* 线索信息，用于定位文本中有问题的部分，辅助人工审核 (Optional) */
    Hint []AudioHint `json:"hint"`

    /* 线索详细信息自定义 (Optional) */
    HitInfos []AudioHintInfo `json:"hitInfos"`
}
