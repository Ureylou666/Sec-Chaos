import Vue from 'vue'
import VueRouter from 'vue-router'
import AdminView from '../views/AdminView.vue'
import LoginView from '../views/LoginView'
import jwtDecode from 'jwt-decode'
Vue.use(VueRouter)

const routes = [
  { path: '/admin', component: AdminView },
  { path: '/', redirect: '/login' },
  { path: '/login', component: LoginView }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

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

export default router
