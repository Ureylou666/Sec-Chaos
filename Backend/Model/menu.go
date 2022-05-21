package Model

import (
	"Backend/Utils/ErrMsg"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	MenuName   string        `gorm:"type:varchar(100);not null;" json:"MenuName"`
	ParentMenu string        `gorm:"type:varchar(100);" json:"ParentMenu"`
	Role       pq.Int64Array `gorm:"type:integer[];" json:"Role"`
}

// 增加Menu
func AddMenu(data *Menu) int {
	data.Role = append(data.Role, int64(99))
	err := db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 添加子节点

// 删除Menu

// 修改Menu

// 查询Menu
//func GetMenu()
