function canSortArray(nums: number[]): boolean {
  function countBits(n: number): number {
    let x = 1;
    let res = 0;
    for (let i = 0; i <= 8; i++) {
      if ((x & n) >= 1) {
        res++;
      }
      x = x << 1;
    }
    return res;
  }
  for (let i = 0; i < nums.length - 1; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[i] > nums[j]) {
        if (countBits(nums[i]) !== countBits(nums[j])) {
          return false;
        }
        const tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
      }
    }
  }
  return true;
}

console.log(canSortArray([1, 256, 64]));
