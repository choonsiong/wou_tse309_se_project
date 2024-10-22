<template>
  <div class="p-4 bg-green-100">
    <p class="font-bold text-center text-3xl">{{ book.title }}</p>
    <div class="m-2.5 overflow-hidden rounded-md h-80 flex justify-center items-center">
      <img class="w-full h-full object-scale-down" :src="appEnvironment.imageURL() + '/' + book.slug + '.jpg'" alt="profile-picture" />
    </div>
    <p class="text-center font-light text-xl">{{ allAuthors }}</p>
    <p class="mt-5 text-center font-bold text-2xl">{{ normalizedPublisherName }} ({{ book.publication_year }})</p>
    <p class="mt-10 font-extralight">{{ book.description }}</p>
    <div class="mt-10">
      <div class="flex flex-wrap mx-3 pb-3 pt-2 px-1">
        <span v-for="genre in book.genres" class="text-sm text-slate-600 font-medium me-2 mb-2 border-2 rounded-md p-1">{{ genre.genre_name.toLowerCase() }}</span>
      </div>
    </div>
  </div>
</template>
<script>
import { store } from '@/store.js'
import appEnvironment from '@/environment.js'

export default {
  name: 'BookDetail',
  data() {
    return {
      book: store.book,
    }
  },
  computed: {
    appEnvironment() {
      return appEnvironment
    },
    allAuthors() {
      let result = ''
      if (store.book.authors.length == 1) {
        return this.capitalizedEachWord(store.book.authors[0].author_name)
      } else {
        for (let i = 0; i < store.book.authors.length; i++) {
          result += store.book.authors[i].author_name
          if (i !== store.book.authors.length - 1) {
            result += ', '
          }
        }
        return this.capitalizedEachWord(result)
      }
    },
    normalizedPublisherName() {
      return this.capitalizedEachWord(store.book.publisher.publisher_name)
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