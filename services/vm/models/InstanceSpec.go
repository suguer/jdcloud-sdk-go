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

import vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"
import disk "github.com/jdcloud-api/jdcloud-sdk-go/services/disk/models"

type InstanceSpec struct {

    /* 高可用组Id。指定了此参数后，只能通过高可用组关联的实例模板创建虚机，并且实例模板中的参数不可覆盖替换。实例模板以外的参数还可以指定。 (Optional) */
    AgId *string `json:"agId"`

    /* 实例模板id，如果没有使用高可用组，那么对于实例模板中没有的信息，需要使用创建虚机的参数进行补充，或者选择覆盖启动模板中的参数。 (Optional) */
    InstanceTemplateId *string `json:"instanceTemplateId"`

    /* 云主机所属的可用区。 (Optional) */
    Az *string `json:"az"`

    /* 实例规格。可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeinstancetypes">DescribeInstanceTypes</a>接口获得指定地域或可用区的规格信息。 (Optional) */
    InstanceType *string `json:"instanceType"`

    /* 镜像ID。可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeimages">DescribeImages</a>接口获得指定地域的镜像信息。 (Optional) */
    ImageId *string `json:"imageId"`

    /* 云主机名称，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。  */
    Name string `json:"name"`

    /* 密码，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional) */
    Password *string `json:"password"`

    /* 密钥对名称，当前只支持传入一个。 (Optional) */
    KeyNames []string `json:"keyNames"`

    /* 主网卡主IP关联的弹性IP规格 (Optional) */
    ElasticIp *vpc.ElasticIpSpec `json:"elasticIp"`

    /* 主网卡配置信息 (Optional) */
    PrimaryNetworkInterface *InstanceNetworkInterfaceAttachmentSpec `json:"primaryNetworkInterface"`

    /* 系统盘配置信息 (Optional) */
    SystemDisk *InstanceDiskAttachmentSpec `json:"systemDisk"`

    /* 数据盘配置信息，本地盘(local类型)做系统盘的云主机可挂载8块数据盘，云硬盘(cloud类型)做系统盘的云主机可挂载7块数据盘。 (Optional) */
    DataDisks []InstanceDiskAttachmentSpec `json:"dataDisks"`

    /* 计费配置
云主机不支持按用量方式计费，默认为按配置计费。
打包创建数据盘的情况下，数据盘的计费方式只能与云主机保持一致。
打包创建弹性公网IP的情况下，若公网IP的计费方式没有指定为按用量计费，那么公网IP计费方式只能与云主机保持一致。
 (Optional) */
    Charge *charge.ChargeSpec `json:"charge"`

    /* 元数据信息，目前只支持传入一个key为"launch-script"，表示首次启动脚本。value为base64格式，编码前数据不能大于16KB。
launch-script：linux系统支持bash和python，编码前须分别以 #!/bin/bash 和 #!/usr/bin/env python 作为内容首行;
launch-script：windows系统支持bat和powershell，编码前须分别以 <cmd></cmd> 和 <powershell></powershell> 作为内容首、尾行。
 (Optional) */
    Userdata []Userdata `json:"userdata"`

    /* 主机描述，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional) */
    Description *string `json:"description"`

    /* 不使用模板中的密码。
仅当不使用Ag，并且使用了模板，并且password参数为空时，此参数(值为true)生效。
若使用模板创建虚机时，又指定了password参数时，此参数无效，以新指定的为准。
 (Optional) */
    NoPassword *bool `json:"noPassword"`

    /* 不使用模板中的密钥。
仅当不使用Ag，并且使用了模板，并且keynames参数为空时，此参数(值为true)生效。
若使用模板创建虚机时，又指定了keynames参数时，此参数无效，以新指定的为准。
 (Optional) */
    NoKeyNames *bool `json:"noKeyNames"`

    /* 不使用模板中的弹性公网IP。
仅当不使用Ag，并且使用了模板，并且elasticIp参数为空时，此参数(值为true)生效。
若使用模板创建虚机时，又指定了elasticIp参数时，此参数无效，以新指定的为准。
 (Optional) */
    NoElasticIp *bool `json:"noElasticIp"`

    /* 用户普通标签集合 (Optional) */
    UserTags []disk.Tag `json:"userTags"`

    /* 关机模式，只支持云盘做系统盘的按配置计费云主机。keepCharging：关机后继续计费；stopCharging：关机后停止计费。 (Optional) */
    ChargeOnStopped *string `json:"chargeOnStopped"`
}
