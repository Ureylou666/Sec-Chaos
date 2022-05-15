# GIN 后端学习笔记

> GIN框架后端学习笔记
---
## 参考文档
- gin
- gorm：https://gorm.io/zh_CN/
- jwt: https://github.com/dgrijalva/jwt-go
- golang教程：https://astaxie.gitbooks.io/build-web-application-with-golang/content/zh/
- 项目实战视频：https://www.bilibili.com/video/BV1AD4y1D7BX
- golang包：https://pkg.go.dev/

## 项目架构

```
Backend
|-- Api
    |-- V1
        |-- user.go     /* 用户注册模块 */
        |-- login.go    /* 用户登录模块 */
|-- Config              /* 存放配置文件 */
    |-- config.ini
|-- Middleware          /* 配置中间价 */
    |--  jwt.go         /* jwt生成*/
|-- Model               /* 数据库相关操作 */
    |-- article.go
    |-- category.go
    |-- db.go
    |-- user.go         /* 定义用户相关增删改查 */
|-- Routers             /* 配置项目api路由 */
    |-- router.go
|-- Utils
    |-- ErrMsg          /* 配置错误控制 */
        |-- err.go
    |-- encrypt.go      /* 处理加解密问题 */
    |-- setting.go      /* 读取config.ini配置文件 */
|-- main.go             /* 项目主程序 */
|-  go.mod
```

## 踩坑笔记

- gorm操作mysql时报错：invalid memory address or nil pointer dereference
```golang
dsn := Utils.DbUser + ":" + Utils.DbPassword + "@tcp(" + Utils.DbHost + ":" + Utils.DbPort + ")/" + Utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
	}
```
应使用 db, err = 而不是db, err :=

- 加密

- gorm钩子

可以使用钩子函数在事务发生前进行加密
```go
func(u *User)Before(){
  u.Password = ScryptPw(u.Password)
}
```
- jwt toke

