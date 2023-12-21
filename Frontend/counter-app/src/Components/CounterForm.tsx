import React, { useState } from 'react';
import './CounterForm.css';

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
        <form onSubmit={handleSubmit} className="counter-form">
            <input
                type="text"
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="Counter name"
            />
            <button type="submit">Create Counter</button>
        </form>
    );
};

export default CounterForm;
