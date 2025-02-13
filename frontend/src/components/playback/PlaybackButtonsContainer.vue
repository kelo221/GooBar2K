<script setup lang="ts">
import PlayStopButton from "@/components/playback/playbackButtons/PlayStopButton.vue";
import PlayPauseButton from "@/components/playback/playbackButtons/PlayPauseButton.vue";
import PlayPreviousButton from "@/components/playback/playbackButtons/PlayPreviousButton.vue";
import PlayNextButton from "@/components/playback/playbackButtons/PlayNextButton.vue";
import PlayRandomButton from "@/components/playback/playbackButtons/PlayRandomButton.vue";
</script>

<template>
  <div>
    <PlayStopButton/>
    <PlayPauseButton/>
    <PlayPreviousButton/>
    <PlayNextButton/>
    <PlayRandomButton/>
  </div>
</template>
