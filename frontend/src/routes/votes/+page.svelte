<script lang="ts">
  import { onMount } from 'svelte';
  import { politiciansStore, votesStore, statuesStore, createIndexedDB } from '$lib/indexedDB'

  export let data;
  export let elems = [];

  onMount(async () => {
    try {
      const db = await createIndexedDB();
      const transaction = db.transaction([
        statuesStore,
        politiciansStore,
        votesStore,
      ], 'readwrite');

      const pStore = transaction.objectStore(politiciansStore)
      const vStore = transaction.objectStore(votesStore)
      const sStore = transaction.objectStore(statuesStore)

      data.statues.forEach(async (val) => {
        await sStore.put(val)
      })

      data.votes.forEach(async (val) => {
        await vStore.put(val)
      })
      data.politicians.forEach(async (val) => {
        await pStore.put(val)
      })

      let votes
      let politicians
      let statues

      let getVotes = vStore.getAll()
      getVotes.onsuccess = (event) => {
        votes = event.target.result

        let getPoliticians = pStore.getAll()
        getPoliticians.onsuccess = (event) => {
          politicians = event.target.result

          let getStatues= sStore.getAll()
          getStatues.onsuccess = (event) => {
            statues = event.target.result

            votes.forEach((vote) => {

              let politician = politicians.find((politician) => {
                return politician.id === vote.politicianId
              })
              let statue = statues.find((statue) => {
                return statue.id === vote.statueId
              })

              let elem = {
                id: vote.id,
                response: vote.response,

                statueId: statue.id,
                statueTitle: statue.title,
                statueSessionNumber: statue.sessionNumber,
                statueTermNumber: statue.termNumber,
                statueVotingNumber: statue.votingNumber,

                politicianId: politician.id,
                politicianName: politician.name,
                politicianParty: politician.party,
              }
              elems.push(elem)
            })
          }
        }
      }

      // close connection
      transaction.oncomplete = () => {
        db.close()
        elems = [...elems];
        console.log(elems)
      }

    } catch (error) {
      console.error('Error creating IndexedDB:', error);
    }
  })
</script>

<a href="/">Home</a>
<table>
  <tbody>
    {#each elems as elem (elem.id)}
      <tr>
        <td>{elem.statueSessionNumber}</td>
        <td>{elem.statueTermNumber}</td>
        <td>{elem.statueVotingNumber}</td>

        <td>{elem.statueTitle}</td>

        <td>{elem.politicianName}</td>
        <td>{elem.politicianParty}</td>

        <td>{elem.response}</td>
      </tr>
    {/each}
  </tbody>
</table>
