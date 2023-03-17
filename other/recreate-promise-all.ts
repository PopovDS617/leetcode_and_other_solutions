  const customPromiseAll = (array) => {
    const results = [];
    let count = 0;
    return new Promise((resolve, rej) => {
      for (let i = 0; i < array.length; i++) {
        array[i]
          .then((res) => {
            results[i] = res;
            count++;

            if (count === array.length) {
              resolve(results);
            }
          })
          .catch((err) => rej(err));
      }
    });
  };
