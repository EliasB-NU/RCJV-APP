<script setup lang="ts">
import HeaderComponent from '@/components/HeaderComponent.vue'
import PopUp from '@/components/PopUp.vue'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import Cookies from 'js-cookie'
import TeamCreateComponent from '@/components/TeamCreateComponent.vue'
import TeamEditComponent from '@/components/TeamEditComponent.vue'

interface Team {
  id: number
  name: string
  league: string
  institution: string
  institutionID: number
}

const teams = ref<Team[]>([])

async function fetchTeams() {
  try {
    await axios
      .get('/api/teams', {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json'
        }
      })
      .then(res => {
        teams.value = res.data.data
      })
  } catch (error) {
    popUp.value?.show('Error fetching teams.')
    console.error(error)
  }
}

const deleteTeam = async (id: number) => {
  try {
    await axios
      .delete(`/api/teams/delete/${id}`, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
      .then(() => {
        fetchTeams()
        popUp.value?.show('Team deleted successfully.')
      })
  } catch (error) {
    console.error(error)
    popUp.value?.show('Error deleting team.')
  }
}

const showCreateTeam = ref<boolean>(false)
const showEditTeam = ref<boolean>(false)
const teamToEdit = ref<Team>({} as Team)

interface Institution {
  id: number
  name: string
  numberTeams: string
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
      .then(res => {
        institutions.value = res.data.data
      })
  } catch (error) {
    popUp.value?.show('Error in fetching institutions.')
    console.error(error)
  }
}

const deleteInstitution = async (id: number) => {
  try {
    await axios
      .delete(`/api/institutions/delete/${id}`, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
      .then(() => {
        fetchInstitutions()
        popUp.value?.show('Institution deleted successfully.')
      })
  } catch (error) {
    console.error(error)
    popUp.value?.show('Error deleting institution.')
  }
}

const createName = ref<string>('')
const createInstitution = async () => {
  try {
    await axios
      .post(`/api/institutions/create/${createName.value}`, {}, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
      .then(() => {
        fetchInstitutions()
        createName.value = ''
        popUp.value?.show('Successfully created institution.')
      })
  } catch (error) {
    console.error(error)
    popUp.value?.show('Error deleting institution.')
  }
}

const updateInstitution = async (id: number, index: number) => {
  try {
    await axios
      .post(`/api/institutions/update/${id}/`, {
        name: institutions.value[index].name
      }, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
      .then(() => {
        fetchInstitutions()
        popUp.value?.show('Institution updated successfully.')
      })
  } catch (error) {
    console.error(error)
    popUp.value?.show('Error updating institution.')
  }
}

const popUp = ref<InstanceType<typeof PopUp> | null>(null)

onMounted(async () => {
  await fetchTeams()
  await fetchInstitutions()
})
</script>

<template>
  <HeaderComponent />

  <div class="grid gap-10 md:grid-cols-1 lg:grid-cols-2 space-y-2 p-10">
    <div class="inset-0 bg-opacity-50 flex items-center justify-center">
      <div class="bg-white p-6 rounded-2xl shadow-lg overflow-y-auto">
        <div>
          <h2 class="text-2xl font-semibold mb-4">Teams</h2>
          <button
            @click="showCreateTeam = true"
            class="bg-gray-800 text-white px-4 py-2 rounded-lg hover:bg-gray-700">
            Create User
          </button>
        </div>

        <div class="bg-white shadow rounded-lg">
          <table class="w-full border-collapse">
            <thead>
            <tr class="bg-gray-200">
              <th class="p-3 text-left">Name</th>
              <th class="p-3 text-left">League</th>
              <th class="p-3 text-left">Institution</th>
              <th class="p-3">Actions</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="team in teams" :key="team.id" class="border-t">
              <td class="p-3">{{ team.name }}</td>
              <td class="p-3">{{ team.league }}</td>
              <td class="p-3">{{ team.institution }}</td>
              <td class="p-3 flex space-x-2">
                <button
                  @click="showEditTeam = true; teamToEdit = team"
                  class="bg-gray-800 text-white px-3 py-1 rounded hover:bg-gray-700"
                >
                  Edit
                </button>
                <button
                  @click="deleteTeam(team.id)"
                  class="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600"
                >
                  Delete
                </button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>

        <TeamCreateComponent
          v-if="showCreateTeam"
          @close="showCreateTeam = false"
          @finished="showCreateTeam = false;fetchTeams()"
        />

        <TeamEditComponent
          v-if="showEditTeam"
          :team="teamToEdit"
          @close="showEditTeam = false"
          @finished="showEditTeam = false;fetchTeams()"
        />
      </div>
    </div>
    <div class="inset-0 bg-opacity-50 flex items-center justify-center">
      <div class="bg-white p-6 rounded-2xl shadow-lg overflow-y-auto">
        <div class="mb-4">
          <h2 class="text-2xl font-semibold mb-4">Institutions</h2>
          <button @click="createInstitution" class="bg-gray-700 text-white px-4 py-2 rounded">
            + Create Institution
          </button>
        </div>

        <div v-for="(institution, index) in institutions" :key="institution.id"
             class="flex items-center space-x-2 mb-2">
          <input
            v-model="institution.name"
            class="border p-1 rounded"
            placeholder="Institution Name"
          />
          <span class="text-gray-500">Teams: {{ institution.numberTeams }}</span>
          <button @click="updateInstitution(institution.id, index)"
                  class="bg-gray-700 text-white px-2 py-1 rounded">
            Update
          </button>
          <button @click="deleteInstitution(institution.id)"
                  class="bg-red-500 text-white px-2 py-1 rounded">
            Delete
          </button>
        </div>
      </div>
    </div>
  </div>
  <PopUp ref='popUp' />
</template>

<style scoped>

</style>
