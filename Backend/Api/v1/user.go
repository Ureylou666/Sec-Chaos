package v1

import (
	"Backend/Config"
	"Backend/Model"
	"Backend/Utils"
	"Backend/Utils/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type RoleEdit struct {
	UserName string
	RoleName string
}

type Users struct {
	Username    string
	RoleName    string
	Status      bool
	UpdatedTime time.Time
}

// 添加用户
func UserAdd(c *gin.Context) {
	var data Model.User
	_ = c.ShouldBindJSON(&data)
	// 默认无权限
	data.RoleUID = ""
	// 数据验证
	msg, code := Utils.Validate(&data)
	if code == ErrMsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": msg,
		})
		return
	}
	code = Model.CheckUser(data.Username)
	data.Password = Utils.Scrypt(data.Password)
	if code == ErrMsg.SUCCESS {
		Model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": ErrMsg.GetErrMsg(code),
	})
}

// 删除用户
func UserDelete(c *gin.Context) {
	var ReqUser, DbUser Model.User
	_ = c.ShouldBindJSON(&ReqUser)
	DbUser = Model.SearchUser(ReqUser.Username)
	// 判断是否找到对应用户
	if DbUser.UID == "" {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_USERDELETE_NOTFOUND,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_USERDELETE_NOTFOUND),
		})
	} else {
		Model.DeleteUser(DbUser)
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.SUCCESS,
			"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
		})
	}
}

// 编辑用户状态
func UserEditStatus(c *gin.Context) {
	var data Users
	_ = c.ShouldBindJSON(&data)
	// 判断用户是否存在
	code := Model.CheckUser(data.Username)
	if code != ErrMsg.ERROR_USERNAME_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_USERNAME_NOT_EXIST),
		})
		return
	}
	// 编辑用户状态
	code = Model.EditUserStatus(data.Username)
	if code != ErrMsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrMsg.GetErrMsg(code),
		})
		return
	}

}

// 编辑用户权限
func UserRoleEdit(c *gin.Context) {
	var data RoleEdit
	_ = c.ShouldBindJSON(&data)
	// 判断用户是否存在
	code := Model.CheckUser(data.UserName)
	if code != ErrMsg.ERROR_USERNAME_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_USERNAME_NOT_EXIST),
		})
		return
	}
	// 判断权限是否存在
	code = Model.CheckRole(data.RoleName)
	if code != ErrMsg.ERROR_ROLE_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrMsg.GetErrMsg(code),
		})
		return
	}
	code = Model.EditRoleToUser(data.UserName, data.RoleName)
}

// 查询用户列表
func UserSearchList(c *gin.Context) {
	var userlist []Users
	var queryinfo Config.UserQueryInfo
	_ = c.ShouldBindJSON(&queryinfo)
	data, res_total, users_total := Model.GetListUser(queryinfo)
	// 返回到前端
	if data == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  ErrMsg.ERROR_USERLIST_WRONG,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_USERLIST_WRONG),
		})
		return
	}
	if data != nil {
		userlist = make([]Users, res_total)
		for i := 0; i < int(res_total); i++ {
			userlist[i].Username = data[i].Username
			userlist[i].RoleName = data[i].RoleName
			userlist[i].Status = data[i].Status
			userlist[i].UpdatedTime = data[i].UpdatedAt
		}
		c.JSON(http.StatusOK, gin.H{
			"Status":         200,
			"Message":        ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
			"Total":          users_total,
			"Response_Users": res_total,
			"Data":           userlist,
		})
	}
}
