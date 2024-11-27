function shortestDistanceAfterQueries(
  n: number,
  queries: number[][]
): number[] {
  const res: number[] = Array();
  const graph: number[][] = new Array(n);
  const dist: number[] = Array(n);
  for (let i = 0; i < n; i++) {
    dist[i] = i;
    if (i < n - 1) {
      graph[i] = [i + 1];
    }
  }

  for (const edge of queries) {
    const v = edge[0];
    const w = edge[1];
    graph[v].push(w);
    if (dist[v] + 1 < dist[w]) {
      dist[w] = dist[v] + 1;
      let queue = [w];
      while (queue.length) {
        const nextQueue = Array();
        for (const u1 of queue) {
          if (u1 === n - 1) {
            continue;
          }
          for (const u2 of graph[u1]) {
            if (dist[u1] + 1 < dist[u2]) {
              dist[u2] = dist[u1] + 1;
              nextQueue.push(u2);
            }
          }
        }
        queue = nextQueue;
      }
    }
    res.push(dist[n - 1]);
  }

  return res;
}
