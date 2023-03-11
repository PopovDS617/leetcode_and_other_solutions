// returns the first deepest value
const get = (key, object) => {
  for (let i in object) {
    if (i === key) {
      return object[i];
    } else if (typeof object[i] === 'object') {
      return get(key, object[i]);
    }
  }
};

// returns key-value pair that is higher in the nesting hierarchy
const get = (key, object) => {
  if (typeof object !== 'object' || object === null) {
    return undefined;
  }

  if (key in object) {
    return object[key];
  }

  for (const objKey in object) {
    const value = get(key, object[objKey]);
    if (value !== undefined) {
      return value;
    }
  }

  return undefined;
};

// using stack to push and pop all the objects one-by-one
const get = (key, object) => {
  if (typeof object !== 'object' || object === null) {
    return undefined;
  }

  const stack = [object];

  while (stack.length > 0) {
    const currentObject = stack.pop();

    if (key in currentObject) {
      return currentObject[key];
    }

    for (const objKey in currentObject) {
      const value = currentObject[objKey];
      if (typeof value === 'object' && value !== null) {
        stack.push(value);
      }
    }
  }

  return undefined;
};
