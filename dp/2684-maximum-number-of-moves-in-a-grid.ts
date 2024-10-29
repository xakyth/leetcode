function maxMoves(grid: number[][]): number {
  const m = grid.length;
  const n = grid[0].length;
  const dp: number[][] = new Array(m)
    .fill(false)
    .map(() => new Array(n).fill(0));

  const helper = (row: number, col: number): number => {
    if (col == n - 1) {
      return col;
    }
    if (dp[row][col] != 0) {
      return dp[row][col];
    }
    const temp: number[] = new Array();
    if (row - 1 >= 0 && grid[row][col] < grid[row - 1][col + 1]) {
      temp.push(helper(row - 1, col + 1));
    }
    if (grid[row][col] < grid[row][col + 1]) {
      temp.push(helper(row, col + 1));
    }
    if (row + 1 < m && grid[row][col] < grid[row + 1][col + 1]) {
      temp.push(helper(row + 1, col + 1));
    }
    dp[row][col] = Math.max(...temp, col);
    return dp[row][col];
  };

  for (let i = 0; i < m; i++) {
    helper(i, 0);
  }

  return Math.max(...dp.map((arr) => arr[0]));
}

console.log(
  maxMoves([
    [2, 4, 3, 5],
    [5, 4, 9, 3],
    [3, 4, 2, 11],
    [10, 9, 13, 15],
  ])
);
