function getMaximumXor(nums: number[], maximumBit: number): number[] {
  const res: number[] = Array();
  function getBestValueForXor(n: number, maximumBit: number): number {
    let res = 0;
    console.log('res', res);
    for (let i = 0; i < maximumBit; i++) {
      res = res << 1;
      if (!(n & 1)) {
        res += 1;
      }
    }
    console.log('res2', res);
    return res;
  }
  let tmp: number;
  nums.forEach((n) => {
    if (!tmp) {
      tmp = n;
    }
    tmp = tmp ^ n;
    res.push(getBestValueForXor(tmp, maximumBit));
  });
  return res;
}

console.log('getMaximumXor', getMaximumXor([0, 1, 1, 3], 2));
