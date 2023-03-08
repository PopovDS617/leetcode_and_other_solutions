const sortAscendingByDate = (array) => {
  if (array.length === 0) {
    return [];
  }

  const getTimeInMs = (str) => {
    const date = new Date(str);
    const ms = date.getTime();
    return ms;
  };

  return array.sort((a, b) => {
    console.log(getTimeInMs(a.date));
    if (getTimeInMs(a.date) > getTimeInMs(b.date)) {
      return 1;
    }
    if (getTimeInMs(a.date) < getTimeInMs(b.date)) {
      return -1;
    }
    return 0;
  });
};
