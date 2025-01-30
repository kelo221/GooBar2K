<script lang="ts">

  import ProgressBar from "./components/playback/ProgressBar.svelte";
  import PlaylistTabs from "./components/ui/PlaylistTabs.svelte";
  import BottomStats from "./components/ui/BottomStats.svelte";
  import PlayBackButtons from "./components/playback/PlayBackButtons.svelte";

  let myTabs : Tab[] = $state([
    {
      id: 'tab1',
      label: 'First Tab',
      content: {
        title: 'Welcome',
        text: 'This is the first tab content'
      }
    },
    {
      id: 'tab2',
      label: 'Second Tab',
      content: {
        title: 'Welcome',
        text: 'This is the second tab content'
      }
    },
  ]);

  function createTab() {
    myTabs = [...myTabs, {
      id: 'tab' + (myTabs.length + 1),
      label: 'Tab ' + (myTabs.length + 1),
      content: {
        title: 'Welcome',
        text: 'This is tab ' + (myTabs.length + 1) + ' content'
      }
    }];
  }

  function removeTab(id: string) {
    myTabs = myTabs.filter(tab => tab.id !== id);
  }

  function renameTab(id: string, newName: string) {
    myTabs = myTabs.map(tab => tab.id === id ? { ...tab, label: newName } : tab);
  }

  const currentBottomStats = {
    codec: 'WavPack',
    bitrate: '998 kbps',
    frequency: '44100 Hz',
    channelType: 'stereo',
    playtime: '1:58 / 4:14',
  }

</script>

<PlayBackButtons/>

<ProgressBar/>
<PlaylistTabs 
  tabs={myTabs} 
  onCreateTab={createTab}
  onRemoveTab={removeTab}
  onRenameTab={renameTab}
/>
<BottomStats currentBottomStats={currentBottomStats} />
