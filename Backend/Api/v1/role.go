package v1

import (
	"Backend/Model"
	"Backend/Utils/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReqRoleToMenu struct {
	RoleUID     string
	Rolename    string
	ReqMenuname []string
}

type ResMainMenu struct {
	MainMenuID   string
	MainMenuName string
	SubMenu      []Model.SubMenu
}

type Response struct {
	RoleUID  string
	Menulist []ResMainMenu
}

// 新增角色
func RoleAdd(c *gin.Context) {
	var data Model.Role
	_ = c.ShouldBindJSON(&data)
	// 检查role是否重复
	code := Model.CheckRole(data.RoleName)
	if code == ErrMsg.ERROR_ROLE_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrMsg.GetErrMsg(code),
		})
		return
	}
	code = Model.AddRole(&data)
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

// 角色编辑
func AddRoleToMenu(c *gin.Context) {
	var reqdata ReqRoleToMenu
	_ = c.ShouldBindJSON(&reqdata)
	// 检查role是否存在
	code := Model.CheckRole(reqdata.Rolename)
	if code == ErrMsg.ERROR_ROLE_NOT_EXIST {
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrMsg.GetErrMsg(code),
		})
		return
	}
	// 重置该角色权限
	Model.DeleteRoleToMenu(reqdata.Rolename)
	// 添加新设定权限
	for i := 0; i < len(reqdata.ReqMenuname); i++ {
		code = Model.CheckPrimaryMenu(reqdata.ReqMenuname[i])
		if code == ErrMsg.ERROR_PrimaryMENU_NOT_EXIST {
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Message": ErrMsg.GetErrMsg(code),
			})
			return
		}
		code = Model.AddRoleToMenu(reqdata.Rolename, reqdata.ReqMenuname[i])
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  ErrMsg.SUCCESS,
		"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
	})
	return
}

// 获取权限对应主菜单列表
func GetRoleToMenu(c *gin.Context) {
	var MainMenu []ResMainMenu
	var ReqRole Model.RoleToMenu
	var MainMenuList []Model.MainMenu
	_ = c.ShouldBindJSON(&ReqRole)
	MainMenuList = Model.GetRoleToMainMenu(ReqRole.RoleUID)
	MainMenu = make([]ResMainMenu, len(MainMenuList))
	for i := 0; i < len(MainMenuList); i++ {
		MainMenu[i].MainMenuID = MainMenuList[i].MenuUID
		MainMenu[i].MainMenuName = MainMenuList[i].MenuName
		MainMenu[i].SubMenu = Model.GetSubMenu(MainMenuList[i].MenuName)
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": ErrMsg.SUCCESS,
		"data":   MainMenu,
	})
	return
}

func GetRolelist(c *gin.Context) {
	var rolelist []Model.Role
	rolelist = Model.GetRoleList()
	c.JSON(http.StatusOK, gin.H{
		"Status": ErrMsg.SUCCESS,
		"data":   rolelist,
	})
}
