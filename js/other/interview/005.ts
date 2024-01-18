// Необходимо написать проверку необработанного кейса
// (должен отображать ошибку на этапе компиляции)

enum Color {
  Red,
  Green,
  Blue,
}

const exhaustiveMatchingGuard = (_: never): never => {
  throw new Error('ye');
};

function getColorName(c: Color): string {
  switch (c) {
    case Color.Red:
      return 'red';
    case Color.Green:
      return 'green';
    case Color.Blue:
      return 'blue'; // all
    default:
      return exhaustiveMatchingGuard(c);
  }
}
