✄┏━━━┓╋╋╋╋╋╋┏━━━┳┓                                                              
✄┃┏━┓┃╋╋╋╋╋╋┃┏━┓┃┃                                                             
✄┃┗━━┳━━┳━━┓┃┃╋┗┫┗━┳━━┳━━┳━━┓                                                           
✄┗━━┓┃┃━┫┏━┛┃┃╋┏┫┏┓┃┏┓┃┏┓┃━━┫                                                             
✄┃┗━┛┃┃━┫┗━┓┃┗━┛┃┃┃┃┏┓┃┗┛┣━━┃                                                             
✄┗━━━┻━━┻━━┛┗━━━┻┛┗┻┛┗┻━━┻━━┛                                                             


> 有很多想法，但还不成熟，就叫他Chaos吧
---
# 参考
- 特别鸣谢：https://github.com/wejectchen/Ginblog

# 项目架构
## 前端
- 框架：vue2
- 组件库：ElementUI

```
FrontEnd
|-- node_modules
|-- N
```

## 后端
- 语言：go
- 代码安全：[腾讯代码安全指南](https://github.com/Tencent/secguide)
- 配置文件解析：
- 数据库：pgsql
- 数据库处理：[gorm](https://gorm.io/zh_CN/)
- 文件存储： 阿里云OSS
- API认证：[JWT](https://github.com/dgrijalva/jwt-go)
- markdown编辑器：markdown editor 
- 日志处理：[logrus](https://github.com/sirupsen/logrus)
- 数据校验：[validator](https://github.com/go-playground/validator)
- CORS跨域：[Cors](https://pkg.go.dev/github.com/gin-contrib/cors) 参考：https://www.ruanyifeng.com/blog/2016/04/cors.html

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
