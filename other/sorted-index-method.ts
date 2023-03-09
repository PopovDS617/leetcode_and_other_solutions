export const sortedIndex = (arr, value) => {
  if (arr.length === 0) {
    return 0;
  }
  let result;

  for (let i = 0; i < arr.length; i++) {
    if (value === arr[i]) {
      result = i;
      break;
    }
    result = i;
    if (value < arr[i]) {
      break;
    } else {
      result = i + 1;
    }
  }

  return result;
};
