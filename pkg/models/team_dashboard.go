package models

type TeamDashboard struct {
	Id          int64 `json:"id"`
	OrgId       int64 `json:"orgId"`
	TeamId      int64 `json:"teamId"`
	DashboardId int64 `json:"dashboardId"`
}

// ---------------------
// COMMANDS

type GetUserTeamDashboardsQuery struct {
	OrgId  int64
	UserId int64
	TeamId int64

	Result map[int64]bool // dashboard ids
}

type UpdateTeamDashboardsCommand struct {
	OrgId        int64
	TeamId       int64
	DashboardIds []int64
}
