<script setup lang="ts">

import { onMounted, ref } from 'vue'
import axios from 'axios'
import Cookies from 'js-cookie'

const props = defineProps({
  team: {
    required: true,
    type: Object,
  }
})

const emit = defineEmits(['close', 'finished'])

interface Team {
  name: string,
  league: string,
  institutionID: number
}

const data = ref<Team>({
  name: props.team.name,
  league: props.team.league,
  institutionID: props.team.institutionID,
})

interface Institution {
  id: number,
  name: string,
}

const institutions = ref<Institution[]>([])

async function fetchInstitutions() {
  try {
    await axios
      .get('/api/institutions', {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json'
        }
      })
      .then((res) => {
        institutions.value = res.data.data
      })
  } catch (error) {
    console.log(error)
  }
}

const updateTeam = async () => {
  try {
    await axios
      .post(`/api/teams/update/${props.team.id}`, {
        name: data.value.name,
        league: data.value.league,
        institutionID: data.value.institutionID
      }, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': `Bearer ${Cookies.get('token')}`
        }
      })
      .then(() => {
        emit('finished')
      })
  } catch (error) {
    console.log(error)
    emit('close')
  }
}

onMounted(fetchInstitutions())

</script>

<template>
  <div class="fixed inset-0 bg-opacity-50 flex items-center justify-center">
    <div class="bg-white p-6 rounded-2xl shadow-lg w-full max-w-2xl max-h-[90vh] overflow-y-auto">
      <h3 class="text-lg font-semibold mb-4">Edit Team</h3>

      <!-- Form to create client -->
      <div class="mb-4">
        <form @submit.prevent="updateTeam">
          <div class="grid grid-cols-2 gap-4">
            <!-- Left Column: Basic Information -->
            <div class="w-1/2 pr-4">
              <div class="mb-4">
                <label for="name" class="block">Name:</label>
                <input v-model="data.name" id="name" type="text"
                       class="input-field border-2 border-b-black rounded-sm" required />
              </div>
              <div class="mb-4">
                <label for="league" class="block">League:</label>
                <select v-model="data.league" id="league"
                        class="input-field border-2 border-b-black rounded-sm">
                  <option value="soccerEntry">Soccer Entry</option>
                  <option value="soccerLightWeightEntry">Soccer LightWeight Entry</option>
                  <option value="soccerLightWeight">Soccer LightWeight int.</option>
                  <option value="soccerOpen">Soccer Open int.</option>
                  <option value="rescueLineEntry">Rescue Line Entry</option>
                  <option value="rescueLine">Rescue Line int.</option>
                  <option value="rescueMazeEntry">Rescue Maze Entry</option>
                  <option value="rescueMaze">Rescue Maze int.</option>
                  <option value="onStageEntry">Onstage Entry</option>
                  <option value="onStage">Onstage int.</option>
                </select>
              </div>
              <div class="mb-4">
                <label for="institution" class="block">Institution:</label>
                <select v-model="data.institutionID" id="institution" class="input-field border-2 border-b-black rounded-sm">
                  <option v-for="inst in institutions" :key="inst.id" :value="inst.id">{{ inst.name }}</option>
                </select>
              </div>
            </div>
          </div>

          <div class="p-3 flex space-x-2 mt-4 justify-end">
            <button type="submit" class="bg-gray-800 text-white px-3 py-1 rounded hover:bg-gray-700">
              Update Team
            </button>
            <button type="button" @click="emit('close');" class="bg-green-600 text-white px-3 py-1 rounded hover:bg-green-700">
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
