<script setup lang="ts">
import {usePlaylistStore} from "@/stores/playlistStore";
import PlaylistHeader from "./PlaylistHeader.vue";
import { computed, ref } from 'vue';

const sortBy = ref("artist");
const sortDirection = ref<"asc" | "desc" | null>(null);

const playListStore = usePlaylistStore();

// Add computed property for sorted songs
const sortedSongs = computed(() => {
  
  if (!sortDirection.value) return playListStore.currentPlayList.songs;
  
  return [...playListStore.currentPlayList.songs].sort((a, b) => {
    const aValue = a[sortBy.value as keyof typeof a];
    const bValue = b[sortBy.value as keyof typeof b];
    
    
    if (sortDirection.value === "asc") {
      return aValue < bValue ? -1 : aValue > bValue ? 1 : 0;
    } else {
      return bValue < aValue ? -1 : bValue > aValue ? 1 : 0;
    }
  });
});

function handleSort(field: string) {
  if (sortBy.value === field) {
    sortDirection.value = sortDirection.value === "asc" ? "desc" : "asc";
  } else {
    sortBy.value = field;
    sortDirection.value = "asc";
  }
  
}

const lastClickTime = ref(0);

function handleSongClick(song: AudioFile) {
  const currentTime = Date.now();
  playListStore.setSelectedSong(song);
  
  if (currentTime - lastClickTime.value < 200) {
    console.log(`Playing: ${song.title}`);
  }
  lastClickTime.value = currentTime;
}
</script>

<template>
  <div class="playlist-container">
    <table class="table table-xs">
      <thead>
      <tr>
        <PlaylistHeader
          :sortable="true"
          :sortDirection="sortBy === 'artist' ? sortDirection : null"
          :onSort="() => handleSort('artist')"
        >
          Artist / Album
        </PlaylistHeader>

        <PlaylistHeader
          :sortable="true"
          :sortDirection="sortBy === 'track' ? sortDirection : null"
          :onSort="() => handleSort('track')"
        >
          Track
        </PlaylistHeader>

        <PlaylistHeader
          :sortable="true"
          :sortDirection="sortBy === 'title' ? sortDirection : null"
          :onSort="() => handleSort('title')"
        >
          Title
        </PlaylistHeader>

        <PlaylistHeader
          :sortable="true"
          :sortDirection="sortBy === 'duration' ? sortDirection : null"
          :onSort="() => handleSort('duration')"
        >
          Duration
        </PlaylistHeader>
      </tr>
      </thead>
      <tbody>

      <tr v-for="song in sortedSongs" 
          :key="song.id" 
          @click="() => handleSongClick(song)"
          class="select-none"
          :class="{ 'selected': playListStore.selectedSong?.id === song.id }"
      >
        <td>{{ song.artist }} - {{ song.album }}</td>
        <td>{{ song.track }}</td>
        <td>{{ song.title }}</td>
        <td>{{ song.duration }}</td>
      </tr>

      </tbody>
    </table>
  </div>
</template>

<style>
.playlist-container {
  height: 100%;
  overflow-y: auto;
}

.sticky-header th {
  position: sticky;
  top: 0;
  z-index: 20;
}

.selected {
    background-color: gray;
}
</style>
