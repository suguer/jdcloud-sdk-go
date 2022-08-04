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


type FileToPath struct {

    /* 键名称，不能重复，最大长度不超过128（字母、数字、-、_和.）  */
    Key string `json:"key"`

    /* 内容（base64） 每个value长度上限为32KB，整个data的长度不能超过1M;  */
    Value string `json:"value"`
}
