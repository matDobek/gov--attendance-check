import { browser } from '$app/environment';
import type { Statue, Vote, Politician } from '$lib/types';

const storeName = {
  statue: 'statues',
  politician: 'politicians',
  vote: 'votes',
}

async function initDB(): Promise<IDBDatabase> {
  return new Promise((resolve, reject) => {
    const request = window.indexedDB.open('GovDB', 1)

    request.onupgradeneeded = () => {
      console.log('IndexedDB :: UPGRADE')
      const db = request.result

      db.createObjectStore(storeName.statue, { keyPath: 'id' })
      db.createObjectStore(storeName.politician, { keyPath: 'id' })
      let vs = db.createObjectStore(storeName.vote, { keyPath: 'id' })

      vs.createIndex('politicianRefId', 'politicianId', { unique: false })
      vs.createIndex('voteRefId', 'statueId', { unique: false })
    }

    request.onsuccess = () => {
      console.log('IndexedDB :: SUCCESS')

      const db = request.result
      resolve(db)
    }

    request.onerror = function () {
      console.log('IndexedDB :: ERROR')
      reject('Error creating IndexedDB')
    }
  });
}

async function fetchWithErr(url: string) {
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

export async function fetchStatues() {
  return await fetchWithErr('http://localhost:8010/proxy/api/v1/statues/');
}

export async function fetchVotes() {
  return await fetchWithErr('http://localhost:8010/proxy/api/v1/votes/');
}

export async function fetchPoliticians() {
  return await fetchWithErr('http://localhost:8010/proxy/api/v1/politicians/');
}

export async function getPoliticians() {
  let result: Politician[] = [];

  result = await fetchPoliticians();

  if (browser) {
    const db = await initDB();
    const transaction = db.transaction([storeName.politician], 'readwrite');
    const store = transaction.objectStore(storeName.politician);

    result.forEach((val) => { store.put(val) });

    transaction.oncomplete = () => { db.close() };
  }

  return result;
}

export async function getStatues() {
  let result: Statue[] = [];

  result = await fetchStatues();

  if (browser) {
    const db = await initDB()
    const transaction = db.transaction([storeName.statue], 'readwrite')
    const store = transaction.objectStore(storeName.statue)

    result.forEach((val) => { store.put(val) });

    transaction.oncomplete = () => { db.close() };
  }

  return result;
}

// TODO: do I want to store this?
// TODO: add try catch for export funcs
// I don't quite get how async workrs? Does it spin another thread?
export async function getVotes() {
  let result: Vote[] = [];

  const db = await initDB()
  const transaction = db.transaction([storeName.vote], 'readwrite')
  const store = transaction.objectStore(storeName.vote)

  result = await fetchVotes();
  result.forEach((val) => { store.put(val) });

  // wrap in promise
  // new Promise((resolve, reject) => {})
  // store
  //   .getAll()
  //   .onsuccess = (event) => { resolve(event.target.result) }

  transaction.oncomplete = () => { db.close() };

  return result;
}
