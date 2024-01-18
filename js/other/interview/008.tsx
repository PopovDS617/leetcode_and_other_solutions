// Jsx
// Хуки
// Жизненный цикл компонента на хуках
// Контекст

import * as React from 'react';

const LifeCycle = () => {
  React.useEffect(() => {
    console.log('1');
    return () => {
      console.log('2');
    };
  });

  console.log('3');

  return <div>LifeCycle</div>;
};
export default LifeCycle;

// 3 2 1
