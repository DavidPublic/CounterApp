import React, { useState } from 'react';
import './CounterForm.css';
import Form from 'react-bootstrap/Form';
import { Button } from 'react-bootstrap';
import { InputGroup } from 'react-bootstrap';

interface CounterFormProps {
    onAdd: () => void;
}

const apiUrl = 'http://localhost:8080'; // Replace with your actual API URL

const CounterForm: React.FC<CounterFormProps> = ({ onAdd }) => {
    const [name, setName] = useState<string>('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!name) return;
        // Assuming apiUrl is defined globally or passed as a prop
        await fetch(`${apiUrl}/counter`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name }),
        });
        onAdd();
        setName('');
    };

    return (     
        <Form onSubmit={handleSubmit}>
            <InputGroup className="mb-3">
            <Button type='submit' variant="dark" id="button-addon1">
                Create Counter
        </Button>
            <Form.Control 
                type="text" 
                placeholder="Name a new counter" 
                value={name}  
                onChange={(e) => setName(e.target.value)} 
            />
            </InputGroup>
        </Form>
    );
};

export default CounterForm;
