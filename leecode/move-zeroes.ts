// move 0,2,0,12 to 2,12,00

var moveZeroes = function (nums: number[]) {
  let pointer = 0;
  const length = nums.length;
  for (let i = 0; i < length; i++) {
    if (nums[i] !== 0) {
      nums[pointer] = nums[i];
      pointer++;
    }
  }

  for (let i = pointer; i < length; i++) {
    nums[i] = 0;
  }
  return nums;
};
