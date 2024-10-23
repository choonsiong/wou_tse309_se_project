<template>
  <section class="bg-green-100">
    <h2 class="pt-10 text-4xl mb-6 font-bold text-center">Manage All Reviews</h2>
    <div class="max-w-4xl mx-auto p-6 space-y-6">
      <div v-for="(review, index) in reviews" :key="review.id">
        <div class="flex flex-col items-center justify-between w-full p-6 bg-green-500/30 rounded-lg md:flex-row">
          <div class="flex-row space-y-2">
            <p class="text-center text-veryDarkViolet md:text-left mr-2">#{{ index + 1 }} - <span class="font-semibold">{{ review.book_name
              }}</span></p>
            <p class="mb-1 font-light">{{ review.created_at.slice(0, 19).replace('T', ' ') }} by {{ review.user_name
              }}</p>
            <p>{{ review.review }}</p>
          </div>
          <div class="flex flex-col items-center justify-end flex-1 space-x-4 space-y-2 md:flex-row md:space-y-0">
            <button class="p-2 px-8 text-white bg-red rounded-lg hover:opacity-70 focus:outline-none"
                    @click="handleDeleteReview(review.id)">Delete
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import Security from '@/security.js'
import appEnvironment from '@/environment.js'
import notie from 'notie'

export default {
  name: 'AllReviews',
  emits: ['forceUpdateEvent'],
  data() {
    return {
      reviews: {}
    }
  },
  methods: {
    handleDeleteReview(reviewId) {
      notie.confirm({
        text: 'Are you sure you want to delete the review id: ' + reviewId,
        submitText: 'Delete',
        submitCallback: () => {

          let payload = {
            id: reviewId
          }

          fetch(appEnvironment.apiURL() + '/admin/reviews/delete', Security.requestOptions(payload))
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
                  text: 'Review deleted successfully'
                })
                this.$emit('forceUpdateEvent')
              }
            })
            .catch((err) => {
              console.log(err)
            })
        }
      })
    }
  },
  beforeMount() {
    Security.requireAdmin()

    fetch(appEnvironment.apiURL() + '/reviews', Security.requestOptions(''))
      .then((resp) => resp.json())
      .then((jsonResp) => {
        if (jsonResp.error) {
          notie.alert({
            type: 'error',
            text: jsonResp.message
          })
        } else {
          this.reviews = jsonResp.data.results
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