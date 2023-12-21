import React, { useState, useEffect } from 'react';
import Counter from './Components/Counter';
import CounterList from './Components/CounterList';
import CounterForm from './Components/CounterForm';
import "./App.css"
import 'bootstrap/dist/css/bootstrap.min.css';
import { Container, Navbar } from 'react-bootstrap';

interface Counter {
    id: number;
    name: string;
    value: number;
}

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

    const deleteCounter = async (name: string) => {
        await fetch(`${apiUrl}/counter/${name}`, { 
            method: 'DELETE' 
        });
        setCounters(counters.filter(counter => counter.name !== name));
    };

    useEffect(() => {
        fetchCounters();
    }, []);

    return (
        <>
            <Navbar expand="lg" className="bg-dark">
                <Container>
                    <Navbar.Brand href="#home" className='text-light' >Davids counter-app
                    </Navbar.Brand>
                </Container>
            </Navbar>
            <div className='app'>
                <CounterForm onAdd={fetchCounters} />
                <CounterList counters={counters} onIncrement={incrementCounter} onDelete={deleteCounter}/>
            </div>
        </>
    );
};

export default App;

