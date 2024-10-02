<template>
  <h2>Manage Books</h2>
</template>

<script>
import Security from '@/security.js'

export default {
  name: 'UserBooks',
  beforeMount() {
    Security.requireToken()
  },
}
</script>

<style scoped>

</style>