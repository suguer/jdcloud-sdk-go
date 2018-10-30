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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    nc "github.com/jdcloud-api/jdcloud-sdk-go/services/nc/apis"
    "encoding/json"
    "errors"
)

type NcClient struct {
    core.JDCloudClient
}

func NewNcClient(credential *core.Credential) *NcClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("nc.jdcloud-api.com")

    return &NcClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "nc",
            Revision:    "0.2.5",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *NcClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *NcClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 批量查询原生容器的详细信息<br>
此接口支持分页查询，默认每页20条。
 */
func (c *NcClient) DescribeContainers(request *nc.DescribeContainersRequest) (*nc.DescribeContainersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.DescribeContainersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建一台或多台指定配置容器。
- 创建容器需要通过实名认证
- 镜像
    - 容器的镜像通过镜像名称来确定
    - nginx:tag 或 mysql/mysql-server:tag 这样命名的镜像表示 docker hub 官方镜像
    - container-registry/image:tag 这样命名的镜像表示私有仓储的镜像
    - 私有仓储必须兼容 docker registry 认证机制，并通过 secret 来保存机密信息
- hostname 规范
    - 支持两种方式：以标签方式书写或以完整主机名方式书写
    - 标签规范
        - 0-9，a-z(不分大小写)和 -（减号），其他的都是无效的字符串
        - 不能以减号开始，也不能以减号结尾
        - 最小1个字符，最大63个字符
    - 完整的主机名由一系列标签与点连接组成
        - 标签与标签之间使用“.”(点)进行连接
        - 不能以“.”(点)开始，也不能以“.”(点)结尾
        - 整个主机名（包括标签以及分隔点“.”）最多有63个ASCII字符
- 网络配置
    - 指定主网卡配置信息
        - 必须指定一个子网
        - 一台云主机创建时必须指定一个安全组，至多指定 5 个安全组
        - 可以指定 elasticIp 规格来约束创建的弹性 IP，带宽取值范围 [1-200]Mbps，步进 1Mbps
        - 可以指定网卡的主 IP(primaryIpAddress)，该 IP 需要在子网 IP 范围内且未被占用，指定子网 IP 时 maxCount 只能为1
        - 安全组 securityGroup 需与子网 Subnet 在同一个私有网络 VPC 内
        - 主网卡 deviceIndex 设置为 1
- 存储
    - volume 分为 root volume 和 data volume，root volume 的挂载目录是 /，data volume 的挂载目录可以随意指定
    - volume 的底层存储介质当前只支持 cloud 类别，也就是云硬盘
    - 系统盘
        - 云硬盘类型可以选择 ssd、premium-hdd
        - 磁盘大小
            - ssd：范围 [10, 100]GB，步长为 10G
            - premium-hdd：范围 [20, 1000]GB，步长为 10G
        - 自动删除
            - 云盘默认跟随容器实例自动删除，如果是包年包月的数据盘或共享型数据盘，此参数不生效
        - 可以选择已存在的云硬盘
    - 数据盘
        - 云硬盘类型可以选择 ssd、premium-hdd
        - 磁盘大小
            - ssd：范围[20,1000]GB，步长为10G
            - premium-hdd：范围[20,3000]GB，步长为10G
        - 自动删除
            - 默认自动删除
        - 可以选择已存在的云硬盘
        - 单个容器最多可以挂载 7 个 data volume
- 计费
  - 弹性IP的计费模式，如果选择按用量类型可以单独设置，其它计费模式都以主机为准
  - 云硬盘的计费模式以主机为准
- 容器日志
    - 默认在本地分配10MB的存储空间，自动 rotate
- 其他
    - 创建完成后，容器状态为running
    - maxCount 为最大努力，不保证一定能达到 maxCount
 */
func (c *NcClient) CreateContainers(request *nc.CreateContainersRequest) (*nc.CreateContainersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.CreateContainersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询单个容器日志
 */
func (c *NcClient) GetLogs(request *nc.GetLogsRequest) (*nc.GetLogsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.GetLogsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询单个 secret 详情
 */
func (c *NcClient) DescribeSecret(request *nc.DescribeSecretRequest) (*nc.DescribeSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.DescribeSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 修改容器的 名称 和 描述。
 */
func (c *NcClient) ModifyContainerAttribute(request *nc.ModifyContainerAttributeRequest) (*nc.ModifyContainerAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.ModifyContainerAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 容器状态必须为 stopped、running 或 error状态。 <br>
按量付费的实例，如不主动删除将一直运行，不再使用的实例，可通过本接口主动停用。<br>
只能支持主动删除按量计费类型的实例。包年包月过期的容器也可以删除，其它的情况还请发工单系统。计费状态异常的容器无法删除。
 */
func (c *NcClient) DeleteContainer(request *nc.DeleteContainerRequest) (*nc.DeleteContainerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.DeleteContainerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 停止处于运行状态的单个实例，处于任务执行中的容器无法启动。
 */
func (c *NcClient) StopContainer(request *nc.StopContainerRequest) (*nc.StopContainerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.StopContainerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 启动处于关闭状态的单个容器，处在任务执行中的容器无法启动。<br>
容器实例或其绑定的云盘已欠费时，容器将无法正常启动。<br>
 */
func (c *NcClient) StartContainer(request *nc.StartContainerRequest) (*nc.StartContainerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.StartContainerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 容器解绑公网 IP，解绑的是主网卡、主内网 IP 对应的弹性 IP.
 */
func (c *NcClient) DisassociateElasticIp(request *nc.DisassociateElasticIpRequest) (*nc.DisassociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.DisassociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询 secret 列表。<br> 
此接口支持分页查询，默认每页20条。
 */
func (c *NcClient) DescribeSecrets(request *nc.DescribeSecretsRequest) (*nc.DescribeSecretsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.DescribeSecretsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询资源的配额，支持：原生容器 pod 和 secret.
 */
func (c *NcClient) DescribeQuota(request *nc.DescribeQuotaRequest) (*nc.DescribeQuotaResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.DescribeQuotaResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询一台原生容器的详细信息
 */
func (c *NcClient) DescribeContainer(request *nc.DescribeContainerRequest) (*nc.DescribeContainerResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.DescribeContainerResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除单个 secret
 */
func (c *NcClient) DeleteSecret(request *nc.DeleteSecretRequest) (*nc.DeleteSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.DeleteSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建一个 secret，用于存放镜像仓库机密相关信息。
 */
func (c *NcClient) CreateSecret(request *nc.CreateSecretRequest) (*nc.CreateSecretResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.CreateSecretResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 容器绑定弹性公网 IP，绑定的是主网卡、主内网IP对应的弹性IP. <br>
一台云主机只能绑定一个弹性公网 IP(主网卡)，若主网卡已存在弹性公网IP，会返回错误。<br>
如果是黑名单中的用户，会返回错误。
 */
func (c *NcClient) AssociateElasticIp(request *nc.AssociateElasticIpRequest) (*nc.AssociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &nc.AssociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

