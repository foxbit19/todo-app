import React from 'react';
import { render, screen } from '@testing-library/react';
import NewTodo from './NewTodo';
import userEvent from '@testing-library/user-event';

test('it clears the input field once save button is clicked', () => {
    render(<NewTodo />);

    const input: HTMLInputElement = screen.getByTestId('todo_description')
    userEvent.type(input, 'wonderful!')
    const button: HTMLInputElement = screen.getByTestId('todo_save')
    userEvent.click(button)
    expect(input.value).toEqual('');
})