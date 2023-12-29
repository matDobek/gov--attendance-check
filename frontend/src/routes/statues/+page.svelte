<script lang="ts">
  import type { PageLoad } from '$types';

  export let data: PageLoad;
</script>

<a href="/">Home</a>
<table>
  <thead>
      <tr>
        <td>Kadencja<td>
        <td>Sesja<td>
        <td>Głosowanie<td>
        <td>Tytuł<td>
      </tr>
  </thead>
  <tbody>
    {#each data.statues as elem (elem.id)}
      <tr>
        <td>{elem.termNumber}</td>
        <td>{elem.sessionNumber}</td>
        <td>{elem.votingNumber}</td>
        <td>{elem.title}</td>
      </tr>
    {/each}
  </tbody>
</table>
