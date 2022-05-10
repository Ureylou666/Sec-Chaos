import Vue from 'vue'
import App from './App.vue'
import router from './router'
import { Button, Container, Main, Header, Aside, Menu, Submenu, MenuItemGroup, MenuItem, Radio, RadioGroup, RadioButton, Form, FormItem, Input } from 'element-ui'
import './assets/css/global.css'
import axios from 'axios'

Vue.config.productionTip = false
Vue.use(Button)
Vue.use(Container)
Vue.use(Main)
Vue.use(Header)
Vue.use(Aside)
Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItem)
Vue.use(MenuItemGroup)
Vue.use(Radio)
Vue.use(RadioGroup)
Vue.use(RadioButton)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Input)

axios.defaults.baseUrl = 'http://127.0.0.1:8080'
Vue.prototype.$http = axios

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
