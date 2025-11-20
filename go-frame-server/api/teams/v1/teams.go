package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 创建团队请求参数
type CreateTeamReq struct {
	g.Meta   `path:"/teams" method:"post" tags:"团队管理" summary:"创建团队" security:"BearerAuth"`
	TeamName string `json:"teamName" v:"required|max-length:100#团队名称不能为空|团队名称不能超过100个字符" dc:"团队名称"`
	TeamCode string `json:"teamCode" v:"required|max-length:50#团队代码不能为空|团队代码不能超过50个字符" dc:"团队代码"`
}

// 创建团队响应参数
type CreateTeamRes struct {
	Id        int64       `json:"id" dc:"团队ID"`
	TeamName  string      `json:"teamName" dc:"团队名称"`
	TeamCode  string      `json:"teamCode" dc:"团队代码"`
	OwnerId   int64       `json:"ownerId" dc:"团队所有者ID"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// 获取团队列表请求参数
type ListTeamsReq struct {
	g.Meta `path:"/teams" method:"get" tags:"团队管理" summary:"获取团队列表" security:"BearerAuth"`
	Page   int `json:"page" d:"1" v:"min:1#页码不能小于1" dc:"页码"`
	Size   int `json:"size" d:"10" v:"min:1|max:50#每页大小不能小于1|每页大小不能大于50" dc:"每页大小"`
}

// 获取团队列表响应参数
type ListTeamsRes struct {
	List  []*TeamItem `json:"list" dc:"团队列表"`
	Total int         `json:"total" dc:"总数"`
	Page  int         `json:"page" dc:"页码"`
	Size  int         `json:"size" dc:"每页大小"`
}

// 团队信息项
type TeamItem struct {
	Id        int64       `json:"id" dc:"团队ID"`
	TeamName  string      `json:"teamName" dc:"团队名称"`
	TeamCode  string      `json:"teamCode" dc:"团队代码"`
	OwnerId   int64       `json:"ownerId" dc:"团队所有者ID"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// 获取团队详情请求参数
type GetTeamReq struct {
	g.Meta `path:"/teams/{id}" method:"get" tags:"团队管理" summary:"获取团队详情" security:"BearerAuth"`
	Id     int64 `json:"id" in:"path" v:"required|min:1#团队ID不能为空|团队ID必须大于0" dc:"团队ID"`
}

// 获取团队详情响应参数
type GetTeamRes struct {
	Id                int64       `json:"id" dc:"团队ID"`
	TeamName          string      `json:"teamName" dc:"团队名称"`
	TeamCode          string      `json:"teamCode" dc:"团队代码"`
	OwnerId           int64       `json:"ownerId" dc:"团队所有者ID"`
	TotalWordsQuota   int         `json:"totalWordsQuota" dc:"总字数配额"`
	TotalStorageQuota int         `json:"totalStorageQuota" dc:"总存储配额"`
	TotalCuQuota      int         `json:"totalCuQuota" dc:"总计算单元配额"`
	TotalTrafficQuota int         `json:"totalTrafficQuota" dc:"总流量配额"`
	WordsUsed         int         `json:"wordsUsed" dc:"已使用字数"`
	StorageUsed       int         `json:"storageUsed" dc:"已使用存储"`
	CuUsed            int         `json:"cuUsed" dc:"已使用计算单元"`
	TrafficUsed       float64     `json:"trafficUsed" dc:"已使用流量"`
	MemberCount       int         `json:"memberCount" dc:"成员数量"`
	Status            string      `json:"status" dc:"状态"`
	CreatedAt         *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// 团队成员信息
type TeamMemberInfo struct {
	Id                  int64       `json:"id" dc:"成员ID"`
	UserId              int64       `json:"userId" dc:"用户ID"`
	Role                string      `json:"role" dc:"角色(owner/admin/member)"`
	AllocatedWordsQuota int         `json:"allocatedWordsQuota" dc:"分配的字数配额"`
	PersonalWordsUsed   int         `json:"personalWordsUsed" dc:"个人已使用字数"`
	PersonalStorageUsed int         `json:"personalStorageUsed" dc:"个人已使用存储"`
	PersonalCuUsed      int         `json:"personalCuUsed" dc:"个人已使用计算单元"`
	PersonalTrafficUsed float64     `json:"personalTrafficUsed" dc:"个人已使用流量"`
	Status              string      `json:"status" dc:"状态(active/inactive)"`
	InvitedBy           int64       `json:"invitedBy" dc:"邀请人ID"`
	JoinedAt            *gtime.Time `json:"joinedAt" dc:"加入时间"`
}

// 获取团队成员列表请求参数
type ListTeamMembersReq struct {
	g.Meta `path:"/teams/{teamId}/members" method:"get" tags:"团队管理" summary:"获取团队成员列表" security:"BearerAuth"`
	TeamId int64 `json:"teamId" in:"path" v:"required|min:1#团队ID不能为空|团队ID必须大于0" dc:"团队ID"`
	Page   int   `json:"page" d:"1" v:"min:1#页码不能小于1" dc:"页码"`
	Size   int   `json:"size" d:"10" v:"min:1|max:50#每页大小不能小于1|每页大小不能大于50" dc:"每页大小"`
}

// 获取团队成员列表响应参数
type ListTeamMembersRes struct {
	List  []*TeamMemberInfo `json:"list" dc:"成员列表"`
	Total int               `json:"total" dc:"总数"`
	Page  int               `json:"page" dc:"页码"`
	Size  int               `json:"size" dc:"每页大小"`
}

// 通过邀请码加入团队请求参数
type JoinTeamByCodeReq struct {
	g.Meta     `path:"/teams/join-by-code" method:"post" tags:"团队管理" summary:"通过邀请码加入团队" security:"BearerAuth"`
	InviteCode string `json:"inviteCode" v:"required|max-length:50#邀请码不能为空|邀请码不能超过50个字符" dc:"团队邀请码"`
}

// 通过邀请码加入团队响应参数
type JoinTeamByCodeRes struct {
	Success bool   `json:"success" dc:"操作是否成功"`
	Message string `json:"message" dc:"操作结果描述"`
	TeamId  int64  `json:"teamId" dc:"加入的团队ID"`
}

// 更新团队邀请码请求参数
type UpdateTeamInviteCodeReq struct {
	g.Meta `path:"/teams/{teamId}/invite-code" method:"put" tags:"团队管理" summary:"更新团队邀请码" security:"BearerAuth"`
	TeamId int64 `json:"teamId" in:"path" v:"required|min:1#团队ID不能为空|团队ID必须大于0" dc:"团队ID"`
}

// 更新团队邀请码响应参数
type UpdateTeamInviteCodeReqRes struct {
	Success    bool   `json:"success" dc:"操作是否成功"`
	Message    string `json:"message" dc:"操作结果描述"`
	InviteCode string `json:"inviteCode" dc:"新的邀请码"`
}

// 移除成员请求参数
type RemoveMemberReq struct {
	g.Meta   `path:"/teams/{teamId}/members/{memberId}" method:"delete" tags:"团队管理" summary:"移除团队成员" security:"BearerAuth"`
	TeamId   int64 `json:"teamId" in:"path" v:"required|min:1#团队ID不能为空|团队ID必须大于0" dc:"团队ID"`
	MemberId int64 `json:"memberId" in:"path" v:"required|min:1#成员ID不能为空|成员ID必须大于0" dc:"成员ID"`
}

// 移除成员响应参数
type RemoveMemberRes struct {
	Success bool   `json:"success" dc:"操作是否成功"`
	Message string `json:"message" dc:"操作结果描述"`
}

// 设置团队成员角色请求参数
type SetMemberRoleReq struct {
	g.Meta   `path:"/teams/{teamId}/members/{memberId}/role" method:"put" tags:"团队管理" summary:"设置团队成员角色" security:"BearerAuth"`
	TeamId   int64  `json:"teamId" in:"path" v:"required|min:1#团队ID不能为空|团队ID必须大于0" dc:"团队ID"`
	MemberId int64  `json:"memberId" in:"path" v:"required|min:1#成员ID不能为空|成员ID必须大于0" dc:"成员ID"`
	Role     string `json:"role" v:"required|in:admin,member#角色不能为空|角色必须是admin或member" dc:"新角色(admin/member)"`
}

// 设置团队成员角色响应参数
type SetMemberRoleRes struct {
	Success bool   `json:"success" dc:"操作是否成功"`
	Message string `json:"message" dc:"操作结果描述"`
	Role    string `json:"role" dc:"新设置的角色"`
}

// 设置团队成员配额请求参数
type SetMemberQuotaReq struct {
	g.Meta              `path:"/teams/{teamId}/members/{memberId}/quota" method:"put" tags:"团队管理" summary:"设置团队成员配额" security:"BearerAuth"`
	TeamId              int64 `json:"teamId" in:"path" v:"required|min:1#团队ID不能为空|团队ID必须大于0" dc:"团队ID"`
	MemberId            int64 `json:"memberId" in:"path" v:"required|min:1#成员ID不能为空|成员ID必须大于0" dc:"成员ID"`
	AllocatedWordsQuota int   `json:"allocatedWordsQuota" v:"min:0#字数配额不能小于0" dc:"分配的字数配额(0表示不限制)"`
	//AllocatedStorageQuota int   `json:"allocatedStorageQuota" v:"min:0#存储配额不能小于0" dc:"分配的存储配额(MB，0表示不限制)"`
	//AllocatedCuQuota      int   `json:"allocatedCuQuota" v:"min:0#计算单元配额不能小于0" dc:"分配的计算单元配额(0表示不限制)"`
	//AllocatedTrafficQuota int   `json:"allocatedTrafficQuota" v:"min:0#流量配额不能小于0" dc:"分配的流量配额(GB，0表示不限制)"`
}

// 设置团队成员配额响应参数
type SetMemberQuotaRes struct {
	Success bool   `json:"success" dc:"操作是否成功"`
	Message string `json:"message" dc:"操作结果描述"`
	Quota   struct {
		AllocatedWordsQuota int `json:"allocatedWordsQuota" dc:"分配的字数配额"`
		//AllocatedStorageQuota int `json:"allocatedStorageQuota" dc:"分配的存储配额"`
		//AllocatedCuQuota      int `json:"allocatedCuQuota" dc:"分配的计算单元配额"`
		//AllocatedTrafficQuota int `json:"allocatedTrafficQuota" dc:"分配的流量配额"`
	} `json:"quota" dc:"更新后的配额信息"`
}
