// Event loop (micro-macro)
// Promise (и его методы)
// async/await

console.log(1);

setTimeout(() => console.log(2), 0);

new Promise((resolve) => {
  console.log(3);
  resolve(undefined);
}).then((data) => {
  console.log(4);
});

console.log(5);

// 1 3 5 4 // 2
