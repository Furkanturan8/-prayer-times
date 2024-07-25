import {createRouter, createWebHistory} from "vue-router";
import MainPage from "@/pages/MainPage.vue";
import CitiesPage from "@/pages/CitiesPage.vue";
import ContactPage from "@/pages/ContactPage.vue";
import PrayerTimesPage from "@/pages/PrayerTimesPage.vue";
import PrayerTimesMonthPage from "@/pages/PrayerTimesMonthPage.vue";

const routes = [
    { path: '/', name: 'home', component: (MainPage) },
    { path: '/prayer-times', name: 'prayer-times', component: (PrayerTimesPage)  },
    { path: '/prayer-times/:month', name: 'prayer-times-month', component: (PrayerTimesMonthPage)  },
    { path: '/cities', name: 'cities', component: (CitiesPage)},
    { path: '/contact', name: 'contact', component: (ContactPage) },
]

const router = createRouter({
    routes,
    history : createWebHistory(),
})

export default router
console.log('Route Params:', routes.params);