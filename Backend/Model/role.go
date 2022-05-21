package Model

import (
	"Backend/Utils/ErrMsg"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName   string        `gorm:"type:varchar(20);not null" json:"RoleName"`
	AuthMenuID pq.Int64Array `gorm:"type:integer[];" json:"AuthMenuID"`
}

// 增加Role
func AddRole(data *Role) int {
	err := db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}
