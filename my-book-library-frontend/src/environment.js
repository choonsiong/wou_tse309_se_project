const appEnvironment = {
  apiURL: function() {
    return import.meta.env.VITE_APP_API_URL
  },
}

export default appEnvironment