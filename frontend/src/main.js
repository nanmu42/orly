import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import VueProgressiveImage from 'vue-progressive-image'

Vue.use(VueProgressiveImage)

Vue.config.productionTip = false

new Vue({
  render: h => h(App)
}).$mount('#app')
