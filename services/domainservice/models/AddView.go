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


type AddView struct {

    /* 主域名  */
    DomainName string `json:"domainName"`

    /* 自定义线路名称, 最多64个字节，允许：数字、字母、下划线，-，中文  */
    ViewName string `json:"viewName"`

    /* 用户输入的此线路的ip段。  
IPv4地址段支持1.2.3.4-5.6.7.8和1.2.3.4/16两种格式。    
IPv6地址段支持CIDR格式，例如：11:22:33:44:55::99/64
  */
    IpRanges []string `json:"ipRanges"`
}
