<template>
  <teleport to="body">
    <div id="closeDiv" @click="$emit('close-event')"></div>
    <dialog open>
      <section class="flex items-center justify-center">
        <div class="flex flex-col my-6 bg-green-50 shadow-sm border border-slate-200 rounded-lg w-[600px]">
          <div class="mx-3 mb-0 border-b border-slate-200 pt-3 pb-2 px-1 text-right">
            <span class="text-sm font-medium text-slate-600 cursor-pointer" @click="this.$emit('close-event')">X</span>
          </div>
          <div class="m-2.5 overflow-hidden rounded-md h-80 flex justify-center items-center">
            <img class="w-full h-full object-scale-down" :src="appEnvironment.imageURL() + '/' + book.slug + '.jpg'" alt="profile-picture" />
          </div>
          <div class="p-4">
            <h5 class="mb-2 text-slate-800 text-2xl font-semibold">{{ book.title }} <span class="font-bold text-xl">({{ book.publication_year }})</span></h5>
            <div class="pb-3">
              <span class="text-slate-600 text-md font-bold">{{ allAuthors }}</span>
            </div>
            <div class="border-b mb-2 pb-2">
              <span class="font-bold text-xl">{{ normalizedPublisherName }}</span>
            </div>
            <p class="pt-3 text-slate-600 leading-normal font-light">{{ book.description }}</p>
          </div>
          <div class="flex flex-wrap mx-3 border-t border-slate-200 pb-3 pt-2 px-1">
            <span v-for="genre in book.genres" class="text-sm text-slate-600 font-medium me-2 mb-2 border-2 rounded-md p-1">{{ this.capitalizedEachWord(genre.genre_name) }}</span>
          </div>
        </div>
      </section>
    </dialog>
  </teleport>
</template>

<script>
import appEnvironment from '@/environment.js'

export default {
  name: 'BookDialog',
  emits: ['close-event'],
  props: ['book'],
  computed: {
    appEnvironment() {
      return appEnvironment
    },
    allAuthors() {
      let result = ''
      if (this.book.authors.length == 1) {
        return this.capitalizedEachWord(this.book.authors[0].author_name)
      } else {
        for (let i = 0; i < this.book.authors.length; i++) {
          result += this.book.authors[i].author_name
          if (i !== this.book.authors.length - 1) {
            result += ', '
          }
        }
        return this.capitalizedEachWord(result)
      }
    },
    normalizedPublisherName() {
      return this.capitalizedEachWord(this.book.publisher.publisher_name)
    },
  },
  methods: {
    capitalizedEachWord(words) {
      const arr = words.split(" ")
      for (let i = 0; i < arr.length; i++) {
        arr[i] = arr[i][0].toUpperCase() + arr[i].substring(1)
      }
      return arr.join(" ")
    }
  }
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