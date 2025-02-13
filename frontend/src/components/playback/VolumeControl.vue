<script setup lang="ts">
import { ref } from "vue";
const volume = ref(1);
</script>

<template>
  <div>
    <input
      type="range"
      min="0"
      max="1"
      step="0.01"
      class="range range-sm w-24"
      v-model="volume"
    />
  </div>
</template>
