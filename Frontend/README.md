✄┏┓╋╋┏┳┓╋┏┳━━━┓                                                                         
✄┃┗┓┏┛┃┃╋┃┃┏━━┛                                                                        
✄┗┓┃┃┏┫┃╋┃┃┗━━┓                                                                        
✄╋┃┗┛┃┃┃╋┃┃┏━━┛                                                                        
✄╋┗┓┏┛┃┗━┛┃┗━━┓                                                                        
✄╋╋┗┛╋┗━━━┻━━━┛                                                                        

> vue2.0 学习笔记
---
# 知识点 / 踩坑记录
## 初始化
### 组件介绍
#### main.js
main.js 程序入口文件，是初始化vue实例并使用需要的插件,加载各种公共组件.
```js
import Vue from 'vue';
import App from './App.vue';
```
#### App.vue
App.vue是我们的主组件，页面入口文件 ，所有页面都是在App.vue下进行切换的。也是整个项目的关键，app.vue负责构建定义及页面组件归集。

#### router index.js
router index.js 把准备好路由组件注册到路由里：
```js
import Vue from 'vue'
import Router from 'vue-router'
import Recommed from 'components/recommend/recommend'
import Singer from 'components/singer/singer'
import Rank from 'components/rank/rank'
import Search from 'components/search/search'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: Recommed
    },
    {
      path: '/recommend',
      component: Recommed
    },
    {
      path: '/singer',
      component: Singer
    },
    {
      path: '/rank',
      component: Rank
    },
    {
      path: '/search',
      component: Search
    }
  ]
})
```
#### 其他
- src放置组件和入口文件 
- node_modules为依赖的模块 
- config中配置了路径端口值等 
- build中配置了webpack的基本配置、开发环境配置、生产环境配置等

### Element-UI 按需引入
我们可以只引入需要的组件，以达到减小项目体积的目的。
虽然有点麻烦，但是打包速度以及性能会提升很大。
具体使用参考：https://element.eleme.cn/#/zh-CN/component/quickstart

## 登录模块
### async与await用法
> 参考：https://cloud.tencent.com/developer/article/1623173

```text
async和await基本是组合使用的，async用来声明一个异步方法，返回的是一个promise对象，如果要获取到对应的返回值，就需要使用.then方法（不清楚的可以查看promise对象）；

await只能在async方面的里面使用，让后面的执行语句或方法要等待当前await方法的结果后才能再执行。
```

### 表单数据验证
elemetUI提供了表单数据验证
```js
<el-form :model="LoginForm" :rules="LoginRules" ref="LoginForm" label-width="80px" class="login_form">

再在data() 中定义LoginRules
LoginRules: {
  Username: [
    { required: true, min: 4, max: 12, message: '请输入正确的用户名', trigger: 'blur' }
  ],
```
### jwt保存 使用
这个其实很简单，就是用window.sessionStorage.setItem('token', res.token) 进行保存
使用的话需要引入jwt-decode，做解析，返回的值就是jwt中body的值

### 登录验证/导航守卫
路由index.js 中 有beforeEach方法 通过判断jwt
```js
router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  console.log(token)
  if (to.path !== '/admin') {
    return next()
  } else {
    if (!token || jwtDecode(token).role_id !== 1) {
      next('/login')
    } else {
      next()
    }
  }
})
```
### 回车输入
element-ui，配置回车键会不生效，这是因为el-input在输入框的外层添加了一层的，把input隐藏在子级，所以el-input添加上了keyup无响应
可以加上.native
```js
 <el-input v-model="LoginForm.Password" prefix-icon="el-icon-lock" type="password" @keyup.enter.native="submitForm('LoginForm')"></el-input>
```

## 管理页面
### layout 布局
参考element文档

### 请求拦截器
通过axios请求拦截器添加token

```js
axios.interceptors.request.use(config => {
  config.headers.Authorization = window.sessionStorage.getItem('token')
  return config
})
```

