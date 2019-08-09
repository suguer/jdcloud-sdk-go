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


type UpdateMonitor struct {

    /* 连续几次触发报警  */
    AlarmLimit int `json:"alarmLimit"`

    /* 监控项ID  */
    Id int `json:"id"`

    /* 备用地址1  */
    IpBackup01 string `json:"ipBackup01"`

    /* 备用地址2  */
    IpBackup02 string `json:"ipBackup02"`

    /* 备用地址列表，存在该参数时，可不填写参数备用地址1、备用地址2 (Optional) */
    BackupAddressList []string `json:"backupAddressList"`

    /* 监控状况 开启监控:2，暂停监控:4  */
    MonitorEnable int `json:"monitorEnable"`

    /* 监控频率，单位秒  */
    MonitorFreq int `json:"monitorFreq"`

    /* 监控端口  */
    MonitorPort int `json:"monitorPort"`

    /* 不做任何修改0，强制暂停解析记录1，自动切换到备用地址2  */
    MonitorRule int `json:"monitorRule"`

    /* 监控路径  */
    MonitorUri string `json:"monitorUri"`

    /* 不发送邮件:0， 发送邮件:1  */
    NotifyEmailEnable int `json:"notifyEmailEnable"`

    /* 不发送通知栏:0， 发送通知栏:1  */
    NotifyMsgBarEnable int `json:"notifyMsgBarEnable"`

    /* 不发送短信:0， 发送短信:1 (Optional) */
    NotifySmsEnable *int `json:"notifySmsEnable"`

    /* https 0，https 1  */
    Protocol int `json:"protocol"`

    /* 0自动恢复 1手动恢复  */
    StopRecoverRule int `json:"stopRecoverRule"`

    /* 0自动恢复至主host 1手动恢复至主host  */
    SwitchRecoverRule int `json:"switchRecoverRule"`
}
