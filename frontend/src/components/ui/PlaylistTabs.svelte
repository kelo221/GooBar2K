<script lang="ts">
  // Props for the tabs data
  export let tabs = [
    { id: 'tabone', label: 'Tab One', content: { title: 'Tab One Content', text: 'Tab content...' } },
    { id: 'tabtwo', label: 'Tab Two', content: { title: 'Tab Two Content', text: 'Tab content...' } },
    { id: 'tabthree', label: 'Tab Three', content: { title: 'Tab Three Content', text: 'Tab content...' } }
  ];


</script>

<div class="cs-tabs">
  {#each tabs as tab, i}
    <input
      class="radiotab"
      name="tabs"
      tabindex="1"
      type="radio"
      id={tab.id}
      checked={i === 0}
    />
    <label class="label" for={tab.id}>{tab.label}</label>
    <div class="panel" tabindex="1">
      <h2>{tab.content.title}</h2>
      <p>{tab.content.text}</p>
    </div>
  {/each}
</div>

<style>
  .cs-tabs {
    display: flex;
    flex-wrap: wrap;
    height: 100%;
  }

</style>