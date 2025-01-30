<script lang="ts">
  import Playlist from "../playlist/Playlist.svelte";

  let { tabs, onCreateTab, onRemoveTab, onRenameTab } = $props<{
    tabs: Tab[],
    onCreateTab: () => void,
    onRemoveTab: (id: string) => void,
    onRenameTab: (id: string, newName: string) => void,
    onReorderTab: (id: string, newIndex: number) => void
  }>();
  let activeTab = $state(0);

  function handleMouseDown(event: MouseEvent, tab: Tab) {
    // Middle mouse button (button 1)
    if (event.button === 1) {
      event.preventDefault();
      onRemoveTab(tab.id);
    }
  }
</script>

<div role="tablist" class="tabs tabs-lift">
  {#each tabs as tab, i}
    <input 
      type="radio" 
      name="my_tabs_3" 
      role="tab" 
      class="tab" 
      aria-label={tab.label}
      checked={i === activeTab}
      onmousedown={(e) => handleMouseDown(e, tab)}
      ondblclick={() => onRenameTab(tab.id, prompt('Enter new name:') || tab.label)}
    />
    <div
      role="tabpanel"
      class="tab-content"
    >
      <Playlist/>
    </div>
  {/each}
  <div class="tab" onclick={onCreateTab}>
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-4">
      <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
    </svg>
  </div>
</div>
