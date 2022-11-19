import React from 'react';
import { render, screen } from '@testing-library/react';
import ShowTodo from './ShowTodo';
import Item from '../../models/item';
import sinon from 'sinon';

describe('show todo', () => {
    let item: Item

    beforeEach(() => {
        item = {
            id: 1,
            description: 'this is my item',
            order: 2
        }
    })

    const clickButton = (testId: string) => {
        screen.getByTestId(testId).click()
    }

    test('it matches snapshot', () => {
        const view = render(<ShowTodo item={item} />);
        expect(view).toMatchSnapshot();
    });

    test('it shows the description of the provided item', () => {
        render(<ShowTodo item={item} />)

        expect(screen.getByText(item.description)).toBeInTheDocument()
    })

    test('it shows the "item details" text', () => {
        render(<ShowTodo item={item} />)

        expect(screen.getByText('Item details')).toBeInTheDocument()
    })

    test('it launch an event when update button is clicked', () => {
        const updateHandler = sinon.spy();

        render(<ShowTodo item={item} onUpdateClick={updateHandler} />)
        clickButton('update_button')
        expect(updateHandler.called).toBeTruthy()
    })

    test('it launch an event when delete button is clicked', () => {
        const deleteHandler = sinon.spy();

        render(<ShowTodo item={item} onDeleteClick={deleteHandler} />)
        clickButton('delete_button')
        expect(deleteHandler.called).toBeTruthy()
    })

    test('it launch update event providing the item as argument', () => {
        const updateHandler = sinon.spy();

        render(<ShowTodo item={item} onUpdateClick={updateHandler} />)
        clickButton('update_button')
        expect(updateHandler.calledWith(item)).toBeTruthy()
    })

    test('it launch delete event providing the item as argument', () => {
        const deleteHandler = sinon.spy();

        render(<ShowTodo item={item} onDeleteClick={deleteHandler} />)
        clickButton('delete_button')
        expect(deleteHandler.calledWith(item)).toBeTruthy()
    })
})