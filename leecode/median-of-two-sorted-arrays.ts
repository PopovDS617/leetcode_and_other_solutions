//sort two array and find median

var findMedianSortedArrays = function (nums1: number[], nums2: number[]) {
  const mergedArray = nums1.concat(nums2);
  const sortedArray = mergedArray.sort((a: number, b: number) => a - b);
  const middle = Math.ceil(mergedArray.length / 2 - 1);
  console.log(mergedArray, middle);
  let result = sortedArray[middle];
  if (mergedArray.length % 2 === 0) {
    result = (sortedArray[middle] + sortedArray[middle + 1]) / 2;
  }

  return result;
};
