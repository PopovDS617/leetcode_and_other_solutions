const memoize = (foo) => {
  const cache = {};

  const inner = (data) => {
    if (cache[data]) {
      console.log('from cache');
      return cache[data];
    }

    const newData = foo(data);
    cache[data] = newData;
    console.log('newly calculated');
    return newData;
  };

  return inner;
};

const multiplication = (data) => data * 2;

const calculate = memoize(multiplication);
