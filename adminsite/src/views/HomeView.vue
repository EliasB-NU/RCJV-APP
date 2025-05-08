<script setup lang="ts">
import HeaderComponent from '@/components/HeaderComponent.vue'
import LeaguesComponent from '@/components/LeaguesComponent.vue'
import UserComponent from '@/components/UserComponent.vue'


</script>

<template>
  <HeaderComponent />
  <div class="grid gap-10 md:grid-cols-1 lg:grid-cols-2">
    <LeaguesComponent />
    <UserComponent />
  </div>

</template>

<style scoped>

</style>
