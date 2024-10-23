<template>
  <section class="bg-white">
    <div class="max-w-4xl mx-auto p-6 space-y-6">
      <h1 class="mt-3 text-4xl font-bold">Write a review for {{ store.book.title }}</h1>
      <div class="m-2.5 overflow-hidden rounded-md h-80 flex justify-center items-center">
        <img class="w-full h-full object-scale-down" :src="appEnvironment.imageURL() + '/' + store.book.slug + '.jpg'"
             alt="profile-picture" />
      </div>
      <form @submit.prevent="addNewReview">
        <div class="relative max-w-4xl w-full min-w-[200px] mb-5">
          <!--          <label class="block mb-2 text-sm text-slate-600">Review</label>-->
          <textarea v-model="review"
                    class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                    placeholder="Enter book review" rows="10"></textarea>
        </div>
        <div class="text-center pt-10">
          <button class="p-3 px-8 text-white bg-blue-600 rounded-lg hover:opacity-70 focus:outline-none me-4"
                  @click="handleCancel">Cancel
          </button>
          <button class="p-3 px-8 text-white bg-rose-500 rounded-lg hover:opacity-70 focus:outline-none">Submit Review</button>
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
  name: 'AddReview',
  computed: {
    appEnvironment() {
      return appEnvironment
    },
    store() {
      return store
    }
  },
  data() {
    return {
      book: store.book,
      review: '',
      rating: 0
    }
  },
  methods: {
    addNewReview() {
      const payload = {
        user_id: store.user.id,
        book_id: this.book.id,
        review: this.review,
        rating: this.rating,
      }

      fetch(appEnvironment.apiURL() + '/admin/reviews/new', Security.requestOptions(payload))
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
              text: 'Review added successfully'
            })
            if (store.isAdmin) {
              //router.push('/admin/manage/reviews')
              router.push('/')
            } else {
              //router.push('/manage/reviews')
              router.push('/')
            }
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
        router.push('/books/' + store.book.id)
      } else {
        router.push('/books/' + store.book.id)
      }
    }
  }
}
</script>