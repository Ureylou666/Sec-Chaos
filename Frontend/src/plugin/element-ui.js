import Vue from 'vue'
import {
  Button,
  Container,
  Main,
  Header,
  Aside,
  Menu,
  Submenu,
  Footer,
  MenuItem,
  Radio,
  RadioGroup,
  RadioButton,
  Form,
  FormItem,
  Input,
  Avatar,
  Message
} from 'element-ui'

Vue.config.productionTip = false
Vue.use(Button)
Vue.use(Container)
Vue.use(Main)
Vue.use(Header)
Vue.use(Aside)
Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItem)
Vue.use(Footer)
Vue.use(Radio)
Vue.use(RadioGroup)
Vue.use(RadioButton)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Input)
Vue.use(Avatar)

Vue.prototype.$message = Message
