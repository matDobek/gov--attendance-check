export const statuesStore = 'Statues'
export const politiciansStore = 'Politicians'
export const votesStore = 'Votes'

export async function createIndexedDB() {
  return new Promise((resolve, reject) => {
    const request = window.indexedDB.open('GovDB', 1)

    request.onupgradeneeded = (event) => {
      console.log('IndexedDB :: UPGRADE')
      const db = event.target.result

      db.createObjectStore(statuesStore, { keyPath: 'id' })
      db.createObjectStore(politiciansStore, { keyPath: 'id' })
      let vs = db.createObjectStore(votesStore, { keyPath: 'id' })

      vs.createIndex('politicianRefId', 'politicianId', { unique: false })
      vs.createIndex('voteRefId', 'statueId', { unique: false })
    }

    request.onsuccess = (event) => {
      console.log('IndexedDB :: SUCCESS')

      const db = event.target.result
      resolve(db)
    }

    request.onerror = function (event) {
      console.log('IndexedDB :: ERROR')
      reject('Error creating IndexedDB')
    }
  });
}
