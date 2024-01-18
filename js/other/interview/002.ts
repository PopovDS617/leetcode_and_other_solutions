// чистые функции
// функции высшего порядка
// стрелочные функции
// замыкания
// Работа с массивами

type Foo<T> = (start: number, end: number) => (value: T) => boolean;

const arr = [1, 2, 3, 4, 5, 6];

const foo: Foo<number> = (start, end) => {
  return (value) => {
    if (value > start && value < end) {
      return true;
    }
    return false;
  };
};

const getBetween = foo(2, 5);

const result = arr.filter(getBetween);

console.log('result: ', result); // [3, 4]

const foo2 = <T extends number, U extends string>(
  start: U,
  end: U
): ((value: T) => boolean) => {
  return (value) => {
    if (value > start.length && value < end.length) {
      return true;
    }
    return false;
  };
};

const getBetween2 = foo2<number, string>('hi', 'hello');
