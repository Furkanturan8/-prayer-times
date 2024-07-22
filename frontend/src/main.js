import { createApp } from 'vue'
import App from './App.vue'
import { loadFonts } from './plugins/webfontloader'
import router from '@/router';
import {registerPlugins} from "@/plugins";

loadFonts()

const app = createApp(App)

registerPlugins(app)

app.use(router)

app.mount('#app')