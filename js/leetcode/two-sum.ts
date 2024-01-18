// find two digits, their sum should be equal to target

var twoSum = function (nums: number[], target: number) {
  const result = [];
  for (let [index, number] of nums.entries()) {
    for (let i = 0; i < nums.length; i++) {
      if (number + nums[i] === target && i !== index) {
        result.push(index, i);
      }
    }
  }
  const array = result.slice(0, 2);
  return array;
};
