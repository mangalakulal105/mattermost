// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import Client from '../client/client.jsx';
import {browserHistory} from 'react-router';
import TeamStore from '../stores/team_store.jsx';
import BrowserStore from '../stores/browser_store.jsx';

const HTTP_UNAUTHORIZED = 401;

class WebClientClass extends Client {
    constructor() {
        super();
        this.enableLogErrorsToConsole(true);
        TeamStore.addChangeListener(this.onTeamStoreChanged);
    }

    onTeamStoreChanged = () => {
        this.setTeamId(TeamStore.getCurrentId());
    }

    track = (category, action, label, property, value) => {
        if (global.window && global.window.analytics) {
            global.window.analytics.track(action, {category, label, property, value});
        }
    }

    trackPage = () => {
        if (global.window && global.window.analytics) {
            global.window.analytics.page();
        }
    }

    handleError = (err, res) => { // eslint-disable-line no-unused-vars
        if (err.status === HTTP_UNAUTHORIZED) {
            const team = window.location.pathname.split('/')[1];
            browserHistory.push('/' + team + '/login?extra=expired');
        }
    }

    // not sure why but super.login doesn't work if using an () => arrow functions.
    // I think this might be a webpack issue.
    webLogin(email, username, password, token, success, error) {
        this.login(
            email,
            username,
            password,
            token,
            (data) => {
                this.track('api', 'api_users_login_success', '', 'email', data.email);
                BrowserStore.signalLogin();

                if (success) {
                    success(data);
                }
            },
            (err) => {
                this.track('api', 'api_users_login_fail', name, 'email', email);
                if (error) {
                    error(err);
                }
            }
        );
    }

    uploadFile = (formData, success, error) => {
        var request = $.ajax({
            url: `${this.getFilesRoute()}/upload`,
            type: 'POST',
            data: formData,
            cache: false,
            contentType: false,
            processData: false,
            success,
            error: function onError(xhr, status, err) {
                if (err !== 'abort') {
                    error(err);
                }
            }
        });

        this.track('api', 'api_files_upload');

        return request;
    }

    uploadProfileImage = (imageData, success, error) => {
        $.ajax({
            url: `${this.getUsersRoute()}/users/newimage`,
            type: 'POST',
            data: imageData,
            cache: false,
            contentType: false,
            processData: false,
            success,
            error: function onError(xhr, status, err) {
                if (err !== 'abort') {
                    error(err);
                }
            }
        });
    }
}

var WebClient = new WebClientClass();
export default WebClient;
