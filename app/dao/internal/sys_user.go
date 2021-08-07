// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SysUserDao is the manager for logic model data accessing and custom defined data operations functions management.
type SysUserDao struct {
	gmvc.M                // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      sysUserColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB         // DB is the raw underlying database management object.
	Table  string         // Table is the underlying table name of the DAO.
}

// SysUserColumns defines and stores column names for table sys_user.
type sysUserColumns struct {
	Id       string // 用户ID
	Passport string // 用户账号
	Password string // 用户密码
	Nickname string // 用户昵称
	CreateAt string // 创建时间
	UpdateAt string // 更新时间
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao() *SysUserDao {
	columns := sysUserColumns{
		Id:       "id",
		Passport: "passport",
		Password: "password",
		Nickname: "nickname",
		CreateAt: "create_at",
		UpdateAt: "update_at",
	}
	return &SysUserDao{
		C:     columns,
		M:     g.DB("default").Model("sys_user").Safe(),
		DB:    g.DB("default"),
		Table: "sys_user",
	}
}
