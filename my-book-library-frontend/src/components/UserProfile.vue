<template>
  <h2>My Profile</h2>
</template>

<script>
import Security from '@/security.js'

export default {
  name: 'UserProfile',
  beforeMount() {
    Security.requireToken()
  }
}
</script>