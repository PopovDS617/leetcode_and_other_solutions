Array.prototype.myFindIndex = function (callback, thisArg) {
  let index = -1;

  for (let i = 0; i < this.length; i++) {
    const result = callback.call(thisArg, this[i], i, this);

    if (result) {
      return i;
    }
  }

  return index;
};
