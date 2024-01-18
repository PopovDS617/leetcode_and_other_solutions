//find single number (it has no pair)

var singleNumber = function (nums: number[]) {
  let duplicates = [] as number[];

  duplicates = nums.filter((value, index) => {
    return nums.indexOf(value) !== index;
  });
  const result = nums.filter((num, index) => {
    if (!duplicates.includes(num)) {
      return num;
    }
  });

  return result;
};
