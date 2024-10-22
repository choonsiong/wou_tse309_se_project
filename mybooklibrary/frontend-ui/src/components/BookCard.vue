<template>
  <div>
    <div class="relative flex flex-col mb-4 sm:my-4 bg-green-100 shadow-sm border border-green-200 h-[600px] w-[100%] sm:w-96 sm:m-2">
      <div class="flex justify-end">
        <div class="hidden md:inline m-4 cursor-pointer justify-end" @click="showBookDialog">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor"
               class="size-6">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M13.5 6H5.25A2.25 2.25 0 0 0 3 8.25v10.5A2.25 2.25 0 0 0 5.25 21h10.5A2.25 2.25 0 0 0 18 18.75V10.5m-10.5 6L21 3m0 0h-5.25M21 3v5.25" />
          </svg>
        </div>
        <div class="md:hidden m-4 cursor-pointer justify-end" @click="showBookDetail">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor"
               class="size-6">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M13.5 6H5.25A2.25 2.25 0 0 0 3 8.25v10.5A2.25 2.25 0 0 0 5.25 21h10.5A2.25 2.25 0 0 0 18 18.75V10.5m-10.5 6L21 3m0 0h-5.25M21 3v5.25" />
          </svg>
        </div>
      </div>
      <div class="relative text-center m-2.5 overflow-hidden text-white">
        <img :src="imagePath" alt="card-image" class="w-full h-[280px] object-scale-down" />
      </div>
      <div class="p-4">
        <!--        <div class="relative flex flex-row flex-wrap gap-2">-->
        <!--          <div v-for="genre in genres" :key="genre.id" class="mb-4 rounded-full bg-green-400 py-0.5 px-2.5 border border-transparent text-xs text-white transition-all shadow-sm w-20 text-center">-->
        <!--            {{ genre.genre_name }}-->
        <!--          </div>-->
        <!--        </div>-->
        <h6 class="mb-2 text-slate-800 text-xl font-semibold">
          {{ title }}
        </h6>
        <p class="text-slate-600 leading-normal font-light">
          {{ shortDescription }}
        </p>
      </div>

      <div class="flex items-center justify-between p-4">
        <div class="flex items-center">
          <!--          <img-->
          <!--            alt="Tania Andrew"-->
          <!--            src="https://images.unsplash.com/photo-1633332755192-727a05c4013d?ixlib=rb-1.2.1&amp;ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&amp;auto=format&amp;fit=crop&amp;w=1480&amp;q=80"-->
          <!--            class="relative inline-block h-8 w-8 rounded-full"-->
          <!--          />-->
          <div class="flex flex-col ml-0 text-sm">
            <span class="text-slate-800 font-semibold">{{ author }}</span>
            <span class="text-slate-600">{{ publicationYear }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'BookCard',
  props: ['id', 'title', 'description', 'imagePath', 'author', 'publicationYear', 'genres'],
  emits: ['show-book-dialog-event'],
  computed: {
    shortDescription() {
      return this.description.length > 200 ? this.description.slice(0, 200 - 1) + '...' : this.description
    }
  },
  methods: {
    showBookDialog() {
      //console.log(this.id)
      this.$emit('show-book-dialog-event', this.id)
    },
    showBookDetail() {
      console.log(this.id)
      this.$emit('show-book-detail-event', this.id)
    },
    capitalizedEachWord(words) {
      const arr = words.split(' ')
      for (let i = 0; i < arr.length; i++) {
        arr[i] = arr[i][0].toUpperCase() + arr[i].substring(1)
      }
      return arr.join(' ')
    }
  }
}
</script>