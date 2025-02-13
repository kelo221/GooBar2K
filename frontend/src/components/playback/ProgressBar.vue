<script setup lang="ts">
import { ref } from 'vue';

const progress = ref<ProgressBarState['progress']>(40);
</script>

<template>
  <div class="w-full px-2">
    <input type="range" min="0" max="1000"
           v-model="progress" class="range range-xs range-primary w-full"/>
  </div>
</template>
