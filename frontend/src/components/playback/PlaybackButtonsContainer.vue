<script setup lang="ts">
import PlayStopButton from "@/components/playback/playbackButtons/PlayStopButton.vue";
import PlayPauseButton from "@/components/playback/playbackButtons/PlayPauseButton.vue";
import PlayPreviousButton from "@/components/playback/playbackButtons/PlayPreviousButton.vue";
import PlayNextButton from "@/components/playback/playbackButtons/PlayNextButton.vue";
import PlayRandomButton from "@/components/playback/playbackButtons/PlayRandomButton.vue";


import { PauseRequest, StopRequest, PlayRequest,Greet } from '../../../wailsjs/go/main/App';

function StopPlayback() {
  console.log("Stopping playback");
  StopRequest();
}

function PlayPausePlayback() {
  console.log("Pausing playback");
  PauseRequest();
}

function PlayPrevious() {
  Greet("test");
  console.log("Playing previous song");
}

function PlayNext() {
  console.log("Playing next song");
}

function PlayRandom() {
  console.log("Playing random song");
  PlayRequest();
}


</script>

<template>
  <div>
    <PlayStopButton @click="StopPlayback"/>
    <PlayPauseButton @click="PlayPausePlayback"/>
    <PlayPreviousButton @click="PlayPrevious" />
    <PlayNextButton @click="PlayNext" />
    <PlayRandomButton @click="PlayRandom" />
  </div>
</template>

