import React from 'react';
import Counter from './Counter';
import ListGroup from 'react-bootstrap/ListGroup';

interface Counter {
    id: number;
    name: string;
    value: number;
}

interface CounterListProps {
    counters: Counter[];
    onIncrement: (name: string) => void;
    onDelete: (name: string) => void;
}

const CounterList: React.FC<CounterListProps> = ({ counters, onIncrement, onDelete }) => (
    <ListGroup>
        <ListGroup.Item variant="dark">
            {counters.map((counter) => (
            <Counter key={counter.id} counter={counter} onDelete={onDelete} onIncrement={onIncrement} />
            ))}
             
        </ListGroup.Item>
    </ListGroup>
);

export default CounterList;
