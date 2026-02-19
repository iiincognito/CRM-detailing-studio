<script lang="ts">
import Authentification from '~/components/Auth/Authentification.vue'

export default defineComponent({
  name: "login-page",
  components: {
    Authentification
  }
})
</script>

<template>
  <div class="login-page">
    <Authentification />
  </div>
</template>

<style scoped>

</style>