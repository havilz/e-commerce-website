import { createApp } from 'vue'
import { createPinia } from 'pinia'
import Toast, { POSITION } from 'vue-toastification'

import App from './App.vue'
import router from './router'
import '@/assets/css/main.css'
import 'vue-toastification/dist/index.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Toast, {
  position:         POSITION.TOP_RIGHT,
  timeout:          3000,
  closeOnClick:     true,
  pauseOnHover:     true,
  draggable:        true,
  showCloseButtonOnHover: false,
  hideProgressBar:  false,
  maxToasts:        5,
})

app.mount('#app')
