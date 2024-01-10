import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import CounterList from './CounterList';

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

// Mock functions for increment and delete actions
const mockIncrement = jest.fn();
const mockDelete = jest.fn();

describe('CounterList Component', () => {
  test('renders a list of counters', () => {
    const counters = [
      { id: 1, name: 'Counter 1', value: 3 },
      { id: 2, name: 'Counter 2', value: 5 }
    ];
    render(<CounterList counters={counters} onIncrement={mockIncrement} onDelete={mockDelete} />);

    expect(screen.getByText('Counter 1')).toBeInTheDocument();
    expect(screen.getByText('Counter 2')).toBeInTheDocument();
  });

  // Additional tests can be added here...
});
