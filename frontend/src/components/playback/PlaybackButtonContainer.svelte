<script>
      import PlayNextButton from "./PlaybackButtons/PlayNextButton.svelte";

      import PlayPauseButton from "./PlaybackButtons/PlayPauseButton.svelte";

      import PlayPreviousButton from "./PlaybackButtons/PlayPreviousButton.svelte";
      import PlayRandomButton from "./PlaybackButtons/PlayRandomButton.svelte";
      import PlayStopButton from "./PlaybackButtons/PlayStopButton.svelte";
</script>

<div class="flex w-full justify-center items-center p-2 gap-1">
      <PlayStopButton />
      <PlayPauseButton />
      <PlayPreviousButton />
      <PlayNextButton />
      <PlayRandomButton />
</div>
