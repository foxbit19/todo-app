import React from 'react';
import { render, screen } from '@testing-library/react';
import sinon from 'sinon';
import userEvent from '@testing-library/user-event';
import DeleteTodo from './DeleteTodo';
import Item from '../../models/item';

describe('delete todo dialog', () => {
    let item: Item

    beforeEach(() => {
        item = new Item(1, 'test', 1)
    })

    test('it matches snapshot', () => {
        const view = render(<DeleteTodo open={true} item={item} />);
        expect(view).toMatchSnapshot();
    });

    test('it calls onClose when close button is clicked', () => {
        const handleClick = sinon.spy()
        render(<DeleteTodo open={true} item={item} onClose={handleClick} />);

        userEvent.click(screen.getByTestId('close_dialog_button'))

        expect(handleClick.called).toBeTruthy();
    });

    test('it show the description of the item to delete', () => {
        render(<DeleteTodo open={true} item={item} />);

        expect(screen.getByText(item.description)).toBeTruthy();
    });
})