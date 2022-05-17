import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView'

Vue.use(VueRouter)

const routes = [
  { path: '/home', component: HomeView },
  { path: '/', redirect: '/login' },
  { path: '/login', component: LoginView }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
