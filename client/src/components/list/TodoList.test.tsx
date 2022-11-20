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

    test('it truncates an item description if this text is longer than 50 characters', () => {
        const items: Item[] = [
            new Item(1, 'this is my first todo', 1),
            new Item(2, 'this is one of the longest todo that I\'ve wrote since then. This todo description is sooo long', 2),
        ]

        render(<TodoList items={items} />)

        for (const item of items) {
            const truncatedDescription = item.description.length > 50 ? `${item.description.substring(0, 50)}...` : item.description
            expect(screen.getByText(truncatedDescription)).toBeInTheDocument()
        }
    })

    test('it launch an event when an item is clicked', () => {
        const items: Item[] = [
            new Item(1, 'this is my first todo', 1),
            new Item(2, 'this is my second todo', 2),
        ]

        const handleItemClick = sinon.spy()

        render(<TodoList items={items} onItemClick={handleItemClick} />)
        screen.getByText(items[0].description).click()
        expect(handleItemClick.called).toBeTruthy()
    })

    test('it pass the item clicked on Item click', () => {
        const items: Item[] = [
            new Item(1, 'this is my first todo', 1),
            new Item(2, 'this is my second todo', 2),
        ]

        const handleItemClick = sinon.spy()

        render(<TodoList items={items} onItemClick={handleItemClick} />)
        screen.getByText(items[0].description).click()
        expect(handleItemClick.calledWithExactly(items[0])).toBeTruthy()
    })
})