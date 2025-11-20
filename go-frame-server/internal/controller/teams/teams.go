package teams

import (
	"context"
	"kgplatform-backend/internal/dao"
	"kgplatform-backend/internal/logic/teams"
	"kgplatform-backend/internal/model/do"
	"kgplatform-backend/internal/model/entity"

	v1 "kgplatform-backend/api/teams/v1"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type ControllerV1 struct {
	teams *teams.Teams
}

func NewV1() *ControllerV1 {
	return &ControllerV1{
		teams: teams.New(),
	}
}

//// getUserIdFromToken 从JWT token中获取用户ID
//func (c *ControllerV1) getUserIdFromToken(ctx context.Context) (int64, error) {
//	r := ghttp.RequestFromCtx(ctx)
//	tokenString := r.Header.Get("Authorization")
//
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte(consts.JwtKey), nil
//	})
//	if err != nil || !token.Valid {
//		return 0, gerror.New("无效的认证token")
//	}
//
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok {
//		return 0, gerror.New("无效的token claims")
//	}
//
//	userIdFloat, ok := claims["Id"].(float64)
//	if !ok {
//		return 0, gerror.New("token中缺少用户ID")
//	}
//
//	return int64(userIdFloat), nil
//}

// CreateTeam 创建团队
func (c *ControllerV1) CreateTeam(ctx context.Context, req *v1.CreateTeamReq) (res *v1.CreateTeamRes, err error) {
	// 获取当前用户ID
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	team, err := c.teams.CreateTeam(ctx, int64(userId), req.TeamName, req.TeamCode)
	if err != nil {
		return nil, err
	}

	res = &v1.CreateTeamRes{
		Id:        team.Id,
		TeamName:  team.TeamName,
		TeamCode:  team.TeamCode,
		OwnerId:   team.OwnerId,
		CreatedAt: team.CreatedAt,
	}
	return
}

// ListTeams 获取团队列表
func (c *ControllerV1) ListTeams(ctx context.Context, req *v1.ListTeamsReq) (res *v1.ListTeamsRes, err error) {
	// 获取当前用户ID
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	teamsList, total, err := c.teams.ListTeams(ctx, int64(userId), req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	// 构造响应列表
	var teamItems []*v1.TeamItem
	for _, team := range teamsList {
		item := &v1.TeamItem{
			Id:        team.Id,
			TeamName:  team.TeamName,
			TeamCode:  team.TeamCode,
			OwnerId:   team.OwnerId,
			CreatedAt: team.CreatedAt,
			UpdatedAt: team.UpdatedAt,
		}
		teamItems = append(teamItems, item)
	}

	res = &v1.ListTeamsRes{
		List:  teamItems,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	}
	return
}

// GetTeam 获取团队详情
func (c *ControllerV1) GetTeam(ctx context.Context, req *v1.GetTeamReq) (res *v1.GetTeamRes, err error) {
	// 获取当前用户ID
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	team, err := c.teams.GetTeam(ctx, req.Id, int64(userId))
	if err != nil {
		return nil, err
	}

	res = &v1.GetTeamRes{
		Id:                team.Id,
		TeamName:          team.TeamName,
		TeamCode:          team.TeamCode,
		OwnerId:           team.OwnerId,
		TotalWordsQuota:   team.TotalWordsQuota,
		TotalStorageQuota: team.TotalStorageQuota,
		TotalCuQuota:      team.TotalCuQuota,
		TotalTrafficQuota: team.TotalTrafficQuota,
		WordsUsed:         team.WordsUsed,
		StorageUsed:       team.StorageUsed,
		CuUsed:            team.CuUsed,
		TrafficUsed:       team.TrafficUsed,
		MemberCount:       team.MemberCount,
		Status:            team.Status,
		CreatedAt:         team.CreatedAt,
		UpdatedAt:         team.UpdatedAt,
	}
	return
}

// ListTeamMembers 获取团队成员列表
func (c *ControllerV1) ListTeamMembers(ctx context.Context, req *v1.ListTeamMembersReq) (res *v1.ListTeamMembersRes, err error) {
	// 获取当前用户ID
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	members, total, err := c.teams.ListTeamMembers(ctx, req.TeamId, int64(userId), req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	// 构造响应列表
	var memberItems []*v1.TeamMemberInfo
	for _, member := range members {
		item := &v1.TeamMemberInfo{
			Id:                  member.Id,
			UserId:              member.UserId,
			Role:                member.Role,
			AllocatedWordsQuota: member.AllocatedWordsQuota,
			PersonalWordsUsed:   member.PersonalWordsUsed,
			PersonalStorageUsed: member.PersonalStorageUsed,
			PersonalCuUsed:      member.PersonalCuUsed,
			PersonalTrafficUsed: member.PersonalTrafficUsed,
			Status:              member.Status,
			InvitedBy:           member.InvitedBy,
			JoinedAt:            member.JoinedAt,
		}
		memberItems = append(memberItems, item)
	}

	res = &v1.ListTeamMembersRes{
		List:  memberItems,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	}
	return
}

// JoinTeamByCode 通过邀请码加入团队
func (c *ControllerV1) JoinTeamByCode(ctx context.Context, req *v1.JoinTeamByCodeReq) (res *v1.JoinTeamByCodeRes, err error) {
	// 获取当前用户ID
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	// 查找团队信息以返回团队ID
	var team entity.Teams
	err = dao.Teams.Ctx(ctx).
		Where(do.Teams{InviteCode: req.InviteCode, Status: "active"}).
		Scan(&team)
	if err != nil {
		return nil, gerror.Wrap(err, "查找团队失败")
	}
	if team.Id == 0 {
		return nil, gerror.New("无效的邀请码")
	}

	// 加入团队
	err = c.teams.JoinTeamByCode(ctx, int64(userId), req.InviteCode)
	if err != nil {
		return nil, err
	}

	res = &v1.JoinTeamByCodeRes{
		Success: true,
		Message: "成功加入团队",
		TeamId:  team.Id,
	}
	return
}

// RemoveMember 移除团队成员
func (c *ControllerV1) RemoveMember(ctx context.Context, req *v1.RemoveMemberReq) (res *v1.RemoveMemberRes, err error) {
	// 获取当前用户ID（操作人ID）
	operatorId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if operatorId == 0 {
		return nil, gerror.New("请先登录")
	}

	err = c.teams.RemoveMember(ctx, req.TeamId, int64(operatorId), req.MemberId)
	if err != nil {
		return nil, err
	}

	res = &v1.RemoveMemberRes{
		Success: true,
		Message: "成员已成功移除",
	}
	return
}

//// UpdateTeamInviteCode 更新团队邀请码
//func (c *ControllerV1) UpdateTeamInviteCode(ctx context.Context, req *v1.UpdateTeamInviteCodeReq) (res *v1.UpdateTeamInviteCodeReqRes, err error) {
//	// 获取当前用户ID
//	operatorId, err := c.getUserIdFromToken(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	// 更新邀请码
//	newInviteCode, err := c.teams.UpdateTeamInviteCode(ctx, req.TeamId, operatorId)
//	if err != nil {
//		return nil, err
//	}
//
//	res = &v1.UpdateTeamInviteCodeReqRes{
//		Success:    true,
//		Message:    "邀请码已更新",
//		InviteCode: newInviteCode,
//	}
//	return
//}

// SetMemberRole 设置团队成员角色
func (c *ControllerV1) SetMemberRole(ctx context.Context, req *v1.SetMemberRoleReq) (res *v1.SetMemberRoleRes, err error) {
	// 获取当前用户ID（操作人ID）
	operatorId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if operatorId == 0 {
		return nil, gerror.New("请先登录")
	}

	// 调用逻辑层方法设置成员角色
	err = c.teams.SetMemberRole(ctx, req.TeamId, int64(operatorId), req.MemberId, req.Role)
	if err != nil {
		return nil, err
	}

	res = &v1.SetMemberRoleRes{
		Success: true,
		Message: "成员角色已成功设置",
		Role:    req.Role,
	}
	return
}

// SetMemberQuota 设置团队成员配额
func (c *ControllerV1) SetMemberQuota(ctx context.Context, req *v1.SetMemberQuotaReq) (res *v1.SetMemberQuotaRes, err error) {
	// 获取当前用户ID（操作人ID）
	operatorId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if operatorId == 0 {
		return nil, gerror.New("请先登录")
	}

	// 调用逻辑层方法设置成员配额
	err = c.teams.SetMemberQuota(ctx, req.TeamId, int64(operatorId), req.MemberId, req.AllocatedWordsQuota)
	if err != nil {
		return nil, err
	}

	res = &v1.SetMemberQuotaRes{
		Success: true,
		Message: "成员配额已成功设置",
		Quota: struct {
			AllocatedWordsQuota int `json:"allocatedWordsQuota" dc:"分配的字数配额"`
			//AllocatedStorageQuota int `json:"allocatedStorageQuota" dc:"分配的存储配额"`
			//AllocatedCuQuota      int `json:"allocatedCuQuota" dc:"分配的计算单元配额"`
			//AllocatedTrafficQuota int `json:"allocatedTrafficQuota" dc:"分配的流量配额"`
		}{
			AllocatedWordsQuota: req.AllocatedWordsQuota,
			//AllocatedStorageQuota: req.AllocatedStorageQuota,
			//AllocatedCuQuota:      req.AllocatedCuQuota,
			//AllocatedTrafficQuota: req.AllocatedTrafficQuota,
		},
	}
	return
}
