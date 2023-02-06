var isPalindrome = function (x: number) {
  const line = x.toString();
  const reversedLine = x.toString().split('').reverse().join('');
  console.log(line, reversedLine);

  return line === reversedLine ? true : false;
};
