import Vue from 'vue'
import VueRouter from 'vue-router'
import AdminView from '../views/AdminView.vue'
import LoginView from '../views/LoginView'

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

export default router
