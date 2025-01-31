<script lang="ts">
  let volume = $state(1); // 0 to 1
</script>

<div class="flex items-center gap-2 p-2">

  <input
    type="range"
    min="0"
    max="1"
    step="0.01"
    class="range range-sm w-24"
    bind:value={volume}
    onchange={() => {
    }}
  />
</div>
