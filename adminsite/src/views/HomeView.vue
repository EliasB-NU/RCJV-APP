<script setup lang="ts">
import HeaderComponent from '@/components/HeaderComponent.vue'
import LeaguesComponent from '@/components/LeaguesComponent.vue'
import UserComponent from '@/components/UserComponent.vue'
import ConfigComponent from '@/components/ConfigComponent.vue'
import PopUp from '@/components/PopUp.vue'
import { ref } from 'vue'

const popUp = ref<InstanceType<typeof PopUp> | null>(null)

const showPopUp = (s: any) => {
  popUp.value?.show(s)
}
</script>

<template>
  <HeaderComponent />
  <div class="grid gap-10 md:grid-cols-1 lg:grid-cols-3">
    <LeaguesComponent />
    <ConfigComponent @popUp="showPopUp" />
    <UserComponent />
    <PopUp ref="popUp" />
  </div>

</template>

<style scoped>

</style>
