<template>
  <div class="bg-green-200">
    <div class="flex justify-between items-center">
      <div class="ml-8">
        <router-link to="/"><img src="../assets/logo.png" alt="logo" class="cursor-pointer w-32 md:w-56"></router-link>
      </div>
      <div class="mr-8 inline">
        <ul class="flex space-x-2 font-bold font-roboto">
          <li v-if="store.isLoggedIn && store.isAdmin"
              class="py-1 px-3 hover:bg-green-400 hover:rounded-full cursor-pointer">
            <router-link to="/admin/manage/users">
              <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" stroke-width="1.5"
                   class="w-8 h-8 inline-block">
                <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
                <g id="SVGRepo_iconCarrier">
                  <g clip-path="url(#clip0_1251_98416)">
                    <path fill-rule="evenodd" clip-rule="evenodd"
                          d="M9 0C5.96243 0 3.5 2.46243 3.5 5.5C3.5 8.53757 5.96243 11 9 11C12.0376 11 14.5 8.53757 14.5 5.5C14.5 2.46243 12.0376 0 9 0ZM5.5 5.5C5.5 3.567 7.067 2 9 2C10.933 2 12.5 3.567 12.5 5.5C12.5 7.433 10.933 9 9 9C7.067 9 5.5 7.433 5.5 5.5Z"
                          fill="#000000"></path>
                    <path
                      d="M15.5 0C14.9477 0 14.5 0.447715 14.5 1C14.5 1.55228 14.9477 2 15.5 2C17.433 2 19 3.567 19 5.5C19 7.433 17.433 9 15.5 9C14.9477 9 14.5 9.44771 14.5 10C14.5 10.5523 14.9477 11 15.5 11C18.5376 11 21 8.53757 21 5.5C21 2.46243 18.5376 0 15.5 0Z"
                      fill="#000000"></path>
                    <path
                      d="M19.0837 14.0157C19.3048 13.5096 19.8943 13.2786 20.4004 13.4997C22.5174 14.4246 24 16.538 24 19V21C24 21.5523 23.5523 22 23 22C22.4477 22 22 21.5523 22 21V19C22 17.3613 21.0145 15.9505 19.5996 15.3324C19.0935 15.1113 18.8625 14.5217 19.0837 14.0157Z"
                      fill="#000000"></path>
                    <path
                      d="M6 13C2.68629 13 0 15.6863 0 19V21C0 21.5523 0.447715 22 1 22C1.55228 22 2 21.5523 2 21V19C2 16.7909 3.79086 15 6 15H12C14.2091 15 16 16.7909 16 19V21C16 21.5523 16.4477 22 17 22C17.5523 22 18 21.5523 18 21V19C18 15.6863 15.3137 13 12 13H6Z"
                      fill="#000000"></path>
                  </g>
                  <defs>
                    <clipPath id="clip0_1251_98416">
                      <rect width="24" height="24" fill="white"></rect>
                    </clipPath>
                  </defs>
                </g>
              </svg>
              Manage Users
            </router-link>
          </li>
          <li v-if="store.isLoggedIn && store.isAdmin"
              class="py-1 px-3 hover:bg-green-400 hover:rounded-full cursor-pointer">
            <router-link to="/admin/manage/books">
            <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg"
                 xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 32 32" enable-background="new 0 0 32 32"
                 xml:space="preserve" fill="#000000" class="w-8 h-8 inline-block"><g id="SVGRepo_bgCarrier" stroke-width="0" ></g>
              <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
              <g id="SVGRepo_iconCarrier"> <rect x="6" y="11" fill="none" stroke="#000000" stroke-width="2"
                                                 stroke-miterlimit="10" width="20" height="13"></rect>
                <rect x="4" y="24" fill="none" stroke="#000000" stroke-width="2" stroke-miterlimit="10" width="24"
                      height="3"></rect>
                <polygon fill="none" stroke="#000000" stroke-width="2" stroke-miterlimit="10"
                         points="28,10 16,5 4,10 4,11 28,11 "></polygon>
                <line fill="none" stroke="#000000" stroke-width="2" stroke-miterlimit="10" x1="18" y1="24" x2="18"
                      y2="11"></line>
                <line fill="none" stroke="#000000" stroke-width="2" stroke-miterlimit="10" x1="22" y1="24" x2="22"
                      y2="11"></line>
                <line fill="none" stroke="#000000" stroke-width="2" stroke-miterlimit="10" x1="10" y1="24" x2="10"
                      y2="11"></line>
                <line fill="none" stroke="#000000" stroke-width="2" stroke-miterlimit="10" x1="14" y1="24" x2="14"
                      y2="11"></line> </g></svg>
            Manage Books
            </router-link>
          </li>
          <li v-if="store.isLoggedIn && !store.isAdmin" class="py-1 px-3 hover:bg-green-400 hover:rounded-full cursor-pointer">
            <router-link to="/manage/books">
              <svg viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg" fill="#000000" class="w-8 h-8 inline-block">
                <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
                <g id="SVGRepo_iconCarrier"><title>ionicons-v5-l</title>
                  <rect x="32" y="96" width="64" height="368" rx="16" ry="16"
                        style="fill:none;stroke:#000000;stroke-linejoin:round;stroke-width:32px"></rect>
                  <line x1="112" y1="224" x2="240" y2="224"
                        style="fill:none;stroke:#000000;stroke-linecap:round;stroke-linejoin:round;stroke-width:32px"></line>
                  <line x1="112" y1="400" x2="240" y2="400"
                        style="fill:none;stroke:#000000;stroke-linecap:round;stroke-linejoin:round;stroke-width:32px"></line>
                  <rect x="112" y="160" width="128" height="304" rx="16" ry="16"
                        style="fill:none;stroke:#000000;stroke-linejoin:round;stroke-width:32px"></rect>
                  <rect x="256" y="48" width="96" height="416" rx="16" ry="16"
                        style="fill:none;stroke:#000000;stroke-linejoin:round;stroke-width:32px"></rect>
                  <path
                    d="M422.46,96.11l-40.4,4.25c-11.12,1.17-19.18,11.57-17.93,23.1l34.92,321.59c1.26,11.53,11.37,20,22.49,18.84l40.4-4.25c11.12-1.17,19.18-11.57,17.93-23.1L445,115C443.69,103.42,433.58,94.94,422.46,96.11Z"
                    style="fill:none;stroke:#000000;stroke-linejoin:round;stroke-width:32px"></path>
                </g>
              </svg>
              My Books
            </router-link>
          </li>
          <li v-if="store.isLoggedIn" class="py-1 px-3 hover:bg-green-400 hover:rounded-full cursor-pointer">
            <router-link to="/manage/profile">
              <svg v-if="store.isAdmin" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" fill="#000000" class="w-8 h-8 inline-block">
                <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
                <g id="SVGRepo_iconCarrier">
                  <g>
                    <path fill="none" d="M0 0h24v24H0z"></path>
                    <path
                      d="M12 14v2a6 6 0 0 0-6 6H4a8 8 0 0 1 8-8zm0-1c-3.315 0-6-2.685-6-6s2.685-6 6-6 6 2.685 6 6-2.685 6-6 6zm0-2c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm9 6h1v5h-8v-5h1v-1a3 3 0 0 1 6 0v1zm-2 0v-1a1 1 0 0 0-2 0v1h2z"></path>
                  </g>
                </g>
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                   stroke="currentColor" class="w-8 h-8 inline-block">
                <path
                  d="M16 7C16 9.20914 14.2091 11 12 11C9.79086 11 8 9.20914 8 7C8 4.79086 9.79086 3 12 3C14.2091 3 16 4.79086 16 7Z"
                  stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path>
                <path d="M12 14C8.13401 14 5 17.134 5 21H19C19 17.134 15.866 14 12 14Z" stroke="#000000"
                      stroke-width="2"
                      stroke-linecap="round" stroke-linejoin="round"></path>
              </svg>
              <span v-if="store.isAdmin">Admin</span>
              <span v-else>Profile</span>
<!--              {{ store.user.first_name }} {{ store.user.last_name }}-->
            </router-link>
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
import { store } from '@/store.js'
import notie from 'notie'
import appEnvironment from '@/environment.js'
import router from '@/router/index.js'
import Security from '@/security.js'

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
      this.$emit('show-login-event')
    },
    logoutUser() {
      const payload = {
        token: store.token
      }
      fetch(appEnvironment.apiURL() + '/users/logout', Security.requestOptions(payload))
        .then((resp) => resp.json())
        .then((jsonResp) => {
          if (jsonResp.error) {
            console.log('error: ' + jsonResp.message)
          } else {
            store.isLoggedIn = false
            store.isAdmin = false
            store.token = ''
            store.user = {}
            document.cookie = '_mybooklibrary_data=; Path=/; SameSite=Strict; Secure; Expires=Thu, 01 Jan 1970 00:00:01 GMT;'
            notie.alert({
              type: 'success',
              text: 'You are logged out'
            })
          }
        })
      router.replace("/")
    }
  }
}
</script>