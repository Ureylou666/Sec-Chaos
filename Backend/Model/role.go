package Model

import (
	"Backend/Utils/ErrMsg"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName string `gorm:"type:varchar(20);not null" json:"RoleName"`
}

type RoleToMenu struct {
	gorm.Model
	RoleName string `gorm:"type:varchar(20);not null" json:"RoleName"`
	MenuName string `gorm:"type:varchar(100);not null;" json:"MenuName"`
}

// 检查role是否存在
func CheckRole(ReqRole string) int {
	var count int64
	db.Model(&Role{}).Where("role_name", ReqRole).Count(&count)
	if count > 0 {
		return ErrMsg.ERROR_ROLE_EXIST
	}
	return ErrMsg.ERROR_ROLE_NOT_EXIST
}

// 增加Role
func AddRole(data *Role) int {
	err := db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 增加用户权限
func AddRoleToMenu(ReqRole string, ReqMenu string) int {
	var data RoleToMenu
	data.RoleName = ReqRole
	data.MenuName = ReqMenu
	err := db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 删除用户权限
func DeleteRoleToMenu(ReqRole string) {
	db.Where("role_name = ?", ReqRole).Delete(&RoleToMenu{})
}

// 获取用户权限
