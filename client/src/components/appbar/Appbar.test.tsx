import React from 'react';
import { render, screen } from '@testing-library/react';
import Item from '../../models/item';
import sinon from 'sinon';
import Appbar from './Appbar';
import userEvent from '@testing-library/user-event';

describe('app bar list', () => {
    test('it matches snapshot', () => {
        const view = render(<Appbar onShowCompletedClick={() => { }} />);
        expect(view).toMatchSnapshot();
    });

    test('it calls the callback when show completed todo is clicked', () => {
        const handleClick = sinon.spy()
        render(<Appbar onShowCompletedClick={handleClick} />);

        userEvent.click(screen.getByTestId('show_completed_button'))

        expect(handleClick.called).toBeTruthy();
    });
})