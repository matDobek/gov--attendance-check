<script lang="ts">
  import type { PageLoad } from '$types';

  export let data: PageLoad;
</script>

<a href="/">Home</a>
<table>
  <thead>
      <tr>
        <td>lp<td>
        <td>Imię<td>
        <td>Partia<td>
      </tr>
  </thead>
  <tbody>
    {#each data.politicians as elem, i (i)}
      <tr>
        <td>{[i+1]}</td>
        <td>{elem.name}</td>
        <td>{elem.party}</td>
      </tr>
    {/each}
  </tbody>
</table>
