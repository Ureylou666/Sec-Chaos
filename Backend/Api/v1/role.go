package v1

import (
	"Backend/Model"
	"Backend/Utils/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 增加Menu
func RoleAdd(c *gin.Context) {
	var data Model.Role
	_ = c.ShouldBindJSON(&data)
	code := Model.AddRole(&data)
	// 添加role错误
	if code == ErrMsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_ROLE_ADD,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_ROLE_ADD),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  ErrMsg.SUCCESS,
		"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
	})
}
