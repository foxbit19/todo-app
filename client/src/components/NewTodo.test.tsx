import React from 'react';
import { render, screen } from '@testing-library/react';
import NewTodo from './NewTodo';
import userEvent from '@testing-library/user-event';

describe('new todo component', () => {
    const getInputElement = (): HTMLInputElement => {
        return screen.getByTestId('todo_description')
    }

    const typeDescription = (description: string) => {
        userEvent.type(getInputElement(), description)
    }

    const clickSaveButton = () => {
        const button: HTMLInputElement = screen.getByTestId('todo_save')
        userEvent.click(button)
    }

    test('it clears the input field once save button is clicked', () => {
        render(<NewTodo />);
        typeDescription('wonderful!')
        clickSaveButton()

        expect(getInputElement().value).toEqual('');
    })

    test('it shows the items into page after save', () => {
        render(<NewTodo />);
        typeDescription('my todo item')
        clickSaveButton()

        const item = screen.getByText(/my todo item/i);
        expect(item).toBeInTheDocument();
    })
})