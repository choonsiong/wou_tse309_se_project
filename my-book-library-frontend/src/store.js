import { reactive } from 'vue'

export const store = reactive({
  token: '',
  isAdmin: true,
  isLoggedIn: false,
  user: {}
})