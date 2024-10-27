function countMaxOrSubsets(nums: number[]): number {
  let result = 0;
  let maxPossible = 0;
  nums.forEach((element) => {
    maxPossible |= element;
  });

  const helper = (i: number, current: number) => {
    if (i == nums.length) {
      return;
    }
    if ((current | nums[i]) === maxPossible) {
      result++;
    }
    helper(i + 1, current | nums[i]);
    helper(i + 1, current);
  };

  helper(0, 0);
  return result;
}

console.log(countMaxOrSubsets([3, 2, 1, 5]));
