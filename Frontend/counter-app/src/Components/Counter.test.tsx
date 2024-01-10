import React from 'react';
import { render, fireEvent, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import Counter from './Counter';

interface CounterProps {
    counter: {
        id: number;
        name: string;
        value: number;
    };
    onIncrement: (name: string) => void;
    onDelete: (name: string) => void;
}

const mockIncrement = jest.fn();
const mockDelete = jest.fn();

describe('Counter Component', () => {
  test('displays the counter value', () => {
    const counter = { id: 1, name: 'Test Counter', value: 5 };
    render(<Counter counter={counter} onIncrement={mockIncrement} onDelete={mockDelete} />);

    expect(screen.getByText('Test Counter')).toBeInTheDocument();
    expect(screen.getByText(/value: 5/i)).toBeInTheDocument();
  });

  test('calls onIncrement when increment button is clicked', () => {
    render(<Counter counter={{ id: 1, name: 'Test Counter', value: 5 }} onIncrement={mockIncrement} onDelete={mockDelete} />);

    fireEvent.click(screen.getByText(/increment/i));
    expect(mockIncrement).toHaveBeenCalledWith('Test Counter');
  });
});
