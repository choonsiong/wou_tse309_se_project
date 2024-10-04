<template>
  <a href="javascript:void(0)">
    <div class="relative flex flex-col my-6 bg-green-100 shadow-sm border border-green-200 h-[600px] w-96 m-2">
      <div class="relative text-center m-2.5 overflow-hidden text-white">
        <img :src="imagePath" alt="card-image" class="w-full h-full object-cover object-center"/>
      </div>
      <div class="p-4">
        <div class="relative flex flex-row flex-wrap gap-2">
          <div v-for="genre in genres" :key="genre.id" class="mb-4 rounded-full bg-green-400 py-0.5 px-2.5 border border-transparent text-xs text-white transition-all shadow-sm w-20 text-center">
            {{ genre.genre_name }}
          </div>
        </div>
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
  </a>
</template>
<script>
export default {
  name: 'BookCard',
  props: ['title', 'description', 'imagePath', 'author', 'publicationYear', 'genres'],
  computed: {
    shortDescription() {
      return this.description.length > 200 ? this.description.slice(0, 200-1) + '...' : this.description;
    }
  }
}
</script>