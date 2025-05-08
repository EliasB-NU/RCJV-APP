<script setup lang="ts">

import { onMounted, ref } from 'vue'
import axios from 'axios'
import Cookies from 'js-cookie'
import UserCreateEditComponent from '@/components/UserCreateEditComponent.vue'

interface User {
  id: number;
  username: string;
  email: string;
}

const users = ref<User[]>([])

async function fetchUsers() {
  try {
    await axios
      .get('/api/users', {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
      .then((res) => {
        users.value = res.data
      })
  } catch (error) {
    console.error(error)
  }
}

const deleteUser = async (id: number) => {
  try {
    await axios
      .delete(`/api/users/${id}`, {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
    .then(() => {
      fetchUsers()
    })
  } catch (error) {
    console.error(error)
  }
}

const showCreateEditComponent = ref<boolean>(false)
const mode = ref<string>('create')
const userToEdit = ref<User>({} as User)

onMounted(fetchUsers())
</script>

<template>
  <div class="inset-0 bg-opacity-50 flex items-center justify-center">
    <div class="bg-white p-6 rounded-2xl shadow-lg overflow-y-auto">
      <h3 class="text-lg font-semibold mb-4">Users:</h3>
      <button
        @click="mode = 'create'; showCreateEditComponent = true" class="bg-gray-800 text-white px-4 py-2 rounded-lg hover:bg-gray-700">
        Create User
      </button>

      <!-- Users Table -->
      <div class="bg-white shadow rounded-lg">
        <table class="w-full border-collapse">
          <thead>
          <tr class="bg-gray-200">
            <th class="p-3 text-left">Name</th>
            <th class="p-3 text-left">Email</th>
            <th class="p-3">Actions</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="user in users" :key="user.id" class="border-t">
            <td class="p-3">{{ user.username }}</td>
            <td class="p-3">{{ user.email }}</td>
            <td class="p-3 flex space-x-2">
              <button
                @click="mode = 'edit'; userToEdit=user; showCreateEditComponent = true"
                class="bg-gray-800 text-white px-3 py-1 rounded hover:bg-gray-700"
              >
                Edit
              </button>
              <button
                @click="deleteUser(user.id)"
                class="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600"
              >
                Delete
              </button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>

      <UserCreateEditComponent
        v-if="showCreateEditComponent"
        :mode="mode"
        :user="userToEdit"
        @close="showCreateEditComponent = false; userToEdit = {} as User"
        @finished="showCreateEditComponent = false; userToEdit = {} as User; fetchUsers()"
      />

    </div>
  </div>
</template>

<style scoped>

</style>
