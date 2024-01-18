const shuffleArray = (array) => {
  for (let i = 0; i < array.length; i++) {
    const randomIndex = Math.floor(Math.random() * array.length);

    const elementA = array[i];
    const elementB = array[randomIndex];

    array[i] = elementB;
    array[randomIndex] = elementA;
  }

  return array;
};
