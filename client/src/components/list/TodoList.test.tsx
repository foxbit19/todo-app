import React from 'react';
import { render, screen } from '@testing-library/react';
import Item from '../../models/item';
import sinon from 'sinon';
import TodoList from './TodoList';

describe('todo list', () => {
    const getItemElement = (itemId: number): HTMLInputElement => {
        return screen.getByTestId(`todo_${itemId}`)
    }

    test('it matches snapshot', () => {
        const view = render(<TodoList items={[]} />);
        expect(view).toMatchSnapshot();
    });

    test('it shows a message when the provided list of elements is empty', () => {
        render(<TodoList items={[]} />)

        expect(screen.getByText(/There are no todos to show/i)).toBeInTheDocument()
    })

    test('it shows the provided list of elements using description field', () => {
        const items: Item[] = [
            new Item(1, 'this is my first todo', 1),
            new Item(2, 'this is my second todo', 2),
            new Item(3, 'this is my third todo', 3)
        ]

        render(<TodoList items={items} />)

        for (const item of items) {
            expect(screen.getByText(item.description)).toBeInTheDocument()
        }
    })
})