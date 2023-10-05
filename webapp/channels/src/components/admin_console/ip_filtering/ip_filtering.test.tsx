import React from 'react';
import { render, fireEvent, waitFor, screen } from '@testing-library/react';
import IPFiltering from './index';
import { Client4 } from 'mattermost-redux/client';
import { AllowedIPRange } from '@mattermost/types/config';
import { IntlProvider } from 'react-intl';
import { Provider } from 'react-redux';
import testConfigureStore from 'tests/test_store';
import {BrowserRouter as Router} from 'react-router-dom';

jest.mock('mattermost-redux/client');

describe('IPFiltering', () => {
    const ipFilters = [
        {
            CIDRBlock: '10.0.0.0/8',
            Description: 'Test IP Filter',
            Enabled: true,
        },
    ] as AllowedIPRange[];

    const intlProviderProps = {
        defaultLocale: 'en',
        locale: 'en',
    };
    const currentIP = '10.0.0.1';
    const applyIPFiltersMock = jest.fn(() => Promise.resolve(ipFilters));
    const getIPFiltersMock = jest.fn(() => Promise.resolve(ipFilters));
    const getCurrentIPMock = jest.fn(() => Promise.resolve({ IP: currentIP }));

    beforeEach(() => {
        Client4.applyIPFilters = applyIPFiltersMock;
        Client4.getIPFilters = getIPFiltersMock;
        Client4.getCurrentIP = getCurrentIPMock;
    });

    afterEach(() => {
        jest.clearAllMocks();
    });

    const mockedStore = testConfigureStore({
        entities: {
            
        },
        views: {
            admin: {
                navigationBlock: {
                    blocked: false,
                },
            }
        }
    });

    const wrapWithIntlProviderAndStore = (component: JSX.Element) => (
        <Router>
            <IntlProvider {...intlProviderProps}>
                <Provider store={mockedStore} >
                    {component}
                </Provider>
            </IntlProvider>
        </Router>
    );


    test('renders the IP Filtering page', async () => {
        const {getByText} = render(wrapWithIntlProviderAndStore(<IPFiltering />));

        expect(getByText('IP Filtering')).toBeInTheDocument();
        expect(getByText('Enable IP Filtering')).toBeInTheDocument();


        await waitFor(() => {
            expect(getByText('Add Filter')).toBeInTheDocument();
            expect(getByText('Test IP Filter')).toBeInTheDocument();
            expect(getByText('10.0.0.0/8')).toBeInTheDocument();
        });

        expect(getByText('Save')).toBeInTheDocument();
        expect(getByText('Cancel')).toBeInTheDocument();
    });

    test('disables IP Filtering when the toggle is turned off', async () => {
        render(wrapWithIntlProviderAndStore(<IPFiltering />));

        await waitFor(() => {
            expect(screen.getByTestId('filterToggle-button')).toBeInTheDocument();
            expect(screen.getByRole('button', { pressed: true })).toBeInTheDocument();
        });

        fireEvent.click(screen.getByTestId('filterToggle-button'));

        await waitFor(() => {
            expect(screen.getByRole('button', { pressed: false })).toBeInTheDocument();
        });
    });

    test('adds a new IP filter when the "Add IP Filter" button is clicked', async () => {
        const { getByLabelText, getByText } = render(wrapWithIntlProviderAndStore(<IPFiltering />));


        await waitFor(() => {
            expect(getByText('Add Filter')).toBeInTheDocument();
        });

        fireEvent.click(getByText('Add Filter'));

        const descriptionInput = getByLabelText('Enter a name for this rule');
        const cidrInput = getByLabelText('Enter IP Range');
        const saveButton = screen.getByTestId('save-add-edit-button');

        fireEvent.change(cidrInput, { target: { value: '192.168.0.0/16' } });
        fireEvent.change(descriptionInput, { target: { value: 'Test IP Filter 2' } });
        fireEvent.click(saveButton);

        await waitFor(() => {
            expect(getByText('Test IP Filter 2')).toBeInTheDocument();
            expect(getByText('192.168.0.0/16')).toBeInTheDocument();
        });
    });

    test('edits an existing IP filter when the "Edit" button is clicked', async () => {
        const { getByLabelText, getByText, queryByText } = render(wrapWithIntlProviderAndStore(<IPFiltering />));

        await waitFor(() => {
            expect(getByText('Test IP Filter')).toBeInTheDocument();
        });

        fireEvent.mouseEnter(screen.getByText('Test IP Filter'));
        fireEvent.click(screen.getByRole('button', {
            name: /Edit/i,
        }));

        const descriptionInput = getByLabelText('Enter a name for this rule');
        const cidrInput = getByLabelText('Enter IP Range');
        const saveButton = screen.getByTestId('save-add-edit-button');

        fireEvent.change(cidrInput, { target: { value: '192.168.0.0/16' } });
        fireEvent.change(descriptionInput, { target: { value: 'zzzzzfilter' } });
        fireEvent.click(saveButton);

        await waitFor(() => {
            expect(getByText('zzzzzfilter')).toBeInTheDocument();
            expect(getByText('192.168.0.0/16')).toBeInTheDocument();
            // ensure that the old description is gone, because we've now changed it
            expect(queryByText('Test IP Filter')).toBeNull();
        });
    });

    test('deletes an existing IP filter when the "Delete" button is clicked', async () => {
        const { getByText, queryByText } = render(wrapWithIntlProviderAndStore(<IPFiltering />));

        await waitFor(() => {
            expect(getByText('Test IP Filter')).toBeInTheDocument();
        });


        fireEvent.mouseEnter(screen.getByText('Test IP Filter'));
        fireEvent.click(screen.getByRole('button', {
            name: /Delete/i
        }));

        const confirmButton = getByText('Delete filter');

        fireEvent.click(confirmButton);

        await waitFor(() => {
            expect(queryByText('Test IP Filter')).not.toBeInTheDocument();
        });
    });

    test('saves changes when the "Save" button is clicked', async () => {
        const { getByText, queryByText, getByLabelText } = render(wrapWithIntlProviderAndStore(<IPFiltering />));

        await waitFor(() => {
            expect(screen.getByTestId('filterToggle-button')).toBeInTheDocument();
            expect(screen.getByRole('button', { pressed: true })).toBeInTheDocument();
        });

        fireEvent.click(screen.getByTestId('filterToggle-button'));

        await waitFor(() => {
            expect(screen.getByRole('button', { pressed: false })).toBeInTheDocument();
        });

        await waitFor(() => {
            expect(queryByText('Test IP Filter')).not.toBeInTheDocument();
        });

        fireEvent.click(getByText('Save'));
        fireEvent.click(screen.getByTestId('save-confirmation-button'));

        await waitFor(() => {
            expect(applyIPFiltersMock).toHaveBeenCalledTimes(1);
            expect(screen.getByRole('button', { pressed: false })).toBeInTheDocument();
        });
    });
});