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
        league: 'soccerEntry'
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
        ...fields.value[index]
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

const selectedFile = ref<File | null>(null)
const selectedOption = ref<string>('')

// Handle file input change
function handleFileChange(event: Event) {
  const input = event.target as HTMLInputElement
  if (input.files && input.files.length > 0) {
    selectedFile.value = input.files[0]
    console.log('File selected:', selectedFile.value.name)
  }
}

// Handle file submission
const uploadXLSX = async () => {
  if (!selectedFile.value) {
    alert('Please upload a file')
    return
  }
  if (!selectedOption.value) {
    alert('Please select a league')
    return
  }
  try {
    await axios
      .post(`/api/matches/upload/${selectedOption.value}`, {
        matches: selectedFile.value
      }, {
        headers: {
          'Content-Type': 'multipart/form-data',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
  } catch (error) {
    console.log(error)
    popUp.value?.show("Error submitting new XLSX file")
  }
}

// Handle dropdown-related action
const newXLSX = async () => {
  if (!selectedOption.value) {
    alert('Please select a league')
    return
  }
  try {
    await axios
      .get(`/api/matches/generate/${selectedOption.value}`, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token'),
        },
        responseType: 'blob'
      })
    .then(res => {
      popUp.value?.show('Successfully generated XLSX file.')

      const blob = new Blob([res.data]);

      // Just use your own name
      const filename = `matches_template_${selectedOption.value}.xlsx`

      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.setAttribute('download', filename);
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
      window.URL.revokeObjectURL(url);
    })
  } catch (error) {
    console.log(error)
    popUp.value?.show('Error generating new XLSX file.')
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
    <!-- File Upload Section -->
    <div class="p-4 border rounded shadow mb-4 max-w-md">
      <h2 class="text-lg font-bold mb-2">Upload XLSX</h2>

      <input type="file" @change="handleFileChange" class="mb-2 text-blue-400 hover:underline" />

      <button
        @click="uploadXLSX"
        class="bg-gray-800 hover:bg-gray-700 text-white px-4 py-2 rounded disabled:opacity-50"
      >
        Upload
      </button>

      <!-- Red Warning Box -->
      <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mt-4">
        ⚠️ Warning: All previous matches will be deleted     <br>
        ⚠️ Warning: Don't forget to select your league below
      </div>

      <!-- Sub-box with Select and Button -->
      <div class="mt-4 p-3 border border-gray-300 rounded bg-gray-50">
        <label class="block mb-1 font-semibold">Select Option</label>
        <div class="flex items-center space-x-2">
          <select v-model="selectedOption" class="border p-1 rounded flex-1">
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

          <button
            @click="newXLSX"
            class="bg-gray-800 hover:bg-gray-700 text-white px-3 py-1 rounded"
          >
            Generate new XLSX
          </button>
        </div>
      </div>
    </div>
    <PopUp ref="popUp" />
  </div>
</template>

<style scoped>

</style>
