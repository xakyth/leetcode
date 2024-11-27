function compressedString(word: string): string {
  let result = '';
  let count = 0;
  let lastChar = '';
  for (let i = 0; i < word.length; i++) {
    if ((word.charAt(i) !== lastChar && count > 0) || count >= 9) {
      result += count.toString() + lastChar;
      count = 1;
    } else {
      count++;
    }
    lastChar = word.charAt(i);
  }
  if (count > 0) {
    result += count.toString() + lastChar;
  }

  return result;
}
