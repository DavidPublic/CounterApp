import React, { useState, useEffect } from 'react';
import Counter from './Components/Counter';
import CounterList from './Components/CounterList';
import CounterForm from './Components/CounterForm';
import "./App.css"


const apiUrl = 'http://localhost:8080'; // Replace with your actual API URL

const App: React.FC = () => {
    const [counters, setCounters] = useState<Counter[]>([]);

    const fetchCounters = async () => {
        const response = await fetch(`${apiUrl}/counters`, { 

          method: 'GET' 
        });
        const data: Counter[] = await response.json();
        setCounters(data);
    };

    const incrementCounter = async (name: string) => {
        await fetch(`${apiUrl}/counter/${name}/increment`, { 

          method: 'POST' 
        });
        fetchCounters();
    };

    useEffect(() => {
        fetchCounters();
    }, []);

    return (
        <div className='app'>
            <h1>David's Counter App</h1>
            <CounterForm onAdd={fetchCounters} />
            <CounterList counters={counters} onIncrement={incrementCounter} />
        </div>
    );
};

export default App;
