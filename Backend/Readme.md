# GIN 后端学习笔记

> GIN框架后端学习笔记
---
## 初始化
### 数据库初始化（迁移）
gorm操作mysql时报错：invalid memory address or nil pointer dereference
```golang
dsn := Utils.DbUser + ":" + Utils.DbPassword + "@tcp(" + Utils.DbHost + ":" + Utils.DbPort + ")/" + Utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
	}
```
应使用 db, err = 而不是db, err :=

考虑从mysql切换到 pgsql，修改数据库表默认变复数

```go
db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
```
### 数据验证

写tag的时候 不能有空格 不然会无法解析
```go
type User struct {
gorm.Model
UID      string `gorm:"type:varchar(50);not null" json:"uid" `
Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,max=12,min=4"`
Password string `gorm:"type:varchar(64);not null" json:"password" validate:"required,max=20,min=8"`
Role     int    `gorm:"type:int; DEFAULT 1;not null" json:"role"`
}
```
## 登录
### 加密

### gorm钩子

可以使用钩子函数在事务发生前进行加密
```go
func(u *User)Before(){
  u.Password = ScryptPw(u.Password)
}
```
### jwt toke

## 权限管理
### gorm定义数组
使用 https://github.com/lib/pq 包
```go
type Role struct {
	gorm.Model
	RoleName	string			`gorm:"type:varchar(20);not null" json:"RoleName"`
	AuthMenuID	pq.Int64Array	`gorm:"type:integer[];" json:"AuthMenuID"`
}
```
