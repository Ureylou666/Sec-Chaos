import Vue from 'vue'
import App from './App.vue'
import router from './router'
import { Button, Container, Main, Header, Aside, Menu, Submenu, MenuItemGroup, MenuItem, Radio, RadioGroup, RadioButton } from 'element-ui'

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

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
