import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import VueProgressiveImage from 'vue-progressive-image'
import i18n from './i18n'

Vue.use(VueProgressiveImage)

Vue.config.productionTip = false

new Vue({
  i18n,
  render: h => h(App)
}).$mount('#app')
