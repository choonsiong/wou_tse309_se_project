import { reactive } from 'vue'

export const store = reactive({
  token: '',
  isLoggedIn: false,
  user: {}
})