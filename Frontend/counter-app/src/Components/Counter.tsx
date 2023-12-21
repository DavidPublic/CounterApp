import React from 'react';
import './Counter.css';
import { Button, CloseButton } from 'react-bootstrap';
import Card from 'react-bootstrap/Card';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

interface CounterProps {
    counter: {
        id: number;
        name: string;
        value: number;
    };
    onIncrement: (name: string) => void;
    onDelete: (name: string) => void; // Add this line
}


    const Counter: React.FC<CounterProps> = ({ counter, onIncrement, onDelete }) => (
    <Card>
        <Card.Body>
        <Row>
            <Col>
                <Button variant="secondary" className="btn-dark" onClick={
                    () => onIncrement(counter.name)}>Increment
                </Button>{' '}
            </Col>
            <Col >
                <h5>{counter.name}</h5>
            </Col>
            <Col>
                <h3>{counter.value}</h3>
            </Col>
            <Col>
                <CloseButton variant="secondary" className="topright" onClick={
                    () => onDelete(counter.name)}>
                </CloseButton>{' '}
            </Col>
        </Row>
        </Card.Body>
    </Card>
);

export default Counter;


