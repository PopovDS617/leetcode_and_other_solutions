// Object:
// {
//   a: {
//     b: {
//       d: {
//         c: 1,
//       }
//     },
//   },
// }

// Path: "abc"

const isPathExist = (obj, path) => {
  let result = { ...obj };

  for (let letter of path) {
    if (result[letter]) {
      result = result[letter];
    } else {
      return false;
    }
  }

  return true;
};
