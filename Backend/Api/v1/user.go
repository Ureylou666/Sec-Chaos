package v1

import (
	"Backend/Model"
	"Backend/Utils"
	"Backend/Utils/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加用户
func UserAdd(c *gin.Context) {
	var data Model.User
	_ = c.ShouldBindJSON(&data)
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
		Model.DeleteUser(ReqUser.UID)
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.SUCCESS,
			"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
		})
	}
}

// 编辑用户密码
func UserEdit(c *gin.Context) {

}

// 查询用户列表
func UserSearchList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data, total := Model.GetListUser(pageSize, pageNum)
	if data == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  ErrMsg.ERROR_USERLIST_WRONG,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_USERLIST_WRONG),
		})
	}
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  200,
			"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
			"Total":   total,
			"Data":    data,
		})
	}
}
