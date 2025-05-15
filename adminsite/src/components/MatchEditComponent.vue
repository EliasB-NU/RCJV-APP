<script setup lang="ts">

import axios from 'axios'
import Cookies from 'js-cookie'
import { type PropType, ref } from 'vue'

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

interface Institution {
  id: number
  name: string
}

const institutions = ref<Institution[]>([])

const matchToEdit = ref<Match>({} as Match)

const props = defineProps({
  match: {
    type: Object as PropType<Match>,
    required: true
  }
})

matchToEdit.value = props.match

const emit = defineEmits(['finished', 'close'])

const updateMatch = async () => {
  try {
    await axios
      .post(`/api/matches/update/${props.match.id}`, {
        ...matchToEdit.value,
      }, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
    .then(() => {
      emit('finished')
    })
  } catch (error) {
    console.error(error)
    emit('close')
  }
}

</script>

<template>
  <div class="fixed inset-0 bg-opacity-50 flex items-center justify-center">
    <div class="bg-white p-6 rounded-2xl shadow-lg w-full max-w-2xl max-h-[90vh] overflow-y-auto">
      <h3 class="text-lg font-semibold mb-4">Create Team</h3>

      <!-- Form to create client -->
      <div class="mb-4">
        <form @submit.prevent="updateMatch">
          <div class="grid grid-cols-2 gap-4">
            <!-- Left Column: Basic Information -->
            <div class="w-1/2 pr-4">
              <div class="mb-4">
                <label for="name" class="block">Name:</label>
                <input v-model="matchToEdit.name" id="name" type="text"
                       class="input-field border-2 border-b-black rounded-sm" required />
              </div>
              <div class="mb-4">
                <label for="institution" class="block">Institution:</label>
                <select v-model="matchToEdit.institutionID" id="institution" class="input-field border-2 border-b-black rounded-sm">
                  <option v-for="inst in institutions" :key="inst.id" :value="inst.id">{{ inst.name }}</option>
                </select>
              </div>
            </div>
          </div>

          <div class="p-3 flex space-x-2 mt-4 justify-end">
            <button type="submit" class="bg-gray-800 text-white px-3 py-1 rounded hover:bg-gray-700">
              Create Team
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
