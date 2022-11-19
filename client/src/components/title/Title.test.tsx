import React from 'react';
import { render, screen } from '@testing-library/react';
import Title from './Title';

describe('title component', () => {
    test('it contains the title', () => {
        render(<Title />);
        expect(screen.getByText(/ToDo App/i)).toBeInTheDocument();
    })

    test('it shows a subtitle', () => {
        render(<Title />);
        expect(screen.getByText(/The perfect solution for yours day to day ToDos/i)).toBeInTheDocument();
    })
})