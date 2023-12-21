import React from 'react';
import Counter from './Counter';

interface CounterListProps {
    counters: {
        id: number;
        name: string;
        value: number;
    }[];
    onIncrement: (name: string) => void;
}

const CounterList: React.FC<CounterListProps> = ({ counters, onIncrement }) => (
    <div className='counter-list'>
        {counters.map((counter) => (
            <Counter key={counter.id} counter={counter} onIncrement={onIncrement} />
        ))}
    </div>
);

export default CounterList;
