<script lang="ts">
	import ProgressBar from "../components/playback/ProgressBar.svelte";
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

<div class="flex flex-col min-h-screen">
    <div class="flex-grow">
        <PlaylistTabs tabs={myTabs} />
    </div>
    <div class="mb-8">
        <ProgressBar />
    </div>
    <BottomStats currentBottomStats={currentBottomStats} />
</div>

