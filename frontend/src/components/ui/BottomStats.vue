<script setup lang="ts">
import {useStatStore} from "@/stores/statStore";
import {computed} from "vue";

const statStore = useStatStore();
const currentBottomStats = computed(() => statStore.stats);
</script>

<template>
  <footer class="footer sm:footer-horizontal bg-neutral text-neutral-content items-center p-1">
    <span class="text-xs">{{ currentBottomStats.codec }} | {{
        currentBottomStats.bitrate
      }} | {{ currentBottomStats.frequency }} | {{ currentBottomStats.channelType }} | {{
        currentBottomStats.playtime
      }}</span>
  </footer>
</template>

<style>
.footer {
  position: relative;
  bottom: auto;
  left: auto;
  right: auto;
  border-top: 1px solid var(--border-color);
}
</style>
