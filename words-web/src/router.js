import { createWebHistory, createRouter } from 'vue-router'

import Main from './pages/Main.vue'
import Login from './pages/Login.vue'
import About from './pages/About.vue'

const routes = [
  { path: '/', name: 'Main', component: Main },
  { path: '/about', name: 'About', component: About },
  { path: '/login', name: 'Login', component: Login },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  if (to.name !== 'Login') next({ name: 'Login' })
  else next()
  // next()
})

export default router