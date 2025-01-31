<script lang="ts">
  import Playlist from "../playlist/Playlist.svelte";

  let { tabs, onCreateTab, onRemoveTab, onRenameTab } = $props<{
    tabs: Tab[],
    onCreateTab: () => void,
    onRemoveTab: (id: string) => void,
    onRenameTab: (id: string, newName: string) => void,
    onReorderTab: (id: string, newIndex: number) => void
  }>();

  function handleMouseDown(event: MouseEvent, tab: Tab) {
    // Middle mouse button (button 1)
    if (event.button === 1) {
      event.preventDefault();
      onRemoveTab(tab.id);
    }
  }

  let activeTab = $state(0);
  let tabContainer: HTMLDivElement;

  // Add this function
  function setActiveTab(index: number) {
    activeTab = index;
  }
</script>

<div role="tablist" class="tabs tabs-lift">
  {#each tabs as tab, i}
    <input 
      type="radio" 
      name="my_tabs_3" 
      role="tab" 
      class="tab truncate min-w-[100px] max-w-[200px]" 
      aria-label={tab.label}
      checked={i === activeTab}
      onchange={() => activeTab = i}
      onmousedown={(e) => handleMouseDown(e, tab)}
      ondblclick={() => onRenameTab(tab.id, prompt('Enter new name:') || tab.label)}
    />
  {/each}
  <div class="tab" onclick={onCreateTab}>
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-4">
      <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
    </svg>
  </div>
</div>

<div class="tab-content-container">
  {#each tabs as tab, i}
    <div
      role="tabpanel"
      class="tab-content"
      class:active={i === activeTab}
    >
      <Playlist/>
    </div>
  {/each}
</div>

<style>
  .tab-content-container {
    flex: 1;
    min-height: 0;
    position: relative;
  }

  .tab-content {
    display: none;
    height: 100%;
  }

  .tab-content.active {
    display: block;
  }
</style>