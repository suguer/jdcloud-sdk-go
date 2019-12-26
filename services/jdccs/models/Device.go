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


type Device struct {

    /* 机房英文标识 (Optional) */
    Idc string `json:"idc"`

    /* 机房名称 (Optional) */
    IdcName string `json:"idcName"`

    /* 设备Id (Optional) */
    DeviceId string `json:"deviceId"`

    /* 设备SN号 (Optional) */
    SnNo string `json:"snNo"`

    /* 机柜编码 (Optional) */
    CabinetNo string `json:"cabinetNo"`

    /* 所在U位 (Optional) */
    RackUIndex string `json:"rackUIndex"`

    /* U数（U） (Optional) */
    UNum int `json:"uNum"`

    /* 品牌 (Optional) */
    Brand string `json:"brand"`

    /* 型号 (Optional) */
    Model string `json:"model"`

    /* 系统IP (Optional) */
    SysIp string `json:"sysIp"`

    /* 管理IP (Optional) */
    ManageIp string `json:"manageIp"`

    /* 设备类型 server:服务器 network:网络设备 storage:存储设备 other:其他设备 (Optional) */
    DeviceType string `json:"deviceType"`

    /* 资产归属 own:自备 lease:租赁 (Optional) */
    AssetBelong string `json:"assetBelong"`

    /* 资产状态 launched:已上架 opened:已开通 canceling:退订中 operating:操作中 modifing:变更中 (Optional) */
    AssetStatus string `json:"assetStatus"`

    /* 开通时间，遵循ISO8601标准，使用UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    DeviceOpenTime string `json:"deviceOpenTime"`

    /* CPU (Optional) */
    CpuCore string `json:"cpuCore"`

    /* 内存 (Optional) */
    Memory string `json:"memory"`

    /* 磁盘 (Optional) */
    Disk string `json:"disk"`
}
