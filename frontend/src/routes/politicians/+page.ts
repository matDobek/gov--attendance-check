import type { Politician } from '$lib/types';
import { getPoliticians } from '$lib/cachedFetch';


export async function load() {
  let politicians: Politician[] = await getPoliticians();

  return { politicians: politicians };
}
