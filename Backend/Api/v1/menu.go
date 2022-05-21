package v1

import (
	"Backend/Model"
	"Backend/Utils/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 增加Menu
func MenuAdd(c *gin.Context) {
	var data Model.Menu
	_ = c.ShouldBindJSON(&data)
	code := Model.AddMenu(&data)
	// 添加menu错误
	if code == ErrMsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_MENU_ADD,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_MENU_ADD),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  ErrMsg.SUCCESS,
		"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
	})
}

// 删除Menu

// 修改Menu

// 获取Menu
// func MenuGet((c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"Status":  ErrMsg.SUCCESS,
// 		"AuthMenu": "用户管理",
// 		"Children": [
// 		{"id": 104, "AuthSubMenu": "用户列表"},]
// 	})
// 	return
// }
