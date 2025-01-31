<script lang="ts">
  import ProgressBar from "./components/playback/ProgressBar.svelte";
  import PlaylistTabs from "./components/ui/PlaylistTabs.svelte";
  import BottomStats from "./components/ui/BottomStats.svelte";
  import PlayBackButtons from "./components/playback/PlaybackButtonContainer.svelte";
  import TopButtonContainer from "./components/ui/TopButtonContainer.svelte";
  import VolumeControl from "./components/playback/VolumeControl.svelte";

  let myTabs: Tab[] = $state([
    {
      id: "tab1",
      label: "First Tab",
      content: {
        title: "Welcome",
        text: "This is the first tab content",
      },
    },
    {
      id: "tab2",
      label: "Second Tab",
      content: {
        title: "Welcome",
        text: "This is the second tab content",
      },
    },
  ]);

  function createTab() {
    myTabs = [
      ...myTabs,
      {
        id: "tab" + (myTabs.length + 1),
        label: "Tab " + (myTabs.length + 1),
        content: {
          title: "Welcome",
          text: "This is tab " + (myTabs.length + 1) + " content",
        },
      },
    ];
  }

  function removeTab(id: string) {
    myTabs = myTabs.filter((tab) => tab.id !== id);
  }

  function renameTab(id: string, newName: string) {
    myTabs = myTabs.map((tab) =>
      tab.id === id ? { ...tab, label: newName } : tab,
    );
  }

  const currentBottomStats = {
    codec: "WavPack",
    bitrate: "998 kbps",
    frequency: "44100 Hz",
    channelType: "stereo",
    playtime: "1:58 / 4:14",
  };
</script>

<div class="app-container">
  <div class="main-content">
    <div class="flex">
      <TopButtonContainer />
      <PlayBackButtons />
      <VolumeControl />
    </div>

    <ProgressBar />
    <PlaylistTabs
      tabs={myTabs}
      onCreateTab={createTab}
      onRemoveTab={removeTab}
      onRenameTab={renameTab}
      class="flex-1 min-h-0"
    />
  </div>
  <BottomStats {currentBottomStats} />
</div>

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