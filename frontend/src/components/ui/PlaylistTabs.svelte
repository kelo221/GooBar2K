<script lang="ts">
  let { tabs, onCreateTab, onRemoveTab } = $props<{
    tabs: Tab[],
    onCreateTab: () => void,
    onRemoveTab: (id: string) => void
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
    />
    <div 
      role="tabpanel"
      class="tab-content bg-base-100 border-base-300 p-6"
      
    >
      <h3>{tab.content.title}</h3>
      <p>{tab.content.text}</p>
    </div>
  {/each}
  <div class="tab" onclick={onCreateTab}>
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-4">
      <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
    </svg>
  </div>
</div>
