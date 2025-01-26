<script lang="ts">
    interface BottomStatsProps {
        codec: string;
        bitrate: string;
        frequency: string;
        channelType: string;
        playtime: string;
    }
    
    export let currentBottomStats: BottomStatsProps;
</script>

<div class="bottom-stats">
  <div class="bottom-stats-item">
    <span>{currentBottomStats.codec} | {currentBottomStats.bitrate} | {currentBottomStats.frequency} | {currentBottomStats.channelType} | {currentBottomStats.playtime}</span>
  </div>
</div>

<style>
  .bottom-stats {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 100;
    padding-left: 8px;
  }

  .bottom-stats-item {
    display: flex;
    gap: 8px;
  }
</style>