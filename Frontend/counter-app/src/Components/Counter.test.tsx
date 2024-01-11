import '@testing-library/jest-dom';
import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import Counter from './Counter';

describe('Counter', () => {
  it('renders the counter with correct values', () => {
    const mockCounter = { id: 1, name: 'Test Counter', value: 5 };
    render(<Counter counter={mockCounter} onIncrement={() => {}} onDelete={() => {}} />);

    expect(screen.getByText('Test Counter')).toBeInTheDocument();
    expect(screen.getByText('5')).toBeInTheDocument();
  });

  it('calls onIncrement when the increment button is clicked', () => {
    const mockCounter = { id: 1, name: 'Test Counter', value: 5 };
    const mockIncrement = vi.fn();
    render(<Counter counter={mockCounter} onIncrement={mockIncrement} onDelete={() => {}} />);

    const incrementButton = screen.getByText('Increment');
    fireEvent.click(incrementButton);

    expect(mockIncrement).toHaveBeenCalledWith('Test Counter');
  });

  it('calls onDelete when the close button is clicked', () => {
    const mockCounter = { id: 1, name: 'Test Counter', value: 5 };
    const mockDelete = vi.fn();
    render(<Counter counter={mockCounter} onIncrement={() => {}} onDelete={mockDelete} />);

    const deleteButton = screen.getByRole('button', { name: /close/i });
    fireEvent.click(deleteButton);

    expect(mockDelete).toHaveBeenCalledWith('Test Counter');
  });
});
