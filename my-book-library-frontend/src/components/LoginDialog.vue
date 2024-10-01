<template>
  <teleport to="body">
    <div @click="$emit('close-event')"></div>
    <dialog open>
      <section>
        <form class="bg-green-900 p-8" @submit.prevent="loginUser">
          <h3 class="text-4xl font-medium mb-5 text-center text-white">Login Now</h3>
          <p
            class="mb-3"><input class="rounded-md w-full p-2 text-black" type="email" placeholder="Enter your email"
                                v-model="userEmail"
          /></p>
          <p
            class="mb-5"><input class="rounded-md w-full p-2 text-black" type="password"
                                placeholder="Enter your password"
                                v-model="userPassword" /></p>
          <button class="font-medium uppercase border-2 rounded-md p-3 w-full hover:bg-green-800 text-white">Login</button>
        </form>
      </section>
    </dialog>
  </teleport>
</template>

<script>
export default {
  data() {
    return {
      userEmail: '',
      userPassword: ''
    }
  },
  emits: ['close-event', 'login-event'],
  methods: {
    loginUser() {
      this.$emit('login-event', this.userEmail, this.userPassword)
    }
  }
}
</script>

<style scoped>
div {
  position: fixed;
  top: 0;
  left: 0;
  height: 100vh;
  width: 100%;
  background-color: rgba(0, 0, 0, 0.75);
  z-index: 10;
}

dialog {
  position: fixed;
  top: 20vh;
  left: 10%;
  width: 80%;
  z-index: 100;
  //border-radius: 12px;
  border: none;
  //box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26);
  padding: 0;
  margin: 0;
  overflow: hidden;
}

@media (min-width: 768px) {
  dialog {
    left: calc(50% - 20rem);
    width: 40rem;
  }
}
</style>