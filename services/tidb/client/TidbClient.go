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
    tidb "github.com/jdcloud-api/jdcloud-sdk-go/services/tidb/apis"
    "encoding/json"
    "errors"
)

type TidbClient struct {
    core.JDCloudClient
}

func NewTidbClient(credential *core.Credential) *TidbClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("tidb.jdcloud-api.com")

    return &TidbClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "tidb",
            Revision:    "1.0.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *TidbClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *TidbClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *TidbClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 创建数据库账号，用户可以使用客户端，应用程序等通过该账号和密码登录数据库实例。 */
func (c *TidbClient) CreateAccount(request *tidb.CreateAccountRequest) (*tidb.CreateAccountResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.CreateAccountResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除白名单分组。 */
func (c *TidbClient) DeleteWhiteListGroup(request *tidb.DeleteWhiteListGroupRequest) (*tidb.DeleteWhiteListGroupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DeleteWhiteListGroupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 升级TiDB引擎版本，例如从4.0.6 升级到4.0.8. 目前支持小版本的升级，可升级到平台支持的最新的小版本 */
func (c *TidbClient) UpgradeEngineVersion(request *tidb.UpgradeEngineVersionRequest) (*tidb.UpgradeEngineVersionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.UpgradeEngineVersionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改实例名称，可支持中文，实例名的具体规则可参见帮助中心文档 */
func (c *TidbClient) ModifyInstanceName(request *tidb.ModifyInstanceNameRequest) (*tidb.ModifyInstanceNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.ModifyInstanceNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 校验需要导入的备份文件在OSS上是否存在，需要的读取权限是否具备 */
func (c *TidbClient) VerifyFilefromOSS(request *tidb.VerifyFilefromOSSRequest) (*tidb.VerifyFilefromOSSResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.VerifyFilefromOSSResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取可用区 */
func (c *TidbClient) DescribeAvailableZones(request *tidb.DescribeAvailableZonesRequest) (*tidb.DescribeAvailableZonesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeAvailableZonesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一个实例全量备份，可以对整个实例所有的数据库进行全量备份。同一时间点，只能有一个正在运行的备份任务 */
func (c *TidbClient) CreateBackup(request *tidb.CreateBackupRequest) (*tidb.CreateBackupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.CreateBackupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改TiDB实例备份策略。 */
func (c *TidbClient) ModifyBackupPolicy(request *tidb.ModifyBackupPolicyRequest) (*tidb.ModifyBackupPolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.ModifyBackupPolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看该实例下所有备份的详细信息 */
func (c *TidbClient) DescribeBackups(request *tidb.DescribeBackupsRequest) (*tidb.DescribeBackupsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeBackupsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询 TiDB 数据迁移任务的信息 */
func (c *TidbClient) DescribeDataMigration(request *tidb.DescribeDataMigrationRequest) (*tidb.DescribeDataMigrationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeDataMigrationResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除一个TiDB实例 */
func (c *TidbClient) DeleteInstance(request *tidb.DeleteInstanceRequest) (*tidb.DeleteInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DeleteInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改TiDB实例的配置参数。部分参数修改后，需要重启才能生效，具体可以参考PingCAP的相关文档 */
func (c *TidbClient) ModifyParameters(request *tidb.ModifyParametersRequest) (*tidb.ModifyParametersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.ModifyParametersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询 TiDB 数据迁移任务的信息 */
func (c *TidbClient) CreateDataMigration(request *tidb.CreateDataMigrationRequest) (*tidb.CreateDataMigrationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.CreateDataMigrationResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询实例列表 */
func (c *TidbClient) DescribeInstances(request *tidb.DescribeInstancesRequest) (*tidb.DescribeInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 增加白名单分组。 */
func (c *TidbClient) AddWhiteListGroup(request *tidb.AddWhiteListGroupRequest) (*tidb.AddWhiteListGroupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.AddWhiteListGroupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 增加实例的节点数量。 */
func (c *TidbClient) ModifyNodeNum(request *tidb.ModifyNodeNumRequest) (*tidb.ModifyNodeNumResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.ModifyNodeNumResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看TiDB实例的配置参数 */
func (c *TidbClient) DescribeParameters(request *tidb.DescribeParametersRequest) (*tidb.DescribeParametersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeParametersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看某个实例下的账号信息 */
func (c *TidbClient) DescribeAccounts(request *tidb.DescribeAccountsRequest) (*tidb.DescribeAccountsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeAccountsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据源实例全量备份创建一个新实例 */
func (c *TidbClient) CreateInstanceFromBackup(request *tidb.CreateInstanceFromBackupRequest) (*tidb.CreateInstanceFromBackupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.CreateInstanceFromBackupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改实例规格，包含节点的水平扩容与垂直扩容 */
func (c *TidbClient) ModifyInstanceSpec(request *tidb.ModifyInstanceSpecRequest) (*tidb.ModifyInstanceSpecResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.ModifyInstanceSpecResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看实例当前白名单。白名单是允许访问当前实例的IP/IP段列表，缺省情况下，白名单对本VPC开放。如果用户开启了外网访问的功能，还需要对外网的IP配置白名单。 */
func (c *TidbClient) DescribeWhiteList(request *tidb.DescribeWhiteListRequest) (*tidb.DescribeWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取TiDB产品提供的所有版本 */
func (c *TidbClient) DescribeVersions(request *tidb.DescribeVersionsRequest) (*tidb.DescribeVersionsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeVersionsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 规格获取接口 */
func (c *TidbClient) DescribeInstanceClasses(request *tidb.DescribeInstanceClassesRequest) (*tidb.DescribeInstanceClassesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeInstanceClassesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改允许访问实例的IP白名单。白名单是允许访问当前实例的IP/IP段列表，缺省情况下，白名单对本VPC开放。如果用户开启了外网访问的功能，还需要对外网的IP配置白名单。 */
func (c *TidbClient) ModifyWhiteList(request *tidb.ModifyWhiteListRequest) (*tidb.ModifyWhiteListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.ModifyWhiteListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建数据库账号，用户可以使用客户端，应用程序等通过该账号和密码登录RDS数据库实例。 */
func (c *TidbClient) ResetPassword(request *tidb.ResetPasswordRequest) (*tidb.ResetPasswordResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.ResetPasswordResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 重启实例的pod */
func (c *TidbClient) RebootPod(request *tidb.RebootPodRequest) (*tidb.RebootPodResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.RebootPodResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看实例当前的备份备份策略。 */
func (c *TidbClient) DescribeBackupPolicy(request *tidb.DescribeBackupPolicyRequest) (*tidb.DescribeBackupPolicyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeBackupPolicyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一个TiDB实例 */
func (c *TidbClient) CreateInstance(request *tidb.CreateInstanceRequest) (*tidb.CreateInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.CreateInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取当前用户售罄信息 */
func (c *TidbClient) DescribeOrderableInstanceType(request *tidb.DescribeOrderableInstanceTypeRequest) (*tidb.DescribeOrderableInstanceTypeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeOrderableInstanceTypeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取某个实例下的节点信息 */
func (c *TidbClient) DescribeNodes(request *tidb.DescribeNodesRequest) (*tidb.DescribeNodesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeNodesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除TiDB的备份，仅允许删除用户生成的备份，系统自动备份不允许删除。 */
func (c *TidbClient) DeleteBackup(request *tidb.DeleteBackupRequest) (*tidb.DeleteBackupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DeleteBackupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询 TiDB 实例的详细信息 */
func (c *TidbClient) DescribeInstanceAttributes(request *tidb.DescribeInstanceAttributesRequest) (*tidb.DescribeInstanceAttributesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeInstanceAttributesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询TiDB数据库的升级计划 */
func (c *TidbClient) DescribeUpgradePlan(request *tidb.DescribeUpgradePlanRequest) (*tidb.DescribeUpgradePlanResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeUpgradePlanResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取TiDB数据库可升级到的版本 */
func (c *TidbClient) DescribeUpgradeVersions(request *tidb.DescribeUpgradeVersionsRequest) (*tidb.DescribeUpgradeVersionsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &tidb.DescribeUpgradeVersionsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

