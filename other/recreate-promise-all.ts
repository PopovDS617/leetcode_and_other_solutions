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

    const p1 = new Promise((res) => {
      setTimeout(() => {
        res(1);
      }, 2500);
    });
    const p2 = new Promise((res) => res(11));

    const fetchAll = async (arr) => {
      const result = await customPromiseAll(arr);

      setState(result);
    };
    fetchAll([p1, p2]);
