import {createRouter, createWebHistory} from "vue-router";
import PrayerTimesPage from "@/pages/PrayerTimesPage.vue";
import MainPage from "@/pages/MainPage.vue";
import CitiesPage from "@/pages/CitiesPage.vue";
import ContactPage from "@/pages/ContactPage.vue";

const routes = [
    { path: '/', name: 'home', component: (MainPage) },
    { path: '/prayer-times', name: 'prayer-times', component: (PrayerTimesPage)  },
    { path: '/cities', name: 'cities', component: (CitiesPage)},
    { path: '/contact', name: 'contact', component: (ContactPage) },
]

const router = createRouter({
    routes,
    history : createWebHistory(),
})

export default router
console.log('Route Params:', routes.params);