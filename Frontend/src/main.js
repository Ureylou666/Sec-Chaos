import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugin/element-ui'
import './assets/css/global.css'
import axios from 'axios'
import { Message } from 'element-ui'

Vue.prototype.$message = Message
Vue.prototype.$http = axios
axios.defaults.baseURL = 'http://127.0.0.1:8080/api/v1/'

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
