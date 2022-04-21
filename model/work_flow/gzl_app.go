package work_flow

import (
	"project/global"
	"project/model/system"
	"project/utils"
)

//GzlApp 应用表
type GzlApp struct {
	global.GSD_MODEL
	Name       string                `json:"name" gorm:"not null;comment:名称"`                               //名称
	Icon       string                `json:"icon" gorm:"not null;comment:图标"`                               //图标
	Form       utils.JSON            `json:"form" gorm:"type:json;comment:表单"`                              //表单
	Flow       utils.JSON            `json:"flow" gorm:"type:json;comment:流程"`                              //流程
	IsEveryone uint8                 `json:"isEveryone" gorm:"not null;default:1;comment:是否所有人(所有人1默认，否2)"` //是否所有人(所有人1默认，否2)
	IsEnable   uint8                 `json:"isEnable" gorm:"not null;default:1;comment:是否启用(不启用1默认，启用2)"`
	Depts      []system.SysDept      `json:"depts" gorm:"many2many:gzl_app_dept"`           //部门
	Authoritys []system.SysAuthority `json:"authoritys" gorm:"many2many:gzl_app_authority"` //角色
	Users      []system.SysUser      `json:"users" gorm:"many2many:gzl_app_user"`           //用户
}