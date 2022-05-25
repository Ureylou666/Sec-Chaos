package v1

import (
	"Backend/Model"
	"Backend/Utils/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 增加PrimaryMenu
func PrimaryMenuAdd(c *gin.Context) {
	var data Model.MainMenu
	_ = c.ShouldBindJSON(&data)
	// 检查主菜单是否存在
	code := Model.CheckPrimaryMenu(data.MenuName)
	if code == ErrMsg.ERROR_PrimaryMENU_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrMsg.GetErrMsg(code),
		})
		return
	}
	code = Model.AddPrimaryMenu(&data)
	// 添加menu错误
	if code == ErrMsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_PrimaryMENU_ADD,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_PrimaryMENU_ADD),
		})
		return
	}
	_ = Model.AddRoleToMenu("超级管理员", data.MenuName)
	c.JSON(http.StatusOK, gin.H{
		"Status":  ErrMsg.SUCCESS,
		"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
	})
}

// 增加SubMenu
func SubMenuAdd(c *gin.Context) {
	var data Model.SubMenu
	_ = c.ShouldBindJSON(&data)
	// 检查主菜单是否存在
	code := Model.CheckPrimaryMenu(data.ParentMenuName)
	if code == ErrMsg.ERROR_PrimaryMENU_NOT_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrMsg.GetErrMsg(code),
		})
		return
	}
	// 检查子菜单是否存在
	code = Model.CheckSubMenu(&data)
	if code == ErrMsg.ERROR_SubMENU_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrMsg.GetErrMsg(code),
		})
		return
	}
	code = Model.AddSubMenu(&data)
	// 添加Submenu错误
	if code == ErrMsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_SubMENU_ADD,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_SubMENU_ADD),
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
