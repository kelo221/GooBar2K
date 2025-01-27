<script lang="ts">
  // Props for the tabs data
  export let tabs = [
    { id: 'tabone', label: 'Tab One', content: { title: 'Tab One Content', text: 'Tab content...' } },
    { id: 'tabtwo', label: 'Tab Two', content: { title: 'Tab Two Content', text: 'Tab content...' } },
    { id: 'tabthree', label: 'Tab Three', content: { title: 'Tab Three Content', text: 'Tab content...' } }
  ];


</script>

<div role="tablist" class="tabs tabs-lift">
  {#each tabs as tab, i}
    <input 
      type="radio" 
      name="my_tabs_3" 
      role="tab" 
      class="tab" 
      aria-label={tab.label}
      checked={i === 0} 
    />
    <div class="tab-content bg-base-100 border-base-300 p-6">
      <h3>{tab.content.title}</h3>
      <p>{tab.content.text}</p>
    </div>
  {/each}
  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-4">
    <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
  </svg>
</div>
