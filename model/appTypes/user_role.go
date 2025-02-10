package appTypes

// RoleID 用户角色
type RoleID int

const (
	Guest RoleID = iota //游客
	User                // 普通用户
	Admin               // 管理员
)
