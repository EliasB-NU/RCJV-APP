<script setup lang="ts">

import { onMounted, ref } from 'vue'
import axios from 'axios'
import Cookies from 'js-cookie'

const emit = defineEmits(['popUp'])

interface Config {
  appEnabled: boolean
  eventName: string
  soccerUrl: string
  soccerAbbreviation: string
}

const config = ref<Config>({} as Config)

async function getConfig() {
  try {
    await axios
      .get('/api/v1/config', {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token'),
        }
      })
    .then(res => {
      config.value = res.data
    })
  } catch (error) {
    emit('popUp', 'Error fetching config')
    console.log(error)
  }
}

const updateConfig = async () => {
  try {
    await axios
      .post('/api/v1/config/update', {
        ...config.value
      }, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': `Bearer ${Cookies.get('token')}`,
        }
      })
    .then(() => {
      emit('popUp', 'Successfully updated config')
      getConfig()
    })
  } catch (error) {
    emit('popUp', 'Error updating config')
    console.log(error)
  }
}

onMounted(getConfig)

</script>

<template>
  <div class="inset-0 bg-opacity-50 flex items-center justify-center">
    <div class="bg-white p-6 rounded-2xl shadow-lg overflow-y-auto">
      <h3 class="text-lg font-semibold mb-4">Config:</h3>


      <div class="mb-4">
        <form @submit.prevent="updateConfig">
          <div class="space-y-2">
            <div class="flex justify-between items-center">
              <label for="appEnabled">App Enabled:</label>
              <input v-model="config.appEnabled" type="checkbox" id="appEnabled"/>
            </div>

            <div class="flex justify-between items-center">
              <label for="eventName">Event name:</label>
              <input v-model="config.eventName" type="text" id="eventName" class="input-field border-2 border-b-black rounded-sm"  required/>
            </div>

            <div class="flex justify-between items-center">
              <label for="soccerUrl">Soccer URL:</label>
              <input v-model="config.soccerUrl" type="text" id="soccerUrl" class="input-field border-2 border-b-black rounded-sm"  required/>
            </div>

            <div class="flex justify-between items-center">
              <label for="soccerAbbreviation">Soccer Abbreviation:</label>
              <input v-model="config.soccerAbbreviation" type="text" id="soccerAbbreviation" class="input-field border-2 border-b-black rounded-sm"  required/>
            </div>
          </div>

          <div class="p-3 flex space-x-2 mt-4 justify-end">
            <button type="submit" class="bg-gray-800 text-white px-3 py-1 rounded hover:bg-gray-700">Update</button>
            <button type="button" @click="getConfig" class="bg-green-600 text-white px-3 py-1 rounded hover:bg-green-700">Reset</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
