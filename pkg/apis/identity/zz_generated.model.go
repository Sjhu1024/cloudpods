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

// Code generated by model-api-gen. DO NOT EDIT.

package identity

import (
	time "time"

	jsonutils "yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/apis"
	tagutils "yunion.io/x/onecloud/pkg/util/tagutils"
)

// SAssignment is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SAssignment.
type SAssignment struct {
	apis.SResourceBase
	Type      string `json:"type"`
	ActorId   string `json:"actor_id"`
	TargetId  string `json:"target_id"`
	RoleId    string `json:"role_id"`
	Inherited *bool  `json:"inherited,omitempty"`
}

// SConfigOption is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SConfigOption.
type SConfigOption struct {
	apis.SResourceBase
	apis.SRecordChecksumResourceBase
	ResType string               `json:"res_type"`
	ResId   string               `json:"domain_id"`
	Group   string               `json:"group"`
	Option  string               `json:"option"`
	Value   jsonutils.JSONObject `json:"value"`
}

// SCredential is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SCredential.
type SCredential struct {
	apis.SStandaloneResourceBase
	UserId        string              `json:"user_id"`
	ProjectId     string              `json:"project_id"`
	Type          string              `json:"type"`
	KeyHash       string              `json:"key_hash"`
	Extra         *jsonutils.JSONDict `json:"extra"`
	EncryptedBlob string              `json:"encrypted_blob"`
	Enabled       *bool               `json:"enabled,omitempty"`
}

// SDomain is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SDomain.
type SDomain struct {
	apis.SStandaloneResourceBase
	Extra    *jsonutils.JSONDict `json:"extra"`
	Enabled  *bool               `json:"enabled,omitempty"`
	IsDomain *bool               `json:"is_domain,omitempty"`
	DomainId string              `json:"domain_id"`
	ParentId string              `json:"parent_id"`
}

// SEnabledIdentityBaseResource is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SEnabledIdentityBaseResource.
type SEnabledIdentityBaseResource struct {
	SIdentityBaseResource
	Enabled *bool `json:"enabled,omitempty"`
}

// SEndpoint is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SEndpoint.
type SEndpoint struct {
	apis.SStandaloneResourceBase
	LegacyEndpointId     string              `json:"legacy_endpoint_id"`
	Interface            string              `json:"interface"`
	ServiceId            string              `json:"service_id"`
	Url                  string              `json:"url"`
	Extra                *jsonutils.JSONDict `json:"extra"`
	Enabled              *bool               `json:"enabled,omitempty"`
	RegionId             string              `json:"region_id"`
	ServiceCertificateId string              `json:"service_certificate_id"`
}

// SFederatedUser is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SFederatedUser.
type SFederatedUser struct {
	apis.SResourceBase
	Id          int    `json:"id"`
	UserId      string `json:"user_id"`
	IdpId       string `json:"idp_id"`
	ProtocolId  string `json:"protocol_id"`
	UniqueId    string `json:"unique_id"`
	DisplayName string `json:"display_name"`
}

// SFederationProtocol is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SFederationProtocol.
type SFederationProtocol struct {
	Id        string `json:"id"`
	IdpId     string `json:"idp_id"`
	MappingId string `json:"mapping_id"`
}

// SFernetKey is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SFernetKey.
type SFernetKey struct {
	Type  string `json:"type"`
	Index int    `json:"index"`
	Key   string `json:"key"`
}

// SGroup is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SGroup.
type SGroup struct {
	SIdentityBaseResource
	Displayname string `json:"displayname"`
}

// SIdentityBaseResource is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SIdentityBaseResource.
type SIdentityBaseResource struct {
	apis.SStandaloneResourceBase
	apis.SDomainizedResourceBase
	Extra *jsonutils.JSONDict `json:"extra"`
}

// SIdentityProvider is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SIdentityProvider.
type SIdentityProvider struct {
	apis.SEnabledStatusStandaloneResourceBase
	apis.SDomainizedResourceBase
	Driver         string `json:"driver"`
	Template       string `json:"template"`
	TargetDomainId string `json:"target_domain_id"`
	// 是否自动创建项目
	AutoCreateProject *bool `json:"auto_create_project,omitempty"`
	// 是否自动创建用户
	AutoCreateUser *bool     `json:"auto_create_user,omitempty"`
	ErrorCount     int       `json:"error_count"`
	SyncStatus     string    `json:"sync_status"`
	LastSync       time.Time `json:"last_sync"`
	// = Column(DateTime, nullable=True)
	LastSyncEndAt       time.Time `json:"last_sync_end_at"`
	SyncIntervalSeconds int       `json:"sync_interval_seconds"`
	// 认证源图标
	IconUri string `json:"icon_uri"`
	// 是否是SSO登录方式
	IsSso *bool `json:"is_sso,omitempty"`
	// 是否是缺省SSO登录方式
	IsDefault *bool `json:"is_default,omitempty"`
}

// SIdmapping is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SIdmapping.
type SIdmapping struct {
	apis.SResourceBase
	PublicId    string `json:"public_id"`
	IdpId       string `json:"domain_id"`
	IdpEntityId string `json:"local_id"`
	EntityType  string `json:"entity_type"`
}

// SIdpRemoteIds is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SIdpRemoteIds.
type SIdpRemoteIds struct {
	IdpId    string `json:"idp_id"`
	RemoteId string `json:"remote_id"`
}

// SImpliedRole is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SImpliedRole.
type SImpliedRole struct {
	PriorRoleId   string `json:"prior_role_id"`
	ImpliedRoleId string `json:"implied_role_id"`
}

// SLocalUser is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SLocalUser.
type SLocalUser struct {
	apis.SResourceBase
	Id                int    `json:"id"`
	UserId            string `json:"user_id"`
	DomainId          string `json:"domain_id"`
	Name              string `json:"name"`
	FailedAuthCount   int    `json:"failed_auth_count"`
	NeedResetPassword *bool  `json:"need_reset_password,omitempty"`
}

// SNonlocalUser is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SNonlocalUser.
type SNonlocalUser struct {
	DomainId string `json:"domain_id"`
	Name     string `json:"name"`
	UserId   string `json:"user_id"`
}

// SPassword is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SPassword.
type SPassword struct {
	apis.SResourceBase
	Id           int    `json:"id"`
	LocalUserId  int    `json:"local_user_id"`
	Password     string `json:"password"`
	SelfService  bool   `json:"self_service"`
	PasswordHash string `json:"password_hash"`
	CreatedAtInt int64  `json:"created_at_int"`
	ExpiresAtInt int64  `json:"expires_at_int"`
}

// SPolicy is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SPolicy.
type SPolicy struct {
	SEnabledIdentityBaseResource
	apis.SSharableBaseResource
	// swagger:ignore
	// Deprecated
	Type string `json:"type"`
	// 权限定义
	Blob jsonutils.JSONObject `json:"blob"`
	// 权限范围
	Scope string `json:"scope"`
	// 是否为系统权限
	IsSystem *bool `json:"is_system,omitempty"`
	// 匹配的项目标签
	ProjectTags []tagutils.STag `json:"project_tags"`
	// 匹配的域标签
	DomainTags []tagutils.STag `json:"domain_tags"`
	// 匹配的资源标签
	ObjectTags []tagutils.STag `json:"object_tags"`
}

// SProject is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SProject.
type SProject struct {
	SIdentityBaseResource
	ParentId string `json:"parent_id"`
	IsDomain *bool  `json:"is_domain,omitempty"`
}

// SProjectExtended is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SProjectExtended.
type SProjectExtended struct {
	SProject
	DomainName string `json:"domain_name"`
}

// SRegion is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SRegion.
type SRegion struct {
	apis.SStandaloneResourceBase
	ParentRegionId string              `json:"parent_region_id"`
	Extra          *jsonutils.JSONDict `json:"extra"`
}

// SRole is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SRole.
type SRole struct {
	SIdentityBaseResource
	apis.SSharableBaseResource
}

// SRolePolicy is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SRolePolicy.
type SRolePolicy struct {
	apis.SResourceBase
	// 角色ID, 主键
	RoleId string `json:"role_id"`
	// 项目ID，主键
	ProjectId string `json:"project_id"`
	// 权限ID, 主键
	PolicyId string `json:"policy_id"`
	// 是否需要认证
	Auth *bool `json:"auth,omitempty"`
	// 匹配的IP白名单
	Ips string `json:"ips"`
	// 匹配开始时间
	ValidSince time.Time `json:"valid_since"`
	// 匹配结束时间
	ValidUntil time.Time `json:"valid_until"`
}

// SScopeResource is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SScopeResource.
type SScopeResource struct {
	DomainId  string `json:"domain_id"`
	ProjectId string `json:"project_id"`
	OwnerId   string `json:"owner_id"`
	RegionId  string `json:"region_id"`
	ServiceId string `json:"service_id"`
	Resource  string `json:"resource"`
	Count     int    `json:"count"`
}

// SService is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SService.
type SService struct {
	apis.SStandaloneResourceBase
	Type          string              `json:"type"`
	Enabled       *bool               `json:"enabled,omitempty"`
	Extra         *jsonutils.JSONDict `json:"extra"`
	ConfigVersion int                 `json:"config_version"`
}

// SServiceCertificate is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SServiceCertificate.
type SServiceCertificate struct {
	apis.SStandaloneResourceBase
	apis.SCertificateResourceBase
	CaCertificate string `json:"ca_certificate"`
	CaPrivateKey  string `json:"ca_private_key"`
}

// SUser is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SUser.
type SUser struct {
	apis.SRecordChecksumResourceBase
	SEnabledIdentityBaseResource
	Email           string    `json:"email"`
	Mobile          string    `json:"mobile"`
	Displayname     string    `json:"displayname"`
	LastActiveAt    time.Time `json:"last_active_at"`
	LastLoginIp     string    `json:"last_login_ip"`
	LastLoginSource string    `json:"last_login_source"`
	IsSystemAccount *bool     `json:"is_system_account,omitempty"`
	// deprecated
	DefaultProjectId string `json:"default_project_id"`
	AllowWebConsole  *bool  `json:"allow_web_console,omitempty"`
	EnableMfa        *bool  `json:"enable_mfa,omitempty"`
	Lang             string `json:"lang"`
}

// SUserOption is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SUserOption.
type SUserOption struct {
	UserId      string `json:"user_id"`
	OptionId    string `json:"option_id"`
	OptionValue string `json:"option_value"`
}

// SUsergroupMembership is an autogenerated struct via yunion.io/x/onecloud/pkg/keystone/models.SUsergroupMembership.
type SUsergroupMembership struct {
	apis.SResourceBase
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`
}
