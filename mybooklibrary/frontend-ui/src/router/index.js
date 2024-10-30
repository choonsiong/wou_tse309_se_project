import { createRouter, createWebHistory } from 'vue-router'
import MainContent from '@/components/MainContent.vue'
import UserBooks from '@/components/UserBooks.vue'
import UserProfile from '@/components/UserProfile.vue'
import AllUsers from '@/components/AllUsers.vue'
import AllBooks from '@/components/AllBooks.vue'
import AddBook from '@/components/AddBook.vue'
import EditBook from '@/components/EditBook.vue'
import BookDetail from '@/components/BookDetail.vue'
import AddReview from '@/components/AddReview.vue'
import AllReviews from '@/components/AllReviews.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: MainContent
  },
  {
    path: '/books/:id',
    name: 'BookDetail',
    component: BookDetail,
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
    path: '/manage/reviews/new',
    name: 'AddReview',
    component: AddReview,
  },
  {
    path: '/admin/manage/users',
    name: 'AllUsers',
    component: AllUsers
  },
  {
    path: '/admin/manage/reviews',
    name: 'AllReviews',
    component: AllReviews
  },
  {
    path: '/admin/manage/books',
    name: 'AllBooks',
    component: AllBooks
  },
  {
    path: '/admin/manage/books/new',
    name: 'AddBook',
    component: AddBook
  },
  {
    path: '/admin/manage/books/edit/:id',
    name: 'EditBook',
    component: EditBook
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