package Model

import (
	"Backend/Utils/ErrMsg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type varchar(20)；not null" json:"name"`
}

// 检查分类是否存在
func CheckCategory(name string) int64 {
	var count int64
	db.Model(&Category{}).Where("name = ? ", name).Count(&count)
	return count
}

// 新增分类
func AddCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	} else {
		return ErrMsg.SUCCESS
	}
}

// 删除分类
func DeleteCategory(name string) int {
	err := db.Where("name = ?", name).Delete(&Category{}).Error
	if err != nil {
		return ErrMsg.ERROR
	} else {
		return ErrMsg.SUCCESS
	}
}

// 修改分类
func EditCategory(name string, updatedName string) int64 {
	err := db.Model(&Category{}).Where("name = ?", name).Update("name", updatedName).Error
	if err != nil {
		return ErrMsg.ERROR
	} else {
		return ErrMsg.SUCCESS
	}
}

// 查询分类

func ListCategory(pageNum int, pageSize int) []Category {
	var result []Category
	if pageNum == 0 || pageSize == 0 {
		db.Limit(-1).Find(&result)
	} else {
		db.Limit(pageNum).Offset((pageSize - 1) * pageNum).Find(&result)
	}
	return result
}
