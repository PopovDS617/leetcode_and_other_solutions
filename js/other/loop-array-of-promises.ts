const processPromise = async (prom: Promise<string | number>) => {
  const result = await prom.then((res) => res).catch((err) => err);
  return result;
};

const waitForResults = async (array: Promise<string | number>[]) => {
  const finalArray: Array<string | number> = [];

  for (const prom of array) {
    const result: string | number = await processPromise(prom);
    if (result) {
      finalArray.push(result);
    }
  }

  return finalArray;
};
const list: Promise<string | number>[] = [
  new Promise((res) => res(1)),
  new Promise((res, rej) => rej('error')),
  new Promise((res) => res(55)),
];

waitForResults(list); //  [1, "error", 55]
