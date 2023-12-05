import type { Statue, Vote, Politician } from '$lib/types';
import type { PageLoad } from '$types';

let statues: Statue[] = [];
let votes: Vote[] = [];
let politicians: Politician[] = [];

async function fetchWithErr(fetch, url: string) {
  console.log(`Fetching ${url}`);

  try {
    const response = await fetch(url);
    if (response.ok) {
      return await response.json();
    } else {
      throw new Error(`failed to fetch ${url}`);
    }
  } catch (error) {
    console.error('Error:', error);
  }
}

async function fetchStatues(fetch) {
  statues = await
    fetchWithErr(fetch, 'http://localhost:8010/proxy/api/v1/statues/');
}

async function fetchVotes(fetch) {
  votes = await
    fetchWithErr(fetch, 'http://localhost:8010/proxy/api/v1/votes/');
}

async function fetchPoliticians(fetch) {
  politicians = await
    fetchWithErr(fetch, 'http://localhost:8010/proxy/api/v1/politicians/');
}

export async function load({fetch, params}: PageLoad) {
  await Promise.all([
    fetchStatues(fetch),
    fetchVotes(fetch),
    fetchPoliticians(fetch),
  ]);

  return {
    statues: statues,
    votes: votes,
    politicians: politicians,
  }
}
