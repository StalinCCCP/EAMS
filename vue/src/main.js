import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import * as bootstrap from "bootstrap"
// import BootstrapTable from '@/plugins/table'

import "@/style.scss"
createApp(App).use(store).use(router).mount('#app')
