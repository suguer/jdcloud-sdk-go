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

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type Container struct {

    /* 容器ID (Optional) */
    ContainerId string `json:"containerId"`

    /* 容器状态 (Optional) */
    Status string `json:"status"`

    /* 实例类型 (Optional) */
    InstanceType string `json:"instanceType"`

    /* 可用区 (Optional) */
    Az string `json:"az"`

    /* 容器名称 (Optional) */
    Name string `json:"name"`

    /* 域名和IP映射的信息 (Optional) */
    HostAliases []HostAlias `json:"hostAliases"`

    /* 主机名 (Optional) */
    Hostname string `json:"hostname"`

    /* 容器执行命令 (Optional) */
    Command []string `json:"command"`

    /* 容器执行命令的参数 (Optional) */
    Args []string `json:"args"`

    /* 动态指定的容器执行的环境变量 (Optional) */
    Envs []EnvVar `json:"envs"`

    /* 镜像名称 (Optional) */
    Image string `json:"image"`

    /* 镜像仓库认证信息名称 (Optional) */
    Secret string `json:"secret"`

    /* 容器是否分配tty (Optional) */
    Tty bool `json:"tty"`

    /* 容器的工作目录 (Optional) */
    WorkingDir string `json:"workingDir"`

    /* 根Volume信息 (Optional) */
    RootVolume VolumeMount `json:"rootVolume"`

    /* 挂载的数据Volume信息 (Optional) */
    DataVolumes []VolumeMount `json:"dataVolumes"`

    /* 主网卡所属VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 主网卡所属子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 主网卡主IP地址 (Optional) */
    PrivateIpAddress string `json:"privateIpAddress"`

    /* 主网卡主IP绑定弹性IP的ID (Optional) */
    ElasticIpId string `json:"elasticIpId"`

    /* 主网卡主IP绑定弹性IP的地址 (Optional) */
    ElasticIpAddress string `json:"elasticIpAddress"`

    /* 主网卡信息 (Optional) */
    PrimaryNetworkInterface InstanceNetworkInterfaceAttachment `json:"primaryNetworkInterface"`

    /* 弹性网卡信息 (Optional) */
    SecondaryNetworkInterfaces []InstanceNetworkInterfaceAttachment `json:"secondaryNetworkInterfaces"`

    /* 容器日志配置信息 (Optional) */
    LogConfiguration LogConfiguration `json:"logConfiguration"`

    /*  (Optional) */
    Tags []Tag `json:"tags"`

    /* 计费配置信息 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* 创建时间 (Optional) */
    LaunchTime string `json:"launchTime"`

    /* 容器终止原因 (Optional) */
    Reason string `json:"reason"`

    /* 容器描述 (Optional) */
    Description string `json:"description"`
}
