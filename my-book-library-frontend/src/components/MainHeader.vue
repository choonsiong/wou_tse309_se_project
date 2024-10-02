<template>
  <div class="bg-green-200">
    <div class="flex justify-between items-center">
      <div class="ml-8">
        <img src="../assets/logo.png" alt="logo" class="cursor-pointer w-56">
      </div>
      <div class="mr-8 hidden md:inline">
        <ul class="flex space-x-2 font-bold font-roboto">
          <li v-if="store.isLoggedIn" class="py-1 px-3 hover:bg-green-400 hover:rounded-full cursor-pointer">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8 inline-block">
              <path d="M16 7C16 9.20914 14.2091 11 12 11C9.79086 11 8 9.20914 8 7C8 4.79086 9.79086 3 12 3C14.2091 3 16 4.79086 16 7Z" stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path>
              <path d="M12 14C8.13401 14 5 17.134 5 21H19C19 17.134 15.866 14 12 14Z" stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path>
            </svg>
            {{ store.user.first_name }} {{store.user.last_name}}
          </li>
          <li v-if="!store.isLoggedIn" class="py-1 px-3 hover:bg-green-400 hover:rounded-full cursor-pointer"
              @click="showLogin">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                 stroke="currentColor" class="w-8 h-8 inline-block">
              <path d="M2.00098 11.999L16.001 11.999M16.001 11.999L12.501 8.99902M16.001 11.999L12.501 14.999"
                    stroke="#1C274C" stroke-linecap="round" stroke-linejoin="round"></path>
              <path
                d="M9.00195 7C9.01406 4.82497 9.11051 3.64706 9.87889 2.87868C10.7576 2 12.1718 2 15.0002 2L16.0002 2C18.8286 2 20.2429 2 21.1215 2.87868C22.0002 3.75736 22.0002 5.17157 22.0002 8L22.0002 16C22.0002 18.8284 22.0002 20.2426 21.1215 21.1213C20.3531 21.8897 19.1752 21.9862 17 21.9983M9.00195 17C9.01406 19.175 9.11051 20.3529 9.87889 21.1213C10.5202 21.7626 11.4467 21.9359 13 21.9827"
                stroke="#1C274C" stroke-linecap="round"></path>
            </svg>
            Login
          </li>
          <li v-else class="py-1 px-3 hover:bg-green-400 hover:rounded-full cursor-pointer" @click="logoutUser">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                 stroke="currentColor" class="w-8 h-8 inline-block">
              <path d="M15 12L2 12M2 12L5.5 9M2 12L5.5 15" stroke="#1C274C" stroke-linecap="round"
                    stroke-linejoin="round"></path>
              <path
                d="M9.00195 7C9.01406 4.82497 9.11051 3.64706 9.87889 2.87868C10.7576 2 12.1718 2 15.0002 2L16.0002 2C18.8286 2 20.2429 2 21.1215 2.87868C22.0002 3.75736 22.0002 5.17157 22.0002 8L22.0002 16C22.0002 18.8284 22.0002 20.2426 21.1215 21.1213C20.3531 21.8897 19.1752 21.9862 17 21.9983M9.00195 17C9.01406 19.175 9.11051 20.3529 9.87889 21.1213C10.5202 21.7626 11.4467 21.9359 13 21.9827"
                stroke="#1C274C" stroke-linecap="round"></path>
            </svg>
            Logout
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import { store } from '@/components/store.js'
import notie from 'notie'
import appEnvironment from '@/environment.js'

export default {
  name: 'MainHeader',
  data() {
    return {
      store
    }
  },
  emits: ['show-login-event'],
  methods: {
    showLogin() {
      //console.log('showLogin()')
      this.$emit('show-login-event')
    },
    logoutUser() {
      //console.log('logoutUser()')
      const payload = {
        token: store.token
      }
      const requestOption = {
        method: 'POST',
        body: JSON.stringify(payload)
      }
      console.log(appEnvironment().apiURL)
      fetch(appEnvironment().apiURL + '/users/logout', requestOption)
        .then((resp) => resp.json())
        .then((jsonResp) => {
          if (jsonResp.error) {
            console.log('error: ' + jsonResp.message)
          } else {
            store.isLoggedIn = false
            store.token = ''
            store.user = {}
            document.cookie = '_mybooklibrary_data=; Path=/; SameSite=Strict; Secure; Expires=Thu, 01 Jan 1970 00:00:01 GMT;'
            notie.alert({
              type: 'success',
              text: 'You are logged out'
            })
            //router.push('/home')
          }
        })
    }
  }
}
</script>