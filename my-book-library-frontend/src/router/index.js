import { createRouter, createWebHistory } from 'vue-router'
import MainContent from "@/components/MainContent.vue"

const routes = [
    {
        path: '/',
        name: 'Home',
        component: MainContent,
    }
]

const router = createRouter({
    history: createWebHistory(), routes
})

export default router