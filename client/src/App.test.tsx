import { render, screen } from '@testing-library/react';
import sinon from 'sinon';
import App from "./App"

describe('main app component', () => {
    test('it contains the new todo button', () => {
        render(<App />);
        expect(screen.getByTestId('new_button')).toBeInTheDocument();
    })
})