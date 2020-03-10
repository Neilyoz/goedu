import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import http from './http'

// 加载自定义样式
import 'normalize.css'
import '@/assets/style/style.scss'

Vue.prototype.$http = http

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
