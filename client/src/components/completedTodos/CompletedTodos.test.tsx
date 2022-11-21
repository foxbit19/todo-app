import React from 'react';
import { render, screen } from '@testing-library/react';
import sinon from 'sinon';
import userEvent from '@testing-library/user-event';
import CompletedTodos from './CompletedTodos';

describe('completed todos dialog', () => {
    test('it matches snapshot', () => {
        const view = render(<CompletedTodos open={true} items={[]} />);
        expect(view).toMatchSnapshot();
    });

    test('it calls onClose when close button is clicked', () => {
        const handleClick = sinon.spy()
        render(<CompletedTodos open={true} items={[]} onClose={handleClick} />);

        userEvent.click(screen.getByTestId('close_dialog_button'))

        expect(handleClick.called).toBeTruthy();
    });
})