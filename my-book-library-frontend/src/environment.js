const appEnvironment = () => {
  const apiURL = import.meta.env.VITE_APP_API_URL
  return {
    apiURL: apiURL,
  }
}

export default appEnvironment