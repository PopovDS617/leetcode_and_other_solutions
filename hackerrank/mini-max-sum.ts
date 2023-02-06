//sort the array, find sum of 1-4 and 2-5

function miniMaxSum(arr: number[]) {
  const sorted = arr.sort((a, b) => a - b);
  const minArray = sorted.slice(0, 4);
  const min = minArray.reduce((c, a) => c + a);
  const maxArray = sorted.slice(1, 5);
  const max = maxArray.reduce((c, a) => c + a);

  const result = min.toString() + ' ' + max.toString();
  console.log(result);
  return result;
}

console.log(miniMaxSum([7, 69, 2, 221, 8974]));
