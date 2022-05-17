✄┏┓╋╋┏┳┓╋┏┳━━━┓                                                                         
✄┃┗┓┏┛┃┃╋┃┃┏━━┛                                                                        
✄┗┓┃┃┏┫┃╋┃┃┗━━┓                                                                        
✄╋┃┗┛┃┃┃╋┃┃┏━━┛                                                                        
✄╋┗┓┏┛┃┗━┛┃┗━━┓                                                                        
✄╋╋┗┛╋┗━━━┻━━━┛                                                                        

> vue2.0 学习笔记
---
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
