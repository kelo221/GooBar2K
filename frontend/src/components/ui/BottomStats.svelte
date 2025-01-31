<script lang="ts">
    export let currentBottomStats: BottomStats;
</script>

<footer class="footer sm:footer-horizontal bg-neutral text-neutral-content items-center p-1">
    <span class="text-xs">{currentBottomStats.codec} | {currentBottomStats.bitrate} | {currentBottomStats.frequency} | {currentBottomStats.channelType} | {currentBottomStats.playtime}</span>
</footer>

<style>
  .footer {
    position: relative;
    bottom: auto;
    left: auto;
    right: auto;
    border-top: 1px solid var(--border-color); /* Optional separator */
  }
  </style>
  