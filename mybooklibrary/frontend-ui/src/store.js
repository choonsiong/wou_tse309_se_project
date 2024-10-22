import { reactive } from 'vue'

export const store = reactive({
  token: '',
  isAdmin: false,
  isLoggedIn: false,
  user: {},
  book: {}
})