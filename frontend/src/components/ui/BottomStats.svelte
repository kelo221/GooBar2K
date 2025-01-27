<script lang="ts">
    export let currentBottomStats: BottomStats;
</script>

<footer class="footer sm:footer-horizontal bg-neutral text-neutral-content items-center p-1">
    <span class="text-xs">{currentBottomStats.codec} | {currentBottomStats.bitrate} | {currentBottomStats.frequency} | {currentBottomStats.channelType} | {currentBottomStats.playtime}</span>
</footer>

<style>
  .footer {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 100
  }

</style>