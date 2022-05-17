import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugin/element-ui'
import './assets/css/global.css'
import axios from 'axios'

axios.defaults.baseUrl = 'http://127.0.0.1:8080'
Vue.prototype.$http = axios

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
