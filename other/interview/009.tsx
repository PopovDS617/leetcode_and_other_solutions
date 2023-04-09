import * as React, {useState} from 'react';

const Timer = () => {
const [count,setCount]=useState(0)


  return (
    <div>
      <h2>seconds: 0</h2>
      <button>Toggle</button>
      <button>Reset</button>
    </div>
  );
};

export default Timer;
