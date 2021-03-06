package Model

import (
	"Backend/Utils/ErrMsg"
	"github.com/google/uuid"
)

type SubMenu struct {
	SubMenuUID     string `gorm:"type:varchar(50);not null" json:"SubMenuUID"`
	SubMenuName    string `gorm:"type:varchar(100);not null;" json:"SubMenuName"`
	SubMenuPath    string `gorm:"type:varchar(100);not null;" json:"SubMenuPath"`
	ParentMenuName string `gorm:"type:varchar(100);not null;" json:"ParentMenuName"`
}

type MainMenu struct {
	MenuUID  string `gorm:"type:varchar(50);not null" json:"MenuUID"`
	MenuName string `gorm:"type:varchar(100);not null;" json:"MenuName"`
	MenuPath string `gorm:"type:varchar(100);not null;" json:"MenuPath"`
}

// 检查Menu是否存在
func CheckPrimaryMenu(ReqMenuName string) int {
	var count int64
	db.Model(&MainMenu{}).Where("menu_name = ? ", ReqMenuName).Count(&count)
	if count > 0 {
		return ErrMsg.ERROR_PrimaryMENU_EXIST
	}
	return ErrMsg.ERROR_PrimaryMENU_NOT_EXIST
}

// 检查SubMenu是否存在
func CheckSubMenu(data *SubMenu) int {
	var count int64
	db.Model(&SubMenu{}).Where("parent_menu_name = ? AND sub_menu_name = ?", data.ParentMenuName, data.SubMenuName).Count(&count)
	if count > 0 {
		return ErrMsg.ERROR_SubMENU_EXIST
	}
	return ErrMsg.SUCCESS
}

// 增加Menu
func AddPrimaryMenu(data *MainMenu) int {
	temp := uuid.New()
	data.MenuUID = temp.String()
	err := db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 添加SubMenu
func AddSubMenu(data *SubMenu) int {
	temp := uuid.New()
	data.SubMenuUID = temp.String()
	err := db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 删除Menu

// 更新Menu

// 通过主菜单获取子菜单
func GetSubMenu(MainMenu string) []SubMenu {
	var SubmenuList []SubMenu
	db.Where("parent_menu_name = ?", MainMenu).Find(&SubmenuList)
	return SubmenuList
}
