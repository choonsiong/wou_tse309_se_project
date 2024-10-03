<template>
  <section class="bg-green-100">
    <h2 class="pt-10 text-4xl mb-6 font-bold text-center">User Profile</h2>
    <div class="max-w-2xl mx-auto p-6 space-y-6">
      <div>
        <form @submit.prevent="submitEditForm"
              class="flex flex-col items-center justify-between w-full p-6 bg-green-500/30 rounded-lg md:flex-col">
          <input v-model="firstName" type="text" placeholder="First Name"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <input v-model="lastName" type="text" placeholder="Last Name"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <input v-model="email" type="email" placeholder="Email"
                 class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          <div class="w-full flex-1">
            <p class="mb-2">Leave the password field empty to use existing password.</p>
            <input v-model="password" type="password" placeholder="Password"
                   class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          </div>
          <div class="w-full flex-1">
            <p class="mb-2">Leave the confirm password field empty to use existing password.</p>
            <input v-model="confirmPassword" type="password" placeholder="Confirm Password"
                   class="w-full flex-1 px-6 pt-3 pb-2 mb-4 rounded-lg border-1 border-white focus:outline-none" />
          </div>
          <div class="flex flex-row justify-between w-full gap-4 mt-10 mb-5">
            <button class="w-full p-3 px-8 text-white bg-blue-600 rounded-lg hover:opacity-70 focus:outline-none"
                    @click="cancelEdit">Cancel
            </button>
            <button class="w-full p-3 px-8 text-white bg-red rounded-lg hover:opacity-70 focus:outline-none">Submit
            </button>
          </div>
        </form>
      </div>
    </div>
    <div class="p-10">&nbsp;</div>
  </section>
</template>

<script>
import Security from '@/security.js'
import router from '@/router/index.js'
import { store } from '@/store.js'
import appEnvironment from '@/environment.js'

export default {
  name: 'UserProfile',
  data() {
    return {
      firstName: '',
      lastName: '',
      email: '',
      password: '',
      confirmPassword: '',
      isActive: 0,
      isAdmin: 0,
    }
  },
  methods: {
    submitEditForm() {
      console.log(this.firstName, this.lastName, this.email, this.password)

      if (this.password !== this.confirmPassword) {
        notie.alert({
          type: 'error',
          text: 'password entered does not match'
        })
        return
      }

      const payload = {
        id: parseInt(String(store.user.id), 10),
        first_name: this.firstName,
        last_name: this.lastName,
        email: this.email,
        password: this.password,
        active: this.isActive,
        is_admin: this.isAdmin,
      }

      fetch(appEnvironment.apiURL() + '/admin/users/edit', Security.requestOptions(payload))
        .then((resp) => resp.json())
        .then((jsonResp) => {
          if (jsonResp.error) {
            notie.alert({
              type: 'error',
              text: jsonResp.message
            })
          } else {
            notie.alert({
              type: 'success',
              text: 'user saved successfully'
            })
          }
        })
        .catch((err) => {
          console.log(err)
        })

      router.push('/')
    },
    cancelEdit() {
      router.push('/')
    }
  },
  beforeMount() {
    Security.requireToken()

    // console.log(store.user.id)

    fetch(appEnvironment.apiURL() + '/admin/users/get/' + store.user.id, Security.requestOptions(''))
    .then((resp) => resp.json())
    .then((jsonResp) => {
      if (jsonResp.error) {
        notie.alert({
          type: 'error',
          text: jsonResp.message
        })
      } else {
        //console.log(jsonResp)
        this.firstName = jsonResp.first_name
        this.lastName = jsonResp.last_name
        this.email = jsonResp.email
        this.password = ''
        this.confirmPassword = ''
        this.isActive = jsonResp.active
        this.isAdmin = jsonResp.is_admin
      }
    })
  }
}
</script>