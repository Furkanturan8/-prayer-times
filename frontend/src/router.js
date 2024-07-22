import {createRouter, createWebHistory} from "vue-router";

const routes = [
    { path: '/', name: 'home', component: () => import('@/pages/MainPage.vue') },
    { path: '/prayer-times', name: 'prayer-times', component: () => import('@/pages/PrayerTimesPage.vue') },
    { path: '/cities', name: 'cities', component: () => import('@/pages/CitiesPage.vue') },
    { path: '/contact', name: 'contact', component: () => import('@/pages/ContactPage.vue') },
]

const router = createRouter({
    routes,
    history : createWebHistory(),
})

export default router
