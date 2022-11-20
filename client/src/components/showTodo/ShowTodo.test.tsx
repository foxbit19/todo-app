import React from 'react';
import { render, screen } from '@testing-library/react';
import ShowTodo from './ShowTodo';
import Item from '../../models/item';
import sinon from 'sinon';

describe('show todo', () => {
    let item: Item

    beforeEach(() => {
        item = new Item(1, 'this is my item', 2)
    })

    const clickButton = (testId: string) => {
        screen.getByTestId(testId).click()
    }

    test('it matches snapshot', () => {
        const view = render(<ShowTodo open={true} item={item} />);
        expect(view).toMatchSnapshot();
    });

    test('it shows the description of the provided item', () => {
        render(<ShowTodo open={true} item={item} />)

        expect(screen.getByText(item.description)).toBeInTheDocument()
    })

    test('it shows the "item details" text', () => {
        render(<ShowTodo open={true} item={item} />)

        expect(screen.getByText('Item details')).toBeInTheDocument()
    })

    test('it launch an event when update button is clicked', () => {
        const updateHandler = sinon.spy();

        render(<ShowTodo open={true} item={item} onUpdateClick={updateHandler} />)
        clickButton('update_button')
        expect(updateHandler.called).toBeTruthy()
    })

    test('it launch update event providing the item as argument', () => {
        const updateHandler = sinon.spy();

        render(<ShowTodo open={true} item={item} onUpdateClick={updateHandler} />)
        clickButton('update_button')
        expect(updateHandler.calledWith(item)).toBeTruthy()
    })
})