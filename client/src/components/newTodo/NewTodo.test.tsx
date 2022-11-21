import React from 'react';
import { render, screen } from '@testing-library/react';
import NewTodo from './NewTodo';
import userEvent from '@testing-library/user-event';
import sinon from 'sinon';

describe('new todo component', () => {
    test('it matches snapshot', () => {
        const view = render(<NewTodo open={true} />);
        expect(view).toMatchSnapshot();
    });

    test('it calls onClose when close button is clicked', () => {
        const handleClick = sinon.spy()
        render(<NewTodo open={true} onClose={handleClick} />);

        userEvent.click(screen.getByTestId('close_dialog_button'))

        expect(handleClick.called).toBeTruthy();
    });

    test('it does not call onSaveClick when save button is clicked and no input is provided', () => {
        const handleClick = sinon.spy()
        render(<NewTodo open={true} onSaveClick={handleClick} />);

        userEvent.click(screen.getByTestId('save_button'))

        expect(handleClick.called).toBeFalsy();
    });
})