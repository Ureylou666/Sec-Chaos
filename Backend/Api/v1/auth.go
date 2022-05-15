package v1

import (
	"Backend/Middleware"
	"Backend/Model"
	"Backend/Utils"
	"Backend/Utils/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginCheck(c *gin.Context) {
	var data Model.User
	_ = c.ShouldBindJSON(&data)
	code := Model.CheckUser(data.Username)
	// 判断用户是否存在
	if code == ErrMsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"code":    ErrMsg.ERROR_PASSWORD_WRONG,
			"message": ErrMsg.GetErrMsg(ErrMsg.ERROR_PASSWORD_WRONG),
		})
	} else {
		// 比对密码是否正确
		data.Password = Utils.Scrypt(data.Password)
		if data.Password != Model.SearchUser(data.Username, "password") {
			c.JSON(http.StatusOK, gin.H{
				"code":    ErrMsg.ERROR_PASSWORD_WRONG,
				"message": ErrMsg.GetErrMsg(ErrMsg.ERROR_PASSWORD_WRONG),
			})
		} else {
			token, _ := Middleware.SetToken(data.Username, data.Password)
			c.JSON(http.StatusOK, gin.H{
				"code":    ErrMsg.SUCCESS,
				"message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
				"token":   token,
			})
		}
	}
}
