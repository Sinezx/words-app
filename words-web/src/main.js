import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import router from './router.js'

//mock data
// import './mock/index'

const app = createApp(App)
app.use(ElementPlus)
app.use(router)
app.mount('#app')
