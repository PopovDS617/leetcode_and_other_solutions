// типы данных
// значение this
// область видимости
// let, const, var
// “поднятие” (hoisting)

const a = 5;
let b = a;

b = 6;

console.log('a: ', a); // a:5
console.log('b: ', b); // b:6

/*-----------------------------------------*/

const obj1 = { val: 5 };
const obj2 = obj1;

obj2.val = 6;

console.log('obj1: ', obj1.val); // obj1: {val:6}
console.log('obj2: ', obj2);
6;

/*-----------------------------------------*/

const obj3 = { val: 5 };
const obj4 = obj3;

const test2 = (obj: { val: number }) => {
  obj.val = 6;
};

test2(obj3);

console.log('obj3: ', obj3); // val:6
console.log('obj4: ', obj4); //val:6
