package model

import (
	"github.com/AbnerEarl/goutils/dbs"
	"time"
)

type AccountStatusStruct struct {
	Inuse  uint `json:"inuse"`
	Forbid uint `json:"forbid"`
}

func AccountStatus() AccountStatusStruct {
	return AccountStatusStruct{
		Inuse:  1,
		Forbid: 2,
	}
}

type AccountRoleStruct struct {
	SuperAdmin  uint `json:"super_admin"`
	NormalAdmin uint `json:"normal_admin"`
}

func AccountRole() AccountRoleStruct {
	return AccountRoleStruct{
		SuperAdmin:  1,
		NormalAdmin: 2,
	}
}

type AccountModel struct {
	dbs.BaseModel
	Username    string    `json:"username" gorm:"column:username;not null;comment:'用户账号'"`
	Nickname    string    `json:"nickname" gorm:"column:nickname;not null;comment:'用户名称'"`
	Password    string    `json:"password" gorm:"column:password;not null;comment:'用户密码'"`
	Role        uint      `json:"role" gorm:"column:role;default:1;not null;comment:'用户角色'"`
	Status      uint      `json:"status" gorm:"column:status;default:1;not null;comment:'用户状态'"`
	Permissions uint      `json:"permissions" gorm:"column:permissions;default:0;comment:'用户权限'"`
	LastLogin   time.Time `json:"last_login" gorm:"column:last_login;comment:'最后登录时间'"`
}

func (m *AccountModel) TableName() string {
	return "account_info"
}
