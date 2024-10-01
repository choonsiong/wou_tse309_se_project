<template>
  <div class="bg-green-200">
    <div class="flex justify-between items-center">
      <div class="ml-8">
        <img src="../assets/logo.png" alt="logo" class="cursor-pointer w-56">
      </div>
      <div class="mr-8 hidden md:inline">
        <ul class="flex space-x-2 font-bold font-roboto">
          <li class="py-1 px-3 hover:bg-green-400 hover:rounded-full cursor-pointer" @click="showLogin">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                 stroke="currentColor" class="w-8 h-8 inline-block">
              <path stroke-linecap="round" stroke-linejoin="round"
                    d="M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15m3 0 3-3m0 0-3-3m3 3H9" />
            </svg>
            Login
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MainHeader',
  emits: ['show-login-event'],
  methods: {
    showLogin() {
      console.log('showLogin()')
      this.$emit('show-login-event')
    },
  }
}
</script>