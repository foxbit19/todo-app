import { render, screen } from '@testing-library/react';
import sinon from 'sinon';
import App from "./App"

describe('main app component', () => {
    test('it contains the title component', () => {
        render(<App />);
        expect(screen.getByTestId('title')).toBeInTheDocument();
    })

    test('it contains the items list component', () => {
        render(<App />);
        expect(screen.getByTestId('list')).toBeInTheDocument();
    })

    test('it contains the new todo button', () => {
        render(<App />);
        expect(screen.getByTestId('new_button')).toBeInTheDocument();
    })
})