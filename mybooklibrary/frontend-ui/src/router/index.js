import { createRouter, createWebHistory } from 'vue-router'
import MainContent from '@/components/MainContent.vue'
import UserBooks from '@/components/UserBooks.vue'
import UserProfile from '@/components/UserProfile.vue'
import AllUsers from '@/components/AllUsers.vue'
import AllBooks from '@/components/AllBooks.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: MainContent
  },
  {
    path: '/manage/books',
    name: 'UserBooks',
    component: UserBooks
  },
  {
    path: '/manage/profile',
    name: 'UserProfile',
    component: UserProfile
  },
  {
    path: '/admin/manage/users',
    name: 'AllUsers',
    component: AllUsers
  },
  {
    path: '/admin/manage/books',
    name: 'AllBooks',
    component: AllBooks
  },
  {
    path: '/:notFound(.*)',
    //component: NotFound
    redirect: '/'
  }
]

const router = createRouter({
  history: createWebHistory(), routes
})

export default router