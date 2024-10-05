const appEnvironment = {
  apiURL: function() {
    return import.meta.env.VITE_APP_API_URL
  },
  imageURL: function() {
    return import.meta.env.VITE_APP_IMAGE_URL
  }
}

export default appEnvironment