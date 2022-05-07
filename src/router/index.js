import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/home',
    component: HomeView
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
