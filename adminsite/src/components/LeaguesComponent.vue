<script setup lang="ts">

import { onMounted, ref } from 'vue'
import axios from 'axios'
import Cookies from 'js-cookie'

interface Leagues {
  rescueLineEntry: boolean;
  rescueLine: boolean;
  rescueMazeEntry: boolean;
  rescueMaze: boolean;

  soccerEntry: boolean;
  soccerLightWeightEntry: boolean;
  soccerLightWeight: boolean;
  soccerOpen: boolean;

  onStageEntry: boolean;
  onStage: boolean;
}

const leagues = ref<Leagues>({} as Leagues)

const fetchLeagues = async () => {
  try {
    await axios
      .get('/api/leagues')
      .then((res) => {
        leagues.value = res.data
      })
  } catch (error) {
    console.log(error)
  }
}

const updateLeagues = async () => {
  try {
    await axios
      .patch('/api/leagues/update', {
        ...leagues.value
      }, {
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
          'Authorization': 'Bearer ' + Cookies.get('token')
        }
      })
  } catch (error) {
    console.log(error)
  }
}

onMounted(fetchLeagues)

</script>

<template>
 <div class="fixed inset-0 bg-opacity-50 flex items-center justify-center">
   <div class="bg-white p-6 rounded-2xl shadow-lg overflow-y-auto">
     <h3 class="text-lg font-semibold mb-4">Leagues:</h3>

     <div class="mb-4">
       <form @submit.prevent="updateLeagues">
         <div class="space-y-2">
           <div class="flex justify-between items-center">
             <label for="rescueLineEntry">Rescue Line Entry:</label>
             <input v-model="leagues.rescueLineEntry" type="checkbox" id="rescueLineEntry" />
           </div>

           <div class="flex justify-between items-center">
             <label for="rescueLine">Rescue Line:</label>
             <input v-model="leagues.rescueLine" type="checkbox" id="rescueLine" />
           </div>

           <div class="flex justify-between items-center">
             <label for="rescueMazeEntry">Rescue Maze Entry:</label>
             <input v-model="leagues.rescueMazeEntry" type="checkbox" id="rescueMazeEntry" />
           </div>

           <div class="flex justify-between items-center">
             <label for="rescueMaze">Rescue Maze:</label>
             <input v-model="leagues.rescueMaze" type="checkbox" id="rescueMaze" />
           </div>

           <div class="flex justify-between items-center">
             <label for="soccerEntry">Soccer Entry:</label>
             <input v-model="leagues.soccerEntry" type="checkbox" id="soccerEntry" />
           </div>

           <div class="flex justify-between items-center">
             <label for="soccerLightWeightEntry">Soccer Lightweight Entry:</label>
             <input v-model="leagues.soccerLightWeightEntry" type="checkbox" id="soccerLightWeightEntry" />
           </div>

           <div class="flex justify-between items-center">
             <label for="soccerLightWeight">Soccer Lightweight:</label>
             <input v-model="leagues.soccerLightWeight" type="checkbox" id="soccerLightWeight" />
           </div>

           <div class="flex justify-between items-center">
             <label for="soccerOpen">Soccer Open:</label>
             <input v-model="leagues.soccerOpen" type="checkbox" id="soccerOpen" />
           </div>
         </div>

         <div class="p-3 flex space-x-2 mt-4 justify-end">
           <button type="submit" class="bg-gray-800 text-white px-3 py-1 rounded hover:bg-gray-700">Update</button>
           <button type="button" @click="fetchLeagues" class="bg-green-600 text-white px-3 py-1 rounded hover:bg-green-700">Reset</button>
         </div>
       </form>
     </div>
   </div>
 </div>
</template>

<style scoped>

</style>
