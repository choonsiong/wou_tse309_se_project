<template>
  <h1>All Users</h1>
</template>

<script>
import Security from '@/security.js'

export default {
  name: 'AllUsers',
  beforeMount() {
    Security.requireAdmin()
  },
}
</script>

<style scoped>

</style>