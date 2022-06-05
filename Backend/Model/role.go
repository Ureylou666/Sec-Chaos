package Model

import (
	"Backend/Utils/ErrMsg"
	"github.com/google/uuid"
)

type Role struct {
	RoleUID  string `gorm:"type:varchar(50);not null" json:"RoleUID"`
	RoleName string `gorm:"type:varchar(20);not null" json:"RoleName"`
}

type RoleToMenu struct {
	RoleUID  string
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
	temp := uuid.New()
	data.RoleUID = temp.String()
	err := db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 增加用户权限
func AddRoleToMenu(ReqRole string, ReqMenu string) int {
	var data RoleToMenu
	var temp Role
	data.RoleName = ReqRole
	data.MenuName = ReqMenu
	db.Model(&Role{}).Where("role_name", ReqRole).First(&temp)
	data.RoleUID = temp.RoleUID
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

// 获取用户主菜单
func GetRoleToMainMenu(ReqRoleUID string) []MainMenu {
	var Menus []RoleToMenu
	db.Where("role_uid = ?", ReqRoleUID).Find(&Menus)
	output := make([]MainMenu, len(Menus))
	for i := 0; i < len(Menus); i++ {
		db.Where("menu_name = ?", Menus[i].MenuName).First(&output[i])
	}
	return output
}

// 获取角色清单
func GetRoleList() []Role {
	var rolelist []Role
	// 不显示超级管理员
	db.Where("role_uid <> 'f2d6e4cf-8865-4f79-b5e4-cdc51d3f4b76'").Find(&rolelist)
	return rolelist
}
