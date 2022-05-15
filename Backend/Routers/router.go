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
	auth := r.Group("api/v1")
	auth.Use(Middleware.JwtToken())
	{
		// 用户模块路由接口
		//查询用户列表
		auth.GET("users", v1.UserSearchList)
		//编辑用户
		auth.POST("users/edit", v1.UserEdit)
		//删除用户
		auth.POST("users/delete", v1.UserDelete)

		// 分类模块路由接口
		auth.POST("category/add", v1.CategoryAdd)
		auth.POST("category/delete", v1.CategoryDelete)
		auth.GET("category", v1.CategoryList)
		auth.POST("category/edit", v1.CategoryEdit)
	}
	public := r.Group("api/v1")
	{
		// 添加用户
		public.POST("user/add", v1.UserAdd)
		// 登录模块路由接口
		//登录验证
		public.POST("login", v1.LoginCheck)
	}

	r.Run(Utils.HttpPort)
}
