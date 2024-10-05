<template>
  <section class="bg-white">
    <div class="max-w-4xl mx-auto p-6 space-y-6">
      <h1 class="mt-3 text-4xl font-bold">Add New Book</h1>
      <hr>
      <form @submit.prevent="addNewBook">
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Book Title</label>
          <input v-model="title" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="Book Title" />
        </div>
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Publication Year</label>
          <input v-model="publicationYear" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="Publication Year" />
        </div>
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Publisher</label>
          <input v-model="publisherName" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="Publisher" />
        </div>
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Book Authors</label>
          <input v-model="authors" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="Jane Doe, John Doe" />
        </div>
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Book Categories</label>
          <input v-model="genres" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="Web Development, Programming" />
        </div>
        <div class="relative max-w-md  w-full min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Description</label>
          <textarea v-model="description"
                    class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                    placeholder="Enter book description"></textarea>
        </div>
        <div class="mb-5">
          <label for="formImageFile" class="block mb-2 text-sm text-slate-600">Book Cover Image</label>
          <input ref="bookCoverRef" type="file" id="formImageFile" accept="image/jpeg" @change="uploadBookCoverImage"
                 class="bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 required />
        </div>
        <div class="text-center pt-10">
          <button class="p-3 px-8 text-white bg-blue-600 rounded-lg hover:opacity-70 focus:outline-none me-4"
                  @click="handleCancel">Cancel
          </button>
          <button class="p-3 px-8 text-white bg-rose-500 rounded-lg hover:opacity-70 focus:outline-none">Submit</button>
        </div>
      </form>
    </div>
  </section>
</template>

<script>
import router from '@/router/index.js'
import appEnvironment from '@/environment.js'
import Security from '@/security.js'
import { store } from '@/store.js'

export default {
  name: 'AddBook',
  data() {
    return {
      title: '',
      description: '',
      publisherName: '',
      publicationYear: '',
      authors: '',
      genres: '',
      bookCover: null
    }
  },
  methods: {
    addNewBook() {
      const payload = {
        title: this.title,
        description: this.description,
        publisher_name: this.publisherName,
        publication_year: parseInt(this.publicationYear, 10),
        authors: this.authors,
        genres: this.genres,
        book_cover: this.bookCover
      }

      fetch(appEnvironment.apiURL() + '/admin/books/new', Security.requestOptions(payload))
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
              text: 'Book added successfully'
            })
            router.push('/admin/manage/books')
          }
        })
        .catch((err) => {
          notie.alert({
            type: 'error',
            text: err
          })
        })
    },
    handleCancel() {
      if (store.isAdmin) {
        router.push('/admin/manage/books')
      } else {
        router.push('/manage/books')
      }
    },
    uploadBookCoverImage() {
      // Get a reference to the input using ref
      const file = this.$refs.bookCoverRef.files[0]

      // Encode the file using the FileReader API
      const reader = new FileReader()
      reader.onloadend = () => {
        const base64String = reader.result.replace('data:', '').replace(/^.+,/, '')
        this.bookCover = base64String
      }
      reader.readAsDataURL(file)
    }
  }
}
</script>