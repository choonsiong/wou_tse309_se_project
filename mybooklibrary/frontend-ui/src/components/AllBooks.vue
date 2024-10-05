<template>
  <section class="bg-green-100">
    <h2 class="pt-10 text-4xl mb-6 font-bold text-center">Manage All Books</h2>
    <div class="max-w-4xl mx-auto p-6 space-y-6">
      <div v-for="(book, index) in books" :key="book.id">
        <div class="px-5 py-5 flex flex-col h-full bg-white items-center">
          <div class="flex flex-row w-full items-start gap-5 mb-10">
            <div class="">
              <img class="w-56 p-2 bg-white object-fill border-2 border-gray-100" :src="appEnvironment.imageURL() + '/' + book.slug + '.jpg'"
                   alt="book image">
            </div>
            <div class="flex-1">
              <p class="mb-2 font-bold text-4xl">{{ book.title }}</p>
              <p class="mb-2"><span class="font-bold text-xl">{{ allAuthors(book) }}</span></p>
              <p class="mb-5">{{ book.description }}</p>
              <p class="font-light">Publisher: <span class="text-gray-500 font-light">{{ book.publisher.publisher_name }}</span></p>
              <p class="font-light mb-5">Publication Year: <span class="text-gray-500 font-light">{{ book.publication_year }}</span></p>
              <div class="flex flex-wrap">
                <span class="p-1 me-2 mb-2 border-slate-200 border-2 me-2 text-gray-500 font-light" v-for="genre in book.genres">{{ genre.genre_name }}</span>
              </div>
            </div>
          </div>
          <div>
            <router-link :to="`/admin/manage/books/edit/${book.id}`" class="p-3 px-8 text-white bg-cyan rounded-lg hover:opacity-70 focus:outline-none me-4">Edit</router-link>
            <button class="p-2 px-8 text-white bg-red rounded-lg hover:opacity-70 focus:outline-none" @click="handleDeleteBook(book.id)">Delete</button>
          </div>
        </div>
      </div>
      <div class="text-center pt-10">
        <router-link to="/admin/manage/books/new" class="p-4 px-8 text-white bg-blue-600 rounded-lg hover:opacity-70 focus:outline-none">Add New Book</router-link>
      </div>
    </div>
    <div class="p-10">&nbsp;</div>
  </section>
</template>

<script>
import Security from '@/security.js'
import appEnvironment from '@/environment.js'

export default {
  name: 'AllBooks',
  computed: {
    appEnvironment() {
      return appEnvironment
    }
  },
  data() {
    return {
      books: []
    }
  },
  methods: {
    allAuthors(book) {
      let result = ''
      if (book.authors.length == 1) {
        return book.authors[0].author_name
      } else {
        for (let i = 0; i < book.authors.length; i++) {
          result += book.authors[i].author_name
          if (i !== book.authors.length - 1) {
            result += ', '
          }
        }
        return result
      }
    },
    handleDeleteBook(bookId) {
      notie.confirm({
        text: 'Are you sure you want to delete the book id: ' + bookId,
        submitText: 'Delete',
        submitCallback: () => {
          let payload = {
            id: bookId,
          }
          fetch(appEnvironment.apiURL() + '/admin/books/delete', Security.requestOptions(payload))
            .then((resp) => resp.json())
            .then((jsonResp) => {
              if (jsonResp.error) {
                console.log('error: ' + jsonResp.message)
                notie.alert({
                  type: 'error',
                  text: jsonResp.message
                })
              } else {
                notie.alert({
                  type: 'success',
                  text: 'Book deleted successfully'
                })
                this.$emit('forceUpdateEvent')
              }
            })
            .catch((err) => {
              console.log(err)
            })
        }
      })
    },
  },
  beforeMount() {
    Security.requireAdmin()

    fetch(appEnvironment.apiURL() + '/books')
      .then((resp) => resp.json())
      .then((jsonResp) => {
        if (jsonResp.error) {
          notie.alert({
            type: 'error',
            text: jsonResp.message
          })
        } else {
          this.books = jsonResp.data.books
          console.log(this.books)
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

<style scoped>

</style>