import React from 'react';
import './Counter.css';

interface Counter {
    id: number;
    name: string;
    value: number;
}

const Counter: React.FC<{ counter: Counter; onIncrement: (name: string) => void }> = ({ counter, onIncrement }) => (
    <div className='counter'>
        {counter.name}: {counter.value}
        <button onClick={() => onIncrement(counter.name)}>Increment</button>
    </div>
);

export default Counter;