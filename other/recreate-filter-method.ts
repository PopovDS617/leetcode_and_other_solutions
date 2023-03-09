Array.prototype.customFilter = function (callback) {
  const result = [];

  for (let i = 0; i < this.length; i++) {
    const shouldPush = callback(this[i], i, this);
    if (shouldPush) {
      result.push(this[i]);
    }
  }

  return result;
};
