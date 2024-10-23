<template>
<!--  <book-dialog v-if="showBookDialog" @close-event="handleCloseEvent" @write-review-event="handleWriteReviewEvent" :book="book"></book-dialog>-->
<!--  <review-dialog v-if="showWriteReviewDialog" @close-event="handleCloseEvent" :book="book"></review-dialog>-->
  <main-hero v-if="!store.isLoggedIn"></main-hero>
  <main class="bg-green-50 m-0 p-5 sm:p-10">
    <div v-if="isLoading" class="text-center text-4xl">Loading...</div>
    <div v-else>
      <transition-group tag="div" appear name="transition-books" class="flex flex-wrap items-start justify-center">
        <div v-for="book in books" :key="book.id">
          <div v-if="book.genre_ids.includes(currentFilter) || currentFilter === 0">
            <book-card :image-path="imagePath + book.slug + '.jpg'"
                       :id="book.id"
                       :title="book.title"
                       :description="book.description"
                       :genres="book.genres"
                       :author="allAuthors(book)"
                       :publication-year="book.publication_year"
                       @show-book-dialog-event="handleShowBookDialogEvent"
                       @show-book-detail-event="handleShowBookDetailEvent">
            </book-card>
          </div>
        </div>
      </transition-group>
    </div>
    <div v-if="store.isLoggedIn">
      <div class="hidden lg:inline mt-10 filters text-center mb-10">
        <div class="flex flex-wrap items-center justify-center">
          <span class="filter me-2 mb-2" :class="{active: currentFilter === 0}" @click="setFilter(0)">ALL</span>
          <span v-for="genre in genres" :key="genre.id" class="filter me-2 mb-2"
                :class="{active: currentFilter === genre.id}"
                @click="setFilter(genre.id)">{{ genre.genre_name.toLowerCase() }}</span>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
import MainHero from '@/components/MainHero.vue'
import { store } from '@/store.js'
import appEnvironment from '@/environment.js'
import BookCard from '@/components/BookCard.vue'
import BookDialog from '@/components/BookDialog.vue'
import router from '@/router/index.js'
import ReviewDialog from '@/components/ReviewDialog.vue'

export default {
  name: 'MainContent',
  components: {
    MainHero,
    BookCard,
    BookDialog,
    ReviewDialog
  },
  data() {
    return {
      store,
      imagePath: appEnvironment.imageURL(),
      book: {},
      books: [{}],
      genres: [{}],
      currentFilter: 0,
      isLoading: true,
      showBookDialog: false,
      showWriteReviewDialog: false
    }
  },
  methods: {
    // handleCloseEvent() {
    //   //console.log('handleCloseEvent')
    //   this.showBookDialog = false
    //   this.showWriteReviewDialog = false
    // },
    handleShowBookDialogEvent(id) {
      //console.log('handleShowBookDialogEvent')
      //console.log(id)
      this.showBookDialog = true
      this.book = this.books.find(book => book.id === id)
    },
    handleShowBookDetailEvent(id) {
      this.showBookDetail = true
      this.book = this.books.find(book => book.id === id)
      store.book = this.book
      console.log(this.book.title)
      router.push('/books/' + id)
    },
    // handleWriteReviewEvent() {
    //   console.log('handle write review event')
    //   this.showBookDialog = false
    //   this.showWriteReviewDialog = true
    // },
    setFilter: function(filter) {
      this.currentFilter = filter
      //console.log(this.currentFilter)
    },
    allAuthors(book) {
      let result = ''
      if (book.authors.length === 1) {
        return this.capitalizedEachWord(book.authors[0].author_name)
      } else {
        for (let i = 0; i < book.authors.length; i++) {
          result += book.authors[i].author_name
          if (i !== book.authors.length - 1) {
            result += ', '
          }
        }
        return this.capitalizedEachWord(result)
      }
    },
    capitalizedEachWord(words) {
      const arr = words.split(' ')
      for (let i = 0; i < arr.length; i++) {
        arr[i] = arr[i][0].toUpperCase() + arr[i].substring(1)
      }
      return arr.join(' ')
    },
    lowercasedEachWord(words) {
      const arr = words.split(' ')
      for (let i = 0; i < arr.length; i++) {
        arr[i] = arr[i][0].toLowerCase() + arr[i].substring(1)
      }
      return arr.join(' ')
    }
  },
  beforeMount() {
    // Fetch all genres
    fetch(appEnvironment.apiURL() + '/genres')
      .then((resp) => resp.json())
      .then((jsonResp) => {
        if (jsonResp.error) {
          notie.alert({
            type: 'error',
            text: jsonResp.message
          })
        } else {
          this.genres = jsonResp.data.genres
          //this.isLoading = false
          //console.log(this.genres)
        }
      })
      .catch((err) => {
        console.log(err)
      })

    // Fetch all books
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
          this.isLoading = false
          //console.log(this.books)
        }
      })
      .catch((err) => {
        console.log(err)
      })
  }
}
</script>

<style scoped>
.filters {
  height: 2.5em;
}

.filter {
  padding: 6px 6px;
  //margin: 8px 4px;
  cursor: pointer;
  border-radius: 6px;
  transition: all 0.35s;
  border: 1px solid lightgreen;
}

.filter.active {
  background: lightgreen;
}

.filter:hover {
  background: lightyellow;
}

.transition-books-move {
  transition: all 0.5s ease-in-out;
}

.transition-books-enter-active {
  transition: all 0.5s ease-out;
}

.transition-books-leave-active {
  transition: all 0.5s ease-in;
}

.transition-books-enter-from, .transition-books-leave-to {
  opacity: 0;
  transform: scale(0.8);
}

.transition-books-enter-to, .transition-books-leave-from {
  opacity: 1;
  transform: scale(1);
}
</style>