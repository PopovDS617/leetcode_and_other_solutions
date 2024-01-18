const callbackAtMostOnce = (callback) => {
  let hasBeenCalled = false;

  return function (...args) {
    if (!hasBeenCalled) {
      hasBeenCalled = true;
      return callback(...args);
    }
  };
};

export const callbackAtMostN = (callback, n) => {
  let count = 0;

  return function (...args) {
    while (count < n) {
      count++;
      return callback(...args);
    }
  };
};
