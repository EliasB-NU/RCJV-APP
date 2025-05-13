<script setup lang="ts">

import { onMounted, ref } from 'vue'
import HeaderComponent from '@/components/HeaderComponent.vue'
import PopUp from '@/components/PopUp.vue'
import axios from 'axios'
import Cookies from 'js-cookie'

interface Field {
  id: number
  name: string
  league: string
}

const fields = ref<Field[]>([])

async function fetchFields() {
  try {
    await axios
      .get('/api/fields', {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
      .then(res => {
        fields.value = res.data
      })
  } catch (error) {
    popUp.value?.show('Error fetching fields.')
    console.log(error)
  }
}

const createField = async () => {
  try {
    await axios
      .post('/api/fields/create', {
        name: 'newField',
        league: 'soccerEntry',
      }, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
    .then(() => {
      fetchFields()
    })
  } catch (error) {
    popUp.value?.show('Error creating field.')
    console.log(error)
  }
}

const updateField = async (fieldID: number, index: number) => {
  try {
    await axios
      .post(`/api/fields/update/${fieldID}`, {
        ...fields.value[index],
      }, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
      .then(() => {
        popUp.value?.show('Successfully updated field.')
        fetchFields()
      })
  } catch (error) {
    popUp.value?.show('Error updating field.')
    console.log(error)
    await fetchFields()
  }
}

const deleteField = async (id: number) => {
  try {
    await axios
      .delete(`/api/fields/delete/${id}`, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
    .then(() => {
      popUp.value?.show('Successfully deleted field.')
      fetchFields()
    })
  } catch (error) {
    popUp.value?.show('Error deleting field.')
    console.log(error)
  }
}

const popUp = ref<InstanceType<typeof PopUp> | null>(null)

onMounted(async () => {
  await fetchFields()
})
</script>

<template>
  <HeaderComponent />
  <div class="grid gap-10 md:grid-cols-1 lg:grid-cols-2 space-y-2 p-10">
    <div class="bg-white p-6 rounded-2xl shadow-lg overflow-y-auto">
      <div class="bg-white p-6 rounded-2xl shadow-lg overflow-y-auto">
        <div class="mb-4">
          <h2 class="text-2xl font-semibold mb-4">Fields</h2>
          <button @click="createField"
                  class="bg-gray-800 hover:bg-gray-700 text-white px-4 py-2 rounded">
            + Create Field
          </button>
        </div>

        <div v-for="(field, index) in fields" :key="field.id"
             class="flex items-center space-x-2 mb-2">
          <input
            v-model="field.name"
            class="input-field border p-1 rounded"
            placeholder="Name"
          />
          <select v-model="field.league" id="league"
                  class="border p-1 rounded-xl">
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
          <button @click="updateField(field.id, index)"
                  class="bg-gray-800 hover:bg-gray-700 text-white px-2 py-1 rounded">
            Update
          </button>
          <button @click="deleteField(field.id)"
                  class="bg-red-500 text-white px-2 py-1 rounded">
            Delete
          </button>
        </div>
      </div>
    </div>
    <div class="bg-white p-6 rounded-2xl shadow-lg overflow-y-auto">

    </div>
    <PopUp ref="popUp" />
  </div>
</template>

<style scoped>

</style>
