<script lang="ts">
    export let sortable = false;
    export let sortDirection: 'asc' | 'desc' | null = null;
    export let onSort: () => void = () => {};
  </script>
  
  <th
    class="group cursor-pointer select-none bg-base-200 hover:bg-base-300 transition-colors !bg-opacity-100"
    on:click={sortable ? onSort : () => {}}
  >
    <div class="flex items-center gap-2">
      <slot />
      {#if sortable}
        <div class="text-xs opacity-50 group-hover:opacity-100">
          {#if sortDirection === 'asc'}
            ↑
          {:else if sortDirection === 'desc'}
            ↓
          {:else}
          ‎ 
          {/if}
        </div>
      {/if}
    </div>
  </th>
  
  <style>
    th {
      position: sticky;
      top: 0;
      z-index: 20;
      padding: 0.75rem 1rem;
      font-weight: 600;
      background-color: hsl(var(--b2));
      box-shadow: 0 1px 0 0 hsl(var(--bc)/0.1);
    }
  </style>