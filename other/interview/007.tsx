// Что такое реакт
// SOLID в рекат
// Virtual Dom (Change tracker) React Reconciliation
// Условия для рендера компонента

import { useEffect } from 'react';

const Child1 = () => {
  console.log('Child1');
  return <div></div>;
};

const Child2 = () => {
  useEffect(() => {
    console.log('Child2');
  }, []);
  return <div></div>;
};

const Container = () => {
  useEffect(() => {
    console.log('Container');
  }, []);
  return (
    <div>
      <Child1 />
      <Child2 />
    </div>
  );
};

export default Container;

// 1 2 Container
