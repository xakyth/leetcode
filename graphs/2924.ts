function findChampion(n: number, edges: number[][]): number {
  const unseen: Set<number> = new Set();
  for (let i = 0; i < n; i++) {
    unseen.add(i);
  }
  for (const edge of edges) {
    unseen.delete(edge[1]);
  }
  if (unseen.size === 1) {
    return unseen.values().next().value;
  }
  return -1;
}
