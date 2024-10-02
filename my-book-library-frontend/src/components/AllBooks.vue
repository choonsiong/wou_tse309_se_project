<template>
  <h1>All Books</h1>
</template>

<script>
import Security from '@/security.js'

export default {
  name: 'AllBooks',
  beforeMount() {
    Security.requireAdmin()
  },
}
</script>

<style scoped>

</style>