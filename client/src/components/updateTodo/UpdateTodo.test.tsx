import React from 'react';
import { render, screen } from '@testing-library/react';
import Item from '../../models/item';
import sinon from 'sinon';
import UpdateTodo from './UpdateTodo';
import userEvent from '@testing-library/user-event';

describe('update todo', () => {
    let item: Item

    beforeEach(() => {
        item = new Item(1, 'this is my item', 2)
    })

    const getInputElement = (): HTMLInputElement => {
        return screen.getByTestId('description')
    }

    const updateButtonClick = () => {
        const buttonElement: HTMLButtonElement = screen.getByTestId('update_button')
        buttonElement.click()
    }

    test('it matches snapshot', () => {
        const view = render(<UpdateTodo open={true} item={item} />);
        expect(view).toMatchSnapshot();
    });

    test('it shows the "item update" text', () => {
        render(<UpdateTodo open={true} item={item} />)

        expect(screen.getByText('Item update')).toBeInTheDocument()
    })

    test('it shows the description of the provided item into an input box', () => {
        render(<UpdateTodo open={true} item={item} />)

        expect(getInputElement().value).toEqual(item.description)
    })

    test('it launch an event when the update button is pressed', () => {
        const updateHandler = sinon.spy();

        render(<UpdateTodo open={true} item={item} onUpdateClick={updateHandler} />)
        updateButtonClick()

        expect(updateHandler.called).toBeTruthy()
    })

    test('it launch an event with the modified item when the update button is pressed', () => {
        const updateHandler = sinon.spy();

        render(<UpdateTodo open={true} item={item} onUpdateClick={updateHandler} />)

        const newDescription = 'remember to test everything'

        userEvent.clear(getInputElement())
        userEvent.type(getInputElement(), newDescription)
        updateButtonClick()

        expect(updateHandler.calledWith({
            id: item.id,
            description: newDescription,
            order: item.order
        })).toBeTruthy()
    })
})