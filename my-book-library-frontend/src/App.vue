<template>
  <login-dialog v-if="showLoginDialog"
                @close-event="handleCloseEvent"
                @login-event="handleLoginUserEvent">
  </login-dialog>
  <main-header @show-login-event="handleShowLoginEvent"></main-header>
  <div>
    <router-view />
  </div>
  <main-footer></main-footer>
</template>

<script>
import MainHeader from '@/components/MainHeader.vue'
import MainFooter from '@/components/MainFooter.vue'
import LoginDialog from '@/components/LoginDialog.vue'

import {store} from '@/components/store.js'

export default {
  name: 'App',
  components: {
    LoginDialog,
    MainHeader,
    MainFooter
  },
  data() {
    return {
      showLoginDialog: false,
      store,
    }
  },
  methods: {
    handleCloseEvent() {
      this.showLoginDialog = false
    },
    handleShowLoginEvent() {
      console.log('handleShowLoginEvent()')
      this.showLoginDialog = true
    },
    handleLoginUserEvent(userEmail, userPassword) {
      console.log('handleLoginUserEvent()')
      console.log('Email: ' + userEmail)
      console.log('Password: ' + userPassword)
      this.showLoginDialog = false

      const payload = {
        email: userEmail,
        password: userPassword
      }

      const requestOptions = {
        method: 'POST',
        body: JSON.stringify(payload)
      }

      fetch('http://localhost:9009/users/login', requestOptions)
        .then((resp) => resp.json())
        .then((jsonResp) => {
          if (jsonResp.error) {
            //console.log(jsonResp.error)
            console.log('error: ', jsonResp.message)
          } else {
            //console.log(jsonResp)
            console.log('token: ', jsonResp.data.token.token)
            store.token = jsonResp.data.token.token
            store.isLoggedIn = true
          }
        })
    }
  }
}
</script>
