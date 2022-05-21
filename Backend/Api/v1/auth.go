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
	var ReqUser, DbUser Model.User
	_ = c.ShouldBindJSON(&ReqUser)
	code := Model.CheckUser(ReqUser.Username)
	// 判断用户是否存在
	if code == ErrMsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"code":    ErrMsg.ERROR_PASSWORD_WRONG,
			"message": ErrMsg.GetErrMsg(ErrMsg.ERROR_PASSWORD_WRONG),
		})
	} else {
		// 比对密码是否正确
		DbUser = Model.SearchUser(ReqUser.Username)
		ReqUser.Password = Utils.Scrypt(ReqUser.Password)
		if ReqUser.Password != DbUser.Password {
			c.JSON(http.StatusOK, gin.H{
				"code":    ErrMsg.ERROR_PASSWORD_WRONG,
				"message": ErrMsg.GetErrMsg(ErrMsg.ERROR_PASSWORD_WRONG),
			})
		} else {
			token, _ := Middleware.SetToken(ReqUser.Username, DbUser.RoleId)
			c.JSON(http.StatusOK, gin.H{
				"code":    ErrMsg.SUCCESS,
				"message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
				"token":   token,
			})
		}
	}
}
