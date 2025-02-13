<template>
  <div class="tabs-container">
    <TabList
      :tabs="tabs"
      :activeIndex="activeTabIndex"
      @create="createTab"
      @remove="removeTab"
      @rename="renameTab"
      @select="setActiveTab"
    />
    <TabPanels
      :tabs="tabs"
      :activeIndex="activeTabIndex"
    />
  </div>
</template>

<script setup lang="ts">
import { useTabStore } from '@/stores/tabStore'
import TabList from './TabList.vue'
import TabPanels from './TabPanels.vue'

const tabStore = useTabStore()
const {
  tabs,
  activeTabIndex,
  createTab,
  removeTab,
  renameTab,
  setActiveTab
} = tabStore
</script>
