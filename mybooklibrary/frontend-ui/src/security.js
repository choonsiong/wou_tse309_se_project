import router from './router'
import { store } from '@/store.js'
import appEnvironment from '@/environment.js'

let Security = {
  // Check whether a valid token is exists
  requireToken: function() {
    //console.log('requireToken')
    if (store.token === '') {
      router.replace('/')
      return false
    }
  },

  // Check whether user is an administrator
  requireAdmin: function() {
    //console.log('requireAdmin')
    //console.log('token: ' + store.token)
    //console.log('isAdmin: ' + store.isAdmin)

    if (store.token === '') {
      router.replace('/')
      return false
    }

    if (!store.isAdmin) {
      router.replace('/')
      return false
    }

    return true
  },

  // validate user token
  validateToken: function() {
    if (store.token !== '') {
      const payload = {
        token: store.token
      }

      const headers = new Headers()
      headers.append('Content-Type', 'application/json')

      let requestOptions = {
        method: 'POST',
        body: JSON.stringify(payload),
        headers: headers
      }

      fetch(appEnvironment.apiURL() + '/token/validate', requestOptions)
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            console.log(response.message)
          } else {
            if (!response.data) {
              store.token = ''
              store.user = {}
              document.cookie = '_mybooklibrary_data=; Path=/; SameSite=strict; Secure; Expires=Thu, 01 Jan 1970 00:00:01 GMT;'
            }
          }
        })
    }
  },

  // Request options for fetch request
  requestOptions: function(payload) {
    const headers = new Headers()

    headers.append('Content-Type', 'application/json')
    headers.append('Authorization', 'Bearer ' + store.token)

    //console.log(payload)
    //console.log(JSON.stringify(payload))

    return {
      method: 'post',
      body: JSON.stringify(payload),
      headers: headers,
    }
  }
}

export default Security