import '@testing-library/jest-dom';
import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import CounterList from './CounterList';

describe('CounterList', () => {
  it('renders a list of counters', () => {
    const mockCounters = [
      { id: 1, name: 'Counter 1', value: 3 },
      { id: 2, name: 'Counter 2', value: 5 }
    ];

    render(<CounterList counters={mockCounters} onIncrement={() => {}} onDelete={() => {}} />);

    expect(screen.getByText('Counter 1')).toBeInTheDocument();
    expect(screen.getByText('3')).toBeInTheDocument();
    expect(screen.getByText('Counter 2')).toBeInTheDocument();
    expect(screen.getByText('5')).toBeInTheDocument();
  });
});
