import Vue from 'vue'
import VueRouter from 'vue-router'
import AdminView from '../views/AdminView.vue'
import LoginView from '../views/LoginView.vue'
// import jwtDecode from 'jwt-decode'
// 页面路由组件
import UserAdd from '../components/admin/User/UserAdd.vue'

Vue.use(VueRouter)
const routes = [
  { path: '/', redirect: '/login' },
  { path: '/login', component: LoginView },
  {
    path: '/admin',
    component: AdminView,
    children: [
      { path: '/admin/users', name: 'UserManager', component: () => import('../components/admin/User/UserManager') },
      { path: '/admin/useradd', name: 'UserAdd', component: UserAdd }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

// router.beforeEach((to, from, next) => {
//   const token = window.sessionStorage.getItem('token')
//  console.log(token)
//  if (to.path !== '/admin') {
//    return next()
//  } else {
//    if (!token || jwtDecode(token).role_id !== 99) {
//      next('/login')
//    } else {
//      next()
//    }
//  }
// })

export default router
