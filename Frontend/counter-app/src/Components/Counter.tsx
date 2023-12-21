import React from 'react';
import './Counter.css';
import { Button, Container } from 'react-bootstrap';
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
                <Button variant="secondary" className="btn-dark" onClick={
                    () => onDelete(counter.name)}>Delete
                </Button>{' '}
            </Col>
        </Row>
        </Card.Body>
    </Card>
);

export default Counter;


