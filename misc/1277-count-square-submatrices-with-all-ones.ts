function countSquares(matrix: number[][]): number {
  const m = matrix.length;
  const n = matrix[0].length;
  let result = 0;

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (matrix[i][j] === 1) {
        result++;
        for (let k = i + 1, l = j + 1; k < m && l < n; k++, l++) {
          let rowOnes = true;
          for (let p = l; p >= j; p--) {
            if (matrix[k][p] != 1) {
              rowOnes = false;
              break;
            }
          }
          if (!rowOnes) {
            break;
          }
          let colOnes = true;
          for (let p = k; p >= i; p--) {
            if (matrix[p][l] != 1) {
              colOnes = false;
              break;
            }
          }
          if (rowOnes && colOnes) {
            result++;
          } else {
            break;
          }
        }
      }
    }
  }

  return result;
}

const matrix = [
  [0, 1, 1, 1],
  [1, 1, 1, 1],
  [0, 1, 1, 1],
];

console.log(countSquares(matrix));
