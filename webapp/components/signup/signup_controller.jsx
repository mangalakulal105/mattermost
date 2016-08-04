// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import React from 'react';

import FormError from 'components/form_error.jsx';
import LoadingScreen from 'components/loading_screen.jsx';

import UserStore from 'stores/user_store.jsx';
import BrowserStore from 'stores/browser_store.jsx';

import * as AsyncClient from 'utils/async_client.jsx';
import Client from 'client/web_client.jsx';
import * as GlobalActions from 'actions/global_actions.jsx';

import logoImage from 'images/logo.png';
import ErrorBar from 'components/error_bar.jsx';

import {FormattedMessage} from 'react-intl';
import {browserHistory, Link} from 'react-router/es6';

export default class SignupController extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            emailEnabled: global.window.mm_config.EnableSignUpWithEmail === 'true',
            gitlabEnabled: global.window.mm_config.EnableSignUpWithGitLab === 'true',
            googleEnabled: global.window.mm_config.EnableSignUpWithGoogle === 'true',
            office365Enabled: global.window.mm_config.EnableSignUpWithOffice365 === 'true',
            ldapEnabled: global.window.mm_license.IsLicensed === 'true' && global.window.mm_config.EnableLdap === 'true',
            samlEnabled: global.window.mm_license.IsLicensed === 'true' && global.window.mm_config.EnableSaml === 'true',
            teamName: '',
            serverError: '',
            noOpenServerError: false,
            loading: true
        };
    }

    componentWillMount() {
        if (window.location.query) {
            const hash = window.location.query.h;
            const data = window.location.query.d;
            const inviteId = window.location.query.id;

            if ((inviteId && inviteId.length > 0) || (hash && hash.length > 0)) {
                if (UserStore.getCurrentUser()) {
                    Client.addUserToTeamFromInvite(
                        data,
                        hash,
                        inviteId,
                        (team) => {
                            GlobalActions.emitInitialLoad(
                                () => {
                                    browserHistory.push('/' + team.name);
                                }
                            );
                        },
                        (err) => {
                            this.setState({
                                serverError: err.message
                            });
                        }
                    );
                } else if (hash) {
                    const parsedData = JSON.parse(data);
                    this.setState({
                        teamName: parsedData.name,
                        usedBefore: BrowserStore.getGlobalItem(hash),
                        loading: false
                    });
                } else {
                    Client.getInviteInfo(
                        inviteId,
                        (inviteData) => {
                            if (!inviteData) {
                                return;
                            }

                            this.setState({
                                serverError: '',
                                teamName: inviteData.name,
                                loading: false
                            });
                        },
                        () => {
                            this.setState({
                                noOpenServerError: true,
                                loading: false,
                                serverError:
                                    <FormattedMessage
                                        id='signup_user_completed.invalid_invite'
                                        defaultMessage='The invite link was invalid.  Please speak with your Administrator to receive an invitation.'
                                    />
                            });
                        }
                    );
                }
            } else if (UserStore.getCurrentUser()) {
                browserHistory.push('/select_team');
            } else if (global.window.mm_config.EnableOpenServer !== 'true' && !UserStore.getNoAccounts()) {
                this.setState({
                    noOpenServerError: true,
                    loading: false,
                    serverError:
                        <FormattedMessage
                            id='signup_user_completed.no_open_server'
                            defaultMessage='This server does not allow open signups.  Please speak with your Administrator to receive an invitation.'
                        />
                });
            }
        } else {
            this.setState({
                loading: false
            });
        }
    }

    componentDidMount() {
        AsyncClient.checkVersion();
    }

    render() {
        if (this.state.loading) {
            return (<LoadingScreen/>);
        }

        if (this.state.usedBefore) {
            return (
                <div>
                    <FormattedMessage
                        id='signup_user_completed.expired'
                        defaultMessage="You've already completed the signup process for this invitation or this invitation has expired."
                    />
                </div>
            );
        }

        let signupControls = [];

        if (this.state.emailEnabled) {
            signupControls.push(
                <Link
                    className='btn btn-custom-login btn--full email'
                    key='email'
                    to={'/signup_email' + window.location.search}
                >

                    <span className='icon fa fa-envelope'/>
                    <span>
                        <FormattedMessage
                            id='signup.email'
                            defaultMessage='Email and Password'
                        />
                    </span>
                </Link>
            );
        }

        if (this.state.gitlabEnabled) {
            signupControls.push(
                <a
                    className='btn btn-custom-login btn--full gitlab'
                    key='gitlab'
                    href={Client.getOAuthRoute() + '/gitlab/signup' + window.location.search}
                >
                    <span className='icon'/>
                    <span>
                        <FormattedMessage
                            id='signup.gitlab'
                            defaultMessage='GitLab Single-Sign-On'
                        />
                    </span>
                </a>
            );
        }

        if (this.state.googleEnabled) {
            signupControls.push(
                <a
                    className='btn btn-custom-login btn--full google'
                    key='google'
                    href={Client.getOAuthRoute() + '/google/signup' + window.location.search + '&team=' + encodeURIComponent(this.state.teamName)}
                >
                    <span className='icon'/>
                    <span>
                        <FormattedMessage
                            id='signup.google'
                            defaultMessage='Google Account'
                        />
                    </span>
                </a>
            );
        }

        if (this.state.office365Enabled) {
            signupControls.push(
                <a
                    className='btn btn-custom-login btn--full office365'
                    key='office365'
                    href={Client.getOAuthRoute() + '/office365/signup' + window.location.search + '&team=' + encodeURIComponent(this.state.teamName)}
                >
                    <span className='icon'/>
                    <span>
                        <FormattedMessage
                            id='signup.office365'
                            defaultMessage='Office 365'
                        />
                    </span>
                </a>
           );
        }

        if (this.state.ldapEnabled) {
            signupControls.push(
                <Link
                    className='btn btn-custom-login btn--full ldap'
                    key='ldap'
                    to={'/signup_ldap'}
                >
                    <span className='icon fa fa-folder-open fa--margin-top'/>
                    <span>
                        <FormattedMessage
                            id='signup.ldap'
                            defaultMessage='LDAP Credentials'
                        />
                    </span>
                </Link>
            );
        }

        if (this.state.samlEnabled) {
            let query = '';
            if (window.location.search) {
                query = '&action=signup';
            } else {
                query = '?action=signup';
            }

            signupControls.push(
                <a
                    className='btn btn-custom-login btn--full saml'
                    key='saml'
                    href={'/login/sso/saml' + window.location.search + query}
                >
                    <span className='icon fa fa-lock fa--margin-top'/>
                    <span>
                        {global.window.mm_config.SamlLoginButtonText}
                    </span>
                </a>
            );
        }

        if (signupControls.length === 0) {
            const signupDisabledError = (
                <FormattedMessage
                    id='signup_user_completed.none'
                    defaultMessage='No user creation method has been enabled. Please contact an administrator for access.'
                />
            );
            signupControls = (
                <FormError
                    error={signupDisabledError}
                    margin={true}
                />
            );
        }

        let serverError = null;
        if (this.state.serverError) {
            serverError = (
                <div className={'form-group has-error'}>
                    <label className='control-label'>{this.state.serverError}</label>
                </div>
            );
        }

        if (this.state.noOpenServerError || this.state.usedBefore) {
            signupControls = null;
        }

        return (
            <div>
                <ErrorBar/>
                <div className='signup-header'>
                    <Link to='/'>
                        <span className='fa fa-chevron-left'/>
                        <FormattedMessage
                            id='web.header.back'
                        />
                    </Link>
                </div>
                <div className='col-sm-12'>
                    <div className='signup-team__container'>
                        <img
                            className='signup-team-logo'
                            src={logoImage}
                        />
                        <div className='signup__content'>
                            <h1>{global.window.mm_config.SiteName}</h1>
                            <h4 className='color--light'>
                                <FormattedMessage
                                    id='web.root.signup_info'
                                />
                            </h4>
                            <div className='margin--extra'>
                                <h5><strong>
                                    <FormattedMessage
                                        id='signup.title'
                                        defaultMessage='Create an account with:'
                                    />
                                </strong></h5>
                            </div>
                            {signupControls}
                            {serverError}
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}
