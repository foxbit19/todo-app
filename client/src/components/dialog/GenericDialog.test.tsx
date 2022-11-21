import React from 'react';
import { render, screen } from '@testing-library/react';
import sinon from 'sinon';
import userEvent from '@testing-library/user-event';
import Item from '../../models/item';
import GenericDialog from './GenericDialog';

describe('generic dialog', () => {
    let item: Item

    beforeEach(() => {
        item = new Item(1, 'test', 1)
    })

    test('it matches snapshot', () => {
        const view = render(<GenericDialog open={true} title='' text='' />);
        expect(view).toMatchSnapshot();
    });

    test('it calls onClose when close button is clicked', () => {
        const handleClick = sinon.spy()
        render(<GenericDialog open={true} title='' text='' onClose={handleClick} />);

        userEvent.click(screen.getByTestId('close_dialog_button'))

        expect(handleClick.called).toBeTruthy();
    });

    test('it show title and text', () => {
        render(<GenericDialog open={true} title='My title' text='My text' />);

        expect(screen.getByText('My title')).toBeTruthy();
        expect(screen.getByText('My text')).toBeTruthy();
    });
})