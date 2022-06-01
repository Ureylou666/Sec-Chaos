package Model

import (
	"Backend/Config"
	"Backend/Utils/ErrMsg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QueryInfo struct {
	Query    string
	Pagenum  int
	Pagesize int
}

type User struct {
	gorm.Model
	UID      string `gorm:"type:varchar(50);not null" json:"uid" `
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,printascii,max=12,min=4" label:"用户名"`
	Password string `gorm:"type:varchar(64);not null" json:"password" validate:"required,printascii,max=20,min=8" label:"密码"`
	RoleUID  string
	RoleName string
	Status   bool `gorm:"type:bool;not null" json:"status" `
}

// 检查用户是否存在
func CheckUser(name string) int {
	var count int64
	db.Model(&User{}).Where("Username = ? ", name).Count(&count)
	if count > 0 {
		return ErrMsg.ERROR_USERNAME_EXIST //1002
	}
	return ErrMsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	temp := uuid.New()
	data.UID = temp.String()
	data.Status = true
	err := db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 用户授权
func EditRoleToUser(username string, rolename string) int {
	var ReqUser User
	var ReqRole Role
	db.Where("username = ?", username).First(&ReqUser)
	db.Where("role_name = ?", rolename).First(&ReqRole)
	err := db.Model(&User{}).Where("username = ?", ReqUser.Username).Update("role_uid", ReqRole.RoleUID).Error
	err = db.Model(&User{}).Where("username = ?", ReqUser.Username).Update("role_name", ReqRole.RoleName).Error
	if err != nil {
		return ErrMsg.ERROR
	} else {
		return ErrMsg.SUCCESS
	}
}

// 删除用户
func DeleteUser(data User) {
	db.Delete(&data)
}

// 编辑用户状态
func EditUserStatus(username string) int {
	var data User
	db.Where("username = ?", username).First(&data)
	data.Status = !data.Status
	err = db.Save(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	} else {
		return ErrMsg.SUCCESS
	}
}

// 查询用户列表
func GetListUser(data Config.UserQueryInfo) ([]User, int64, int64) {
	var result []User
	var res_total, users_totle int64
	// 当没有定义查询用户时
	if data.Query == "" {
		db.Limit(-1).Find(&result).Count(&users_totle)
		if data.Pagenum == 0 || data.Pagesize == 0 {
			db.Limit(-1).Find(&result)
			res_total = int64(len(result))
		} else {
			db.Limit(data.Pagesize).Offset((data.Pagenum - 1) * data.Pagesize).Find(&result)
			res_total = int64(len(result))
		}
		return result, res_total, users_totle
	}
	// 当有目标查询用户时
	db.Where("username like ?", "%"+data.Query+"%").Find(&result).Count(&users_totle)
	if data.Pagenum == 0 || data.Pagesize == 0 {
		db.Where("username like ?", "%"+data.Query+"%").Find(&result)
		res_total = int64(len(result))
	} else {
		db.Where("username like ?", "%"+data.Query+"%").Limit(data.Pagesize).Offset((data.Pagenum - 1) * data.Pagesize).Find(&result)
		res_total = int64(len(result))
	}
	return result, res_total, users_totle
}

// 查询单个用户
func SearchUser(name string) User {
	var result User
	db.Where("username = ?", name).First(&result)
	return result
}
