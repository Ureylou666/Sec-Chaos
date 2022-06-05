package Routers

import (
	"Backend/Api/v1"
	"Backend/Middleware"
	"Backend/Utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(Utils.AppMode)
	r := gin.Default()
	r.Use(Middleware.LogPrint(), Middleware.Cors())
	auth := r.Group("api/v1")
	auth.Use(Middleware.JwtToken())
	{
		// 用户模块路由接口
		//查询用户列表
		auth.POST("users", v1.UserSearchList)
		//编辑用户权限
		auth.POST("users/editRole", v1.UserRoleEdit)
		//编辑用户权限
		auth.POST("users/editStatus", v1.UserEditStatus)
		//编辑用户密码
		auth.POST("users/editPassword", v1.UserPasswordEdit)
		//删除用户
		auth.POST("users/delete", v1.UserDelete)

		// 分类模块路由接口
		auth.POST("category/add", v1.CategoryAdd)
		auth.POST("category/delete", v1.CategoryDelete)
		auth.GET("category", v1.CategoryList)
		auth.POST("category/edit", v1.CategoryEdit)

		//	Menu菜单栏接口
		//auth.POST("primarymenu/add", v1.PrimaryMenuAdd)
		//auth.POST("submenu/add", v1.SubM)
	}
	public := r.Group("api/v1")
	{
		// 登录模块路由接口
		//登录验证
		public.POST("login", v1.LoginCheck)
		//！！！temp 测试用
		// 添加用户
		public.POST("users/add", v1.UserAdd)
		// 菜单管理
		public.POST("menu/add", v1.PrimaryMenuAdd)
		public.POST("submenu/add", v1.SubMenuAdd)
		public.POST("menu", v1.GetRoleToMenu)

		// 角色权限管理
		public.POST("role/add", v1.RoleAdd)
		public.POST("role/edit", v1.AddRoleToMenu)
		public.GET("roles", v1.GetRolelist)
	}

	r.Run(Utils.HttpPort)
}
