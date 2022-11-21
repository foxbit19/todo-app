import React from 'react';
import { render } from '@testing-library/react';
import Footer from './Footer';

describe('footer', () => {
    test('it matches snapshot', () => {
        const view = render(<Footer />);
        expect(view).toMatchSnapshot();
    });
})