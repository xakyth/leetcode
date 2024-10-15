function minimumSteps(s: string): number {
  let res = 0;
  for (let l = 0, r = s.length - 1; l <= r; ) {
    if (s[l] != '1') {
      l++;
      continue;
    }
    if (s[r] != '0') {
      r--;
      continue;
    }
    res += r - l;
    r--;
    l++;
  }
  return res;
}
