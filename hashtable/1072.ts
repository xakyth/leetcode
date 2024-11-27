function maxEqualRowsAfterFlips(matrix: number[][]): number {
  const rowsCount = new Map<string, number>();

  function flipArray(arr: number[]): number[] {
    return arr.map((x) => (x == 0 ? 1 : 0));
  }

  for (const row of matrix) {
    const strRow = row.join('');

    const inversedStrRow = flipArray(row).join('');
    if (rowsCount.has(strRow)) {
      const val = rowsCount.get(strRow);
      if (val) {
        rowsCount.set(strRow, val + 1);
      }
    } else if (rowsCount.has(inversedStrRow)) {
      const val = rowsCount.get(inversedStrRow);
      if (val) {
        rowsCount.set(inversedStrRow, val + 1);
      }
    } else {
      rowsCount.set(strRow, 1);
    }
  }

  return Math.max(...Array.from(rowsCount.values()));
}

const res = maxEqualRowsAfterFlips([
  [0, 0, 0],
  [0, 0, 1],
  [1, 1, 0],
]);

console.log(res);
