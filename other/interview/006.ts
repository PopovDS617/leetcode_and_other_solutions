// Напишите функцию-дженерик showLength(obj) которая на вход будет принимать только объект obj имеющий поле length.
// Функция должна выводить в консоль obj.length и возвращать сам объект

type Lengthable = {
  length: number;
};

const showLength = <T extends Lengthable>(obj: T) => {
  return obj.length;
};
