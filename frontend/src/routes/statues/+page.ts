import type { Statue } from '$lib/types';
import { getStatues } from '$lib/cachedFetch';


export async function load() {
  let statues: Statue[] = await getStatues();

  return { statues: statues };
}
