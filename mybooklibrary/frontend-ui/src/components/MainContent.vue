<template>
  <main-hero v-if="!store.isLoggedIn"></main-hero>
  <main class="bg-green-50 m-0 p-10">
    <div class="grid grid-cols-4 grid-flow-row gap-3">
      <div
        class="w-[320px] mb-5 bg-white rounded-sm mx-auto shadow-md hover:bg-green-200 hover:-translate-y-3 cursor-pointer">
        <img class="w-[320px]" src="../assets/images/object-oriented-java-script.png">
        <div class="mt-3 p-3">
          <h3 class="font-bold text-center">The Principles of Object-Oriented Java Script</h3>
          <p class="pt-2 font-light text-center">by Teiva Harsanyi</p>
          <p class="pt-2">Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
        </div>
      </div>
      <div
        class="w-[320px] mb-5 bg-white rounded-sm mx-auto shadow-md hover:bg-green-200 hover:-translate-y-3 cursor-pointer">
        <img class="w-[320px]" src="../assets/images/object-oriented-java-script.png">
        <div class="mt-3 p-3">
          <h3 class="font-bold text-center">The Principles of Object-Oriented Java Script</h3>
          <p class="pt-2 font-light text-center">by Teiva Harsanyi</p>
          <p class="pt-2">Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
        </div>
      </div>
      <div
        class="w-[320px] mb-5 bg-white rounded-sm mx-auto shadow-md hover:bg-green-200 hover:-translate-y-3 cursor-pointer">
        <img class="w-[320px]" src="../assets/images/object-oriented-java-script.png">
        <div class="mt-3 p-3">
          <h3 class="font-bold text-center">The Principles of Object-Oriented Java Script</h3>
          <p class="pt-2 font-light text-center">by Teiva Harsanyi</p>
          <p class="pt-2">Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
        </div>
      </div>
      <div
        class="w-[320px] mb-5 bg-white rounded-sm mx-auto shadow-md hover:bg-green-200 hover:-translate-y-3 cursor-pointer">
        <img class="w-[320px]" src="../assets/images/object-oriented-java-script.png">
        <div class="mt-3 p-3">
          <h3 class="font-bold text-center">The Principles of Object-Oriented Java Script</h3>
          <p class="pt-2 font-light text-center">by Teiva Harsanyi</p>
          <p class="pt-2">Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
        </div>
      </div>
      <div
        class="w-[320px] mb-5 bg-white rounded-sm mx-auto shadow-md hover:bg-green-200 hover:-translate-y-3 cursor-pointer">
        <img class="w-[320px]" src="../assets/images/object-oriented-java-script.png">
        <div class="mt-3 p-3">
          <h3 class="font-bold text-center">The Principles of Object-Oriented Java Script</h3>
          <p class="pt-2 font-light text-center">by Teiva Harsanyi</p>
          <p class="pt-2">Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
        </div>
      </div>
      <div
        class="w-[320px] mb-5 bg-white rounded-sm mx-auto shadow-md hover:bg-green-200 hover:-translate-y-3 cursor-pointer">
        <img class="w-[320px]" src="../assets/images/object-oriented-java-script.png">
        <div class="mt-3 p-3">
          <h3 class="font-bold text-center">The Principles of Object-Oriented Java Script</h3>
          <p class="pt-2 font-light text-center">by Teiva Harsanyi</p>
          <p class="pt-2">Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
        </div>
      </div>
      <div
        class="w-[320px] mb-5 bg-white rounded-sm mx-auto shadow-md hover:bg-green-200 hover:-translate-y-3 cursor-pointer">
        <img class="w-[320px]" src="../assets/images/object-oriented-java-script.png">
        <div class="mt-3 p-3">
          <h3 class="font-bold text-center">The Principles of Object-Oriented Java Script</h3>
          <p class="pt-2 font-light text-center">by Teiva Harsanyi</p>
          <p class="pt-2">Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
        </div>
      </div>
      <div
        class="w-[320px] mb-5 bg-white rounded-sm mx-auto shadow-md hover:bg-green-200 hover:-translate-y-3 cursor-pointer">
        <img class="w-[320px]" src="../assets/images/object-oriented-java-script.png">
        <div class="mt-3 p-3">
          <h3 class="font-bold text-center">The Principles of Object-Oriented Java Script</h3>
          <p class="pt-2 font-light text-center">by Teiva Harsanyi</p>
          <p class="pt-2">Lorem ipsum dolor sit amet, consectetur adipisicing elit...</p>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
import MainHero from '@/components/MainHero.vue'
import { store } from '@/store.js'

export default {
  name: 'MainContent',
  components: {
    MainHero
  },
  data() {
    return {
      store
    }
  }
}
</script>