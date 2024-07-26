import { createApp } from 'vue'
import App from './App.vue'
import { loadFonts } from './plugins/webfontloader'
import router from '@/router';
import {registerPlugins} from "@/plugins";
import { faGithub, faTwitter, faInstagram, faWhatsapp } from '@fortawesome/free-brands-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {library} from "@fortawesome/fontawesome-svg-core";

library.add(faGithub, faTwitter, faInstagram, faWhatsapp)


loadFonts()

const app = createApp(App)

registerPlugins(app)

app.use(router)
app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')