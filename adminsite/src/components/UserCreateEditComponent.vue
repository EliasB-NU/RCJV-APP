<script setup lang="ts">
import axios from 'axios'
import { ref } from 'vue'
import Cookies from 'js-cookie'

const props = defineProps({
  mode: {
    type: String,
    default: 'edit',
    required: true,
  },
  user: {
    type: Object,
    default: () => {
      return {
        id: 0,
        username: '',
        email: '',
      }
    }
  }
})

const emit = defineEmits(['close', 'finished'])


const username = ref<string>(props.user.username)
const email = ref<string>(props.user.email)
const password = ref<string>('')

const editOrCreate = async () => {
  const url = ref<string>('')
  if (props.mode === 'edit' && props.user !== null) {
    url.value = `/api/users/update/${props.user.id}`
  } else if (props.mode === 'create') {
    url.value = "/api/users/create"
  }
  try {
    await axios
      .post(url.value, {
        username: username.value,
        email: email.value,
        password: password.value,
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

</script>

<template>
  <div class="fixed inset-0 bg-opacity-50 flex items-center justify-center">
    <div class="bg-white p-6 rounded-2xl shadow-lg w-full max-w-2xl max-h-[90vh] overflow-y-auto">
      <h3 class="text-lg font-semibold mb-4">Create User</h3>

      <!-- Form to create user -->
      <div class="mb-4">
        <form @submit.prevent="editOrCreate">
          <div class="grid grid-cols-2 gap-4">
            <!-- Left Column: Basic Information -->
            <div class="w-1/2 pr-4">
              <div class="mb-4">
                <label for="name" class="block">Name:</label>
                <input v-model="username" id="name" type="text" class="input-field border-2 border-b-black rounded-sm" required />
              </div>
              <div class="mb-4">
                <label for="email" class="block">Email:</label>
                <input v-model="email" id="email" type="email" class="input-field border-2 border-b-black rounded-sm" required />
              </div>
              <div class="mb-4">
                <label for="password" class="block">Password:</label>
                <input v-model="password" id="password" type="password" class="input-field border-2 border-b-black rounded-sm" />
              </div>
            </div>
          </div>

          <div class="p-3 flex space-x-2 mt-4 justify-end">
            <button v-if="props.mode === 'create'" type="submit" class="bg-gray-800 text-white px-3 py-1 rounded hover:bg-gray-700">Create</button>
            <button v-else type="submit" class="bg-gray-800 text-white px-3 py-1 rounded hover:bg-gray-700">Update</button>
            <button type="button" @click="emit('close')" class="bg-green-600 text-white px-3 py-1 rounded hover:bg-green-700" >Cancel</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
