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

    test('it matches snapshot', () => {
        const view = render(<NewTodo open={true} />);
        expect(view).toMatchSnapshot();
    });

})