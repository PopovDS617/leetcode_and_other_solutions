Array.prototype.myReduce = function (callback, initialValue) {
  let acc = initialValue === undefined ? undefined : initialValue;

  for (let i = 0; i < this.length; i++) {
    if (acc !== undefined) {
      acc = callback.call(null, acc, this[i], i, this);
    } else {
      acc = this[i];
    }
  }

  if (acc === undefined) {
    throw new TypeError('Reduce of empty array with no initial value');
  }

  return acc;
};
