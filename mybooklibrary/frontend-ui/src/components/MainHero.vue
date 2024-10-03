<template>
  <section class="bg-[url('/src/assets/images/header-bg.jpg')] bg-center bg-cover">
    <div class="bg-slate-900/60">
      <div class="py-12 mx-40 text-white">
        <div class="flex justify-center items-center space-x-10">
          <div class="lg:w-[65%] hidden md:inline">
            <h1 class="text-6xl font-bold">Manage your book collections and sharing with other</h1>
            <p class="text-2xl mt-8 pl-8">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="3.0"
                   stroke="currentColor"
                   class="w-8 h-8 mr-4 inline-block bg-green-200/20 text-green-400 rounded-md">
                <path stroke-linecap="round" stroke-linejoin="round" d="m4.5 12.75 6 6 9-13.5" />
              </svg>
              Manage your personal book collection anywhere, anytime
            </p>
            <p class="text-2xl mt-5 pl-8">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="3.0"
                   stroke="currentColor"
                   class="w-8 h-8 mr-4 inline-block bg-green-200/20 text-green-400 rounded-md">
                <path stroke-linecap="round" stroke-linejoin="round" d="m4.5 12.75 6 6 9-13.5" />
              </svg>
              Review other user books and discover more books to read
            </p>
            <p class="text-2xl mt-5 pl-8 whitespace-nowrap">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="3.0"
                   stroke="currentColor"
                   class="w-8 h-8 mr-4 inline-block bg-green-200/20 text-green-400 rounded-md">
                <path stroke-linecap="round" stroke-linejoin="round" d="m4.5 12.75 6 6 9-13.5" />
              </svg>
              Keep track of your books borrowing so that you will never lost it
            </p>
          </div>
          <div>
            <form @submit.prevent="submitHandler" class="bg-green-900/70 p-8 rounded-md">
              <h3 class="text-4xl font-medium mb-3 text-center">Sign Up Today</h3>
              <p class="text-xl text-center mb-5">Please fill up this form to register</p>
              <p class="mb-3">
                <input v-model="firstName" class="rounded-md w-full p-2 text-black" type="text"
                       placeholder="First Name" />
              </p>
              <p class="mb-3">
                <input v-model="lastName" class="rounded-md w-full p-2 text-black" type="text"
                       placeholder="Last Name" /></p>
              <p class="mb-3">
                <input v-model="email" class="rounded-md w-full p-2 text-black" type="email"
                       placeholder="Email" /></p>
              <p class="mb-3">
                <input v-model="password" class="rounded-md w-full p-2 text-black" type="password"
                       placeholder="Password" /></p>
              <p class="mb-5">
                <input v-model="confirmPassword" class="rounded-md w-full p-2 text-black" type="password"
                       placeholder="Confirm Password" /></p>
              <button class="font-medium uppercase border-2 rounded-md p-3 w-full hover:bg-green-800">Register</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import appEnvironment from '@/environment.js'
import router from '@/router/index.js'

export default {
  name: 'MainHero',
  data() {
    return {
      firstName: '',
      lastName: '',
      email: '',
      password: '',
      confirmPassword: ''
    }
  },
  methods: {
    submitHandler() {
      //console.log(this.firstName, this.lastName, this.email, this.password)

      if (this.firstName === '' || this.lastName === '' || this.email === '') {
        notie.alert({
          type: 'error',
          text: 'All fields are required'
        })
        return
      }

      if (this.password === '' || this.confirmPassword === '') {
        notie.alert({
          type: 'error',
          text: 'Password and confirm password are required'
        })
        return
      }

      if (this.password !== this.confirmPassword) {
        notie.alert({
          type: 'error',
          text: 'Password entered does not match'
        })
        return
      }

      const payload = {
        first_name: this.firstName,
        last_name: this.lastName,
        email: this.email,
        password: this.password,
        active: 1,
        is_admin: 0
      }

      const headers = new Headers()

      headers.append('Content-Type', 'application/json')

      const requestOptions = {
        method: 'POST',
        body: JSON.stringify(payload),
        headers: headers,
      }

      fetch(appEnvironment.apiURL() + '/users/register', requestOptions)
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
              text: 'User registered successfully'
            })
          }
        })
        .catch((err) => {
          console.log(err)
        })

      this.firstName = ''
      this.lastName = ''
      this.email = ''
      this.password = ''
      this.confirmPassword = ''

      router.push('/')
    }
  }
}
</script>