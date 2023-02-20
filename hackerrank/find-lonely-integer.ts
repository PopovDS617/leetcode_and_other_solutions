//find a lonely integer in array of duplicates

// var 1
const foo = (array) => {
  let duplicates: any = [];
  duplicates = a.filter((num, index) => {
    return a.indexOf(num) !== index;
  });

  const result: number[] | number = a.filter((item) => {
    if (!duplicates.includes(item)) {
      return item;
    }
  });
  const final: number = result[0];

  return final;
};

// var 2
const foo = (array) => {
  const store = {};

  for (let number in array) {
    const testNumber = array[number];

    if (store[testNumber]) {
      store[testNumber]++;
    } else {
      store[testNumber] = 1;
    }
  }

  const result = [];

  for (let [key, value] of Object.entries(store)) {
    console.log(key, value);
    if (value < 2) {
      result.push(+key);
    }
  }

  return result;
};
