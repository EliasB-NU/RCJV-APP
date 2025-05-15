<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import axios from 'axios'
import PopUp from '@/components/PopUp.vue'
import HeaderComponent from '@/components/HeaderComponent.vue'
import Cookies from 'js-cookie'

interface Match {
  id: number

  league: string
  name: string
  startTime: string
  duration: string
  field: string

  institutionID: number
  institutionName: string
  teamID: number
  teamName: string
}

const matches = ref<Match[]>([])

const league = ref<string>('soccerEntry')

async function fetchMatches() {
  try {
    await axios
      .get(`/api/matches/league/${league.value}`, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json'
        }
      })
      .then(res => {
        matches.value = res.data.data
      })
  } catch (error) {
    popUp.value?.show("Error fetching matches")
    console.error(error)
  }
}

const deleteMatch = async (id: number) => {
  try {
    await axios
    .delete(`/api/matches/delete/${id}`, {
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
        'Authorization': `Bearer ${Cookies.get('token')}`
      }
    })
    .then(() => {
      fetchMatches()
      popUp.value?.show("Successfully deleted match")
    })
  } catch (error) {
    console.error(error)
    popUp.value?.show("Error fetching matches")
  }
}

const popUp = ref<InstanceType<typeof PopUp> | null>(null)

onMounted(async () => {
  await fetchMatches()
})

watch(league, async () => {
  await fetchMatches()
}, {
  immediate: true,
  deep: true
})
</script>

<template>
  <HeaderComponent />
  <div class="inset-0 grid grid-cols-1 items-center justify-center">
    <div>
      <div class="bg-opacity-50 bg-white p-6 rounded-2xl shadow-lg overflow-y-auto">
        <h3 class="text-lg font-semibold mb-4">Matches :</h3>
        <select v-model="league" class="border p-1 rounded flex-1">
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

      <div class="bg-white shadow rounded-lg p-6 overflow-y-auto">
        <table class="w-full border-collapse">
          <thead>
          <tr class="bg-gray-200">
            <th class="p-3 text-left">Name</th>
            <th class="p-3 text-left">Start Time</th>
            <th class="p-3 text-left">Duration</th>
            <th class="p-3 text-left">Field</th>
            <th class="p-3 text-left">Team</th>
            <th class="p-3 text-left">Institution</th>
            <th class="p-3">Actions</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="match in matches" :key="match.id" class="border-t">
            <td class="p-3">{{ match.name }}</td>
            <td class="p-3">{{ match.startTime }}</td>
            <td class="p-3">{{ match.duration }}</td>
            <td class="p-3">{{ match.field }}</td>
            <td class="p-3">{{ match.teamName }}</td>
            <td class="p-3">{{ match.institutionName }}</td>
            <td class="p-3 flex space-x-2">
              <button
                @click="deleteMatch(match.id)"
                class="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600"
              >
                Delete
              </button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>

    <PopUp ref="popUp" />
  </div>
</template>

<style scoped>

</style>
