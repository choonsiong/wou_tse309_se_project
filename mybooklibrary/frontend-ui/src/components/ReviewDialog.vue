<template>
  <teleport to="body">
    <div id="closeDiv" @click="$emit('close-event')"></div>
    <dialog open>
      <section class="flex items-center justify-center">
        <div class="flex flex-col my-6 bg-green-50 shadow-sm border border-slate-200 rounded-lg w-[600px]">
          <div class="mx-3 mb-0 border-b border-slate-200 pt-3 pb-2 px-1 text-right">
            <span class="text-sm font-medium text-slate-600 cursor-pointer" @click="this.$emit('close-event')">X</span>
          </div>
          <h1 class="mt-5 text-center text-2xl font-bold">Write a Review for<br/> {{ book.title }}</h1>
          <div class="m-2.5 overflow-hidden rounded-md h-80 flex justify-center items-center">
            <img class="w-full h-full object-scale-down" :src="appEnvironment.imageURL() + '/' + book.slug + '.jpg'" alt="profile-picture" />
          </div>
          <div class="p-4 relative max-w-md w-full min-w-[500px]">
            <label class="block mb-2 text-sm text-slate-600">Book Review:</label>
            <textarea class="w-full bg-white placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                      placeholder="Enter book review" rows="10"></textarea>
          </div>
          <div class="p-4 mb-5">
            <label class="mb-2 text-sm text-slate-600">Rating: </label>
            <select>
              <option>1</option>
              <option>2</option>
              <option>3</option>
              <option>4</option>
              <option>5</option>
            </select>
          </div>
          <div class="text-center mb-5 mt-5">
            <button class="p-3 px-8 text-white bg-green-500 rounded-lg hover:opacity-70 focus:outline-none">Submit</button>
          </div>
        </div>
      </section>
    </dialog>
  </teleport>
</template>

<script>
import appEnvironment from '@/environment.js'
import { store } from '@/store.js'

export default {
  name: 'ReviewDialog',
  emits: ['close-event'],
  props: ['book'],
  computed: {
    store() {
      return store
    },
    appEnvironment() {
      return appEnvironment
    },
  },
}
</script>

<style scoped>
#closeDiv {
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
  top: 5vh;
  left: 10%;
  width: 80%;
  z-index: 100;
  //border-radius: 12px;
  //border: none;
  background-color: transparent;
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