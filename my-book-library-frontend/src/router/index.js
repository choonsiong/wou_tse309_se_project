import { createRouter, createWebHistory } from 'vue-router'
import MainContent from "@/components/MainContent.vue"
import UserBooks from '@/components/UserBooks.vue'
import UserProfile from '@/components/UserProfile.vue'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: MainContent,
    },
    {
        path: '/manage/books',
        name: 'UserBooks',
        component: UserBooks,
    },
    {
        path: '/manage/profile',
        name: 'UserProfile',
        component: UserProfile,
    },
]

const router = createRouter({
    history: createWebHistory(), routes
})

export default router