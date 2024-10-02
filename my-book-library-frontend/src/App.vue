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

import {store} from '@/store.js'
import notie from 'notie'
import appEnvironment from '@/environment.js'
import Security from '@/security.js'

const getCookie = (name) => {
  return document.cookie.split('; ').reduce((r, v) => {
    const parts = v.split('=');
    return parts[0] === name ? decodeURIComponent(parts[1]) : r
  }, "")
}

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
      //console.log('handleShowLoginEvent()')
      this.showLoginDialog = true
    },
    handleLoginUserEvent(userEmail, userPassword) {
      //console.log('handleLoginUserEvent()')
      //console.log('Email: ' + userEmail)
      //console.log('Password: ' + userPassword)

      this.showLoginDialog = false

      const payload = {
        email: userEmail,
        password: userPassword
      }

      // const requestOptions = {
      //   method: 'POST',
      //   body: JSON.stringify(payload)
      // }

      //console.log(appEnvironment.apiURL())

      fetch(appEnvironment.apiURL() + '/users/login', Security.requestOptions(payload))
        .then((resp) => resp.json())
        .then((jsonResp) => {
          if (jsonResp.error) {
            //console.log(jsonResp.error)
            //console.log('error: ', jsonResp.message)
            notie.alert({
              type: 'error',
              text: 'Error: ' + jsonResp.message,
            })
          } else {
            //console.log(jsonResp)
            //console.log('token: ', jsonResp.data.token.token)
            store.token = jsonResp.data.token.token
            store.user = {
              id: jsonResp.data.user.id,
              first_name: jsonResp.data.user.first_name,
              last_name: jsonResp.data.user.last_name,
              email: jsonResp.data.user.email,
              is_admin: jsonResp.data.user.is_admin === 1,
            }
            store.isLoggedIn = true
            store.isAdmin = jsonResp.data.user.is_admin === 1

            // save user logged in info to cookie
            let date = new Date()
            let expiredDays = 1
            date.setTime(date.getTime() + (expiredDays * 24 * 60 * 60 * 1000))
            const expires = "expires=" + date.toUTCString()
            document.cookie = "_mybooklibrary_data=" + JSON.stringify(jsonResp.data) + "; " + expires + "; path=/; SameSite=strict; Secure;"

            //router.push('/home')
            notie.alert({
              type: 'success',
              text: 'Logged in successfully',
            })
          }
        })
    }
  },
  beforeMount() {
    // check if a cookie exists?
    let cookieData = getCookie('_mybooklibrary_data');
    if (cookieData !== "") {
      let data = JSON.parse(cookieData);
      store.token = data.token.token;
      store.user = {
        id: data.user.id,
        first_name: data.user.first_name,
        last_name: data.user.last_name,
        email: data.user.email,
        is_admin: data.user.is_admin,
      }
      store.isLoggedIn = true
      store.isAdmin = data.user.isAdmin
    }
  },
  mounted() {
    // const payload = {
    //   data: 'test',
    // }
    // const headers = new Headers()
    // headers.append("Content-Type", "application/json")
    // headers.append("Authorization", "Bearer " + store.token)
    // const requestOption = {
    //   method: 'POST',
    //   headers: headers,
    //   body: JSON.stringify(payload)
    // }
    //
    // fetch('http://localhost:9009/admin/test', requestOption)
    //   .then((resp) => resp.json())
    //   .then((jsonResp) => {
    //     if (jsonResp.error) {
    //       console.log('error: ' + jsonResp.message)
    //     } else {
    //       console.log(jsonResp)
    //     }
    //   })
  }
}
</script>
