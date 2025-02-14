<script setup lang="ts">
import TopButtonContainer from "@/components/ui/TopButtonContainer.vue";
import PlaybackButtonsContainer from "@/components/playback/PlaybackButtonsContainer.vue";
import VolumeControl from "@/components/playback/VolumeControl.vue";
import ProgressBar from "@/components/playback/ProgressBar.vue";
import PlaylistTabs from "@/components/ui/PlaylistTabs.vue";
import BottomStats from "@/components/ui/BottomStats.vue";
</script>

<template>
  <main>
    <div class="app-container">
      <div class="main-content">
          <div class="flex items-center gap-2">
            <TopButtonContainer />
            <PlaybackButtonsContainer />
            <VolumeControl />
          </div>
        <ProgressBar />
        <PlaylistTabs />
      </div>
      <BottomStats />
    </div>
  </main>
</template>

<style>
.app-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  overflow: hidden;
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  min-height: 0;
  gap: 0.5rem;
}
</style>

<!--<script lang="ts" setup>-->
<!--import {reactive} from 'vue'-->
<!--import {Greet} from '../wailsjs/go/main/App'-->

<!--const data = reactive({-->
<!--  name: "",-->
<!--  resultText: "Please enter your name below 👇",-->
<!--})-->

<!--function greet() {-->
<!--  Greet(data.name).then((result: string) => {-->
<!--    data.resultText = result-->
<!--  })-->
<!--}-->

<!--</script>-->

<!--<template>-->
<!--  <main>-->
<!--    <div class="flex justify-center items-center">-->
<!--      <div class="card">-->
<!--        <div id="result" class="result">{{ data.resultText }}</div>-->
<!--        <div id="input" class="input-box">-->
<!--          <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>-->
<!--          <button class="btn btn-xs sm:btn-sm md:btn-md lg:btn-lg xl:btn-xl" @click="greet">Greet</button>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->
<!--  </main>-->
<!--</template>-->
