import React from 'react';
import { render, screen } from '@testing-library/react';
import Item from '../../model/item';
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
            {
                id: 1,
                description: 'this is my first todo',
                order: 1
            },
            {
                id: 2,
                description: 'this is my second todo',
                order: 2
            },
            {
                id: 3,
                description: 'this is my third todo',
                order: 3
            },
        ]

        render(<TodoList items={items} />)

        for (const item of items) {
            expect(screen.getByText(item.description)).toBeInTheDocument()
        }
    })

    test('it truncates an item description if this text is longer than 50 characters', () => {
        const items: Item[] = [
            {
                id: 1,
                description: 'this is my first todo',
                order: 1
            },
            {
                id: 2,
                description: 'this is one of the longest todo that I\'ve wrote since then. This todo description is sooo long',
                order: 2
            },
        ]

        render(<TodoList items={items} />)

        for (const item of items) {
            const truncatedDescription = item.description.length > 50 ? `${item.description.substring(0, 50)}...` : item.description
            expect(screen.getByText(truncatedDescription)).toBeInTheDocument()
        }
    })

    test('it launch an event when an item is clicked', () => {
        const items: Item[] = [
            {
                id: 1,
                description: 'this is my first todo',
                order: 1
            },
            {
                id: 2,
                description: 'this is my second todo',
                order: 2
            },
        ]

        const handleItemClick = sinon.spy()

        render(<TodoList items={items} onItemClick={handleItemClick} />)
        getItemElement(items[0].id).click()
        expect(handleItemClick.called).toBeTruthy()
    })

    test('it pass the item click on Item click', () => {
        const items: Item[] = [
            {
                id: 1,
                description: 'this is my first todo',
                order: 1
            },
            {
                id: 2,
                description: 'this is my second todo',
                order: 2
            },
        ]

        const handleItemClick = sinon.spy()

        render(<TodoList items={items} onItemClick={handleItemClick} />)
        getItemElement(items[0].id).click()
        expect(handleItemClick.calledWithExactly(items[0])).toBeTruthy()
    })
})