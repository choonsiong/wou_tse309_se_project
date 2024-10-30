<template>
  <section class="bg-white">
    <div class="max-w-4xl mx-auto p-6 space-y-6">
      <h1 class="mt-3 text-4xl font-bold">Edit Book</h1>
      <hr>
      <div>
        <img :src="appEnvironment.imageURL() + '/' + book.slug + '.jpg'" />
      </div>
      <form @submit.prevent="editBook">
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Book Title</label>
          <input v-model.trim="book.title" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 bg-gray-50 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="Book Title" disabled />
        </div>
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Publication Year</label>
          <input v-model.trim="book.publication_year" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="Publication Year" />
        </div>
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Publisher</label>
          <input v-model.trim="book.publisher.publisher_name" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="Publisher" />
        </div>
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">ISBN</label>
          <input v-model.trim="book.isbn" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="ISBN" />
        </div>
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Book Authors</label>
          <input v-model.trim="authors" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="Jane Doe, John Doe" />
        </div>
        <div class="w-full max-w-md min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Book Categories</label>
          <input v-model.trim="genres" type="text"
                 class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                 placeholder="Web Development, Programming" />
        </div>
        <div class="relative max-w-md  w-full min-w-[200px] mb-5">
          <label class="block mb-2 text-sm text-slate-600">Description</label>
          <textarea rows="20" v-model.trim="book.description"
                    class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                    placeholder="Enter book description"></textarea>
        </div>
        <div class="mb-5">
          <label for="formImageFile" class="block mb-2 text-sm text-slate-600">Book Cover Image</label>
          <input ref="bookCoverRef" type="file" id="formImageFile" accept="image/jpeg" @change="uploadBookCoverImage"
                 class="bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow" />
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
import appEnvironment from '@/environment.js'
import Security from '@/security.js'
import router from '@/router/index.js'

export default {
  name: 'EditBook',
  data() {
    return {
      book: {},
      authors: '',
      genres: '',
      bookCover: null
    }
  },
  computed: {
    appEnvironment() {
      return appEnvironment
    },
  },
  methods: {
    initializeAuthors() {
      let result = ''
      if (this.book.authors.length === 1) {
        result += this.book.authors[0].author_name
      } else {
        for (let i = 0; i < this.book.authors.length; i++) {
          result += this.book.authors[i].author_name
          if (i !== this.book.authors.length - 1) {
            result += ', '
          }
        }
      }
      this.authors = result
      return result
    },
    initializeGenres() {
      let result = ''
      if (this.book.genres.length === 1) {
        result += this.book.genres[0].genre_name
      } else {
        for (let i = 0; i < this.book.genres.length; i++) {
          result += this.book.genres[i].genre_name
          if (i !== this.book.genres.length - 1) {
            result += ', '
          }
        }
      }
      this.genres = result
      return result
    },
    editBook() {
      const payload = {
        id: this.book.id,
        title: this.book.title,
        publisher_name: this.book.publisher.publisher_name,
        publication_year: parseInt(this.book.publication_year, 10),
        authors: this.authors,
        genres: this.genres,
        description: this.book.description,
        isbn: this.book.isbn,
        slug: this.book.slug,
        book_cover: this.bookCover
      }

      fetch(appEnvironment.apiURL() + '/admin/books/edit', Security.requestOptions(payload))
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
              text: 'Book updated successfully'
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
      router.push('/admin/manage/books')
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
  },
  beforeMount() {
    fetch(appEnvironment.apiURL() + '/books/' + this.$route.params.id)
      .then((resp) => resp.json())
      .then((jsonResp) => {
        if (jsonResp.error) {
          console.log(jsonResp.message)
        } else {
          this.book = jsonResp.data.book
          //console.log(this.book)
          this.initializeAuthors()
          this.initializeGenres()
        }
      })
      .catch((err) => {
        notie.alert({
          type: 'error',
          text: err
        })
      })
  }
}
</script>