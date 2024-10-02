<template>
  <h1>All Users</h1>
</template>

<script>
import Security from '@/security.js'
import appEnvironment from '@/environment.js'
import notie from 'notie'

export default {
  name: 'AllUsers',
  data() {
    return {
      users: []
    }
  },
  beforeMount() {
    Security.requireAdmin()

    fetch(appEnvironment.apiURL() + '/admin/users/all', Security.requestOptions(""))
      .then((resp) => resp.json())
      .then((jsonResp) => {
        if (jsonResp.error) {
          //console.log('error: ' + jsonResp.message)
          notie.alert({
            type: 'error',
            text: jsonResp.message,
          })
        } else {
          //console.log(jsonResp)
          this.users = jsonResp.data.users
        }
      })
      .catch((err) => {
        notie.alert({
          type: 'error',
          text: err
        })
      })
  },
}
</script>

<style scoped>

</style>