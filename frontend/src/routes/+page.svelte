<script lang="ts">
	import BottomStats from "../components/ui/BottomStats.svelte";
	import PlaylistTabs from "../components/ui/PlaylistTabs.svelte";

    const myTabs = [
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
    ];


    const currentBottomStats = {
    codec: 'WavPack',
    bitrate: '998 kbps',
    frequency: '44100 Hz',
    channelType: 'stereo',
    playtime: '1:58 / 4:14',
  }

  </script>
  <PlaylistTabs tabs={myTabs} />
  <BottomStats currentBottomStats={currentBottomStats} />