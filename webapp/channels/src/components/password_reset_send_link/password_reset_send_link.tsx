// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, { PureComponent, createRef } from 'react';
import { FormattedMessage, injectIntl, WrappedComponentProps } from 'react-intl';

import type {ServerError} from '@mattermost/types/errors';

import {isEmail} from 'mattermost-redux/utils/helpers';

import BackButton from 'components/common/back_button';

// import LocalizedInput from 'components/localized_input/localized_input';

import {t} from 'utils/i18n';

interface Props extends WrappedComponentProps {
    actions: {
        sendPasswordResetEmail: (email: string) => Promise<{data: any; error: ServerError}>;
    };
}

interface State {
    error: React.ReactNode;
    updateText: React.ReactNode;
}

class PasswordResetSendLink extends PureComponent<Props, State> {
    state = {
        error: null,
        updateText: null,
    };
    resetForm = React.createRef<HTMLFormElement>();
    emailInput = React.createRef<HTMLInputElement>();

    handleSendLink = async (e: React.FormEvent) => {
        e.preventDefault();

        const email = this.emailInput.current!.value.trim().toLowerCase();
        if (!email || !isEmail(email)) {
            this.setState({
                error: (
                    <FormattedMessage
                        id='password_send.error'
                        defaultMessage='Please enter a valid email address.'
                    />
                ),
            });
            return;
        }

        // End of error checking clear error
        this.setState({error: null});

        const {data, error} = await this.props.actions.sendPasswordResetEmail(email);
        if (data) {
            this.setState({
                error: null,
                updateText: (
                    <div
                        id='passwordResetEmailSent'
                        className='reset-form alert alert-success'
                    >
                        <FormattedMessage
                            id='password_send.link'
                            defaultMessage='If the account exists, a password reset email will be sent to:'
                        />
                        <div>
                            <b>{email}</b>
                        </div>
                        <br/>
                        <FormattedMessage
                            id='password_send.checkInbox'
                            defaultMessage='Please check your inbox.'
                        />
                    </div>
                ),
            });
            if (this.resetForm.current) {
                this.resetForm.current.hidden = true;
            }
        } else if (error) {
            this.setState({
                error: error.message,
                updateText: null,
            });
        }
    };

    render() {
        const { formatMessage } = this.props.intl;
        const { error, updateText } = this.state;
        
        let error = null;
        if (error) {
          errorComponent = (
            <div className='form-group has-error'>
              <label className='control-label'>{error}</label>
            </div>
          );
        }

        let formClass = 'form-group';
        if (errorComponent) {
            formClass += ' has-error';
        }

        return (
            <div>
                <BackButton/>
                <div className='col-sm-12'>
                    <div className='signup-team__container'>
                        <FormattedMessage
                            id='password_send.title'
                            tagName='h1'
                            defaultMessage='Password Reset'
                        />
                        {updateText}
                        <form
                            onSubmit={this.handleSendLink}
                            ref={this.resetForm}
                        >
                            <p>
                                <FormattedMessage
                                    id='password_send.description'
                                    defaultMessage='To reset your password, enter the email address you used to sign up'
                                />
                            </p>
                            <div className={formClass}>
                                <input
                                    id='passwordResetEmailInput'
                                    type='email'
                                    className='form-control'
                                    name='email'
                                    placeholder={{
                                        id: 'password_send.email',
                                        defaultMessage: 'Email',
                                    }}
                                    ref={this.emailInput}
                                    spellCheck='false'
                                    autoFocus={true}
                                />
                            </div>
                            {errorComponent}
                            <button
                                id='passwordResetButton'
                                type='submit'
                                className='btn btn-primary'
                            >
                                <FormattedMessage
                                    id='password_send.reset'
                                    defaultMessage='Reset my password'
                                />
                            </button>
                        </form>
                    </div>
                </div>
            </div>
        );
    }
}

export default injectIntl(PasswordResetSendLink);
