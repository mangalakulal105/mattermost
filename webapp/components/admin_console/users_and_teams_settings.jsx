// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import React from 'react';

import * as Utils from 'utils/utils.jsx';

import AdminSettings from './admin_settings.jsx';
import BooleanSetting from './boolean_setting.jsx';
import {FormattedMessage} from 'react-intl';
import SettingsGroup from './settings_group.jsx';
import TextSetting from './text_setting.jsx';

export class UsersAndTeamsSettingsPage extends AdminSettings {
    constructor(props) {
        super(props);

        this.getConfigFromState = this.getConfigFromState.bind(this);

        this.renderSettings = this.renderSettings.bind(this);

        this.state = Object.assign(this.state, {
            enableUserCreation: props.config.TeamSettings.EnableUserCreation,
            enableTeamCreation: props.config.TeamSettings.EnableTeamCreation,
            maxUsersPerTeam: props.config.TeamSettings.MaxUsersPerTeam,
            restrictCreationToDomains: props.config.TeamSettings.RestrictCreationToDomains,
            restrictTeamNames: props.config.TeamSettings.RestrictTeamNames
        });
    }

    getConfigFromState(config) {
        config.TeamSettings.EnableUserCreation = this.state.enableUserCreation;
        config.TeamSettings.EnableTeamCreation = this.state.enableTeamCreation;
        config.TeamSettings.MaxUsersPerTeam = this.parseIntNonZero(this.state.maxUsersPerTeam);
        config.TeamSettings.RestrictCreationToDomains = this.state.restrictCreationToDomains;
        config.TeamSettings.RestrictTeamNames = this.state.restrictTeamNames;

        return config;
    }

    renderTitle() {
        return (
            <h3>
                <FormattedMessage
                    id='admin.general.title'
                    defaultMessage='General Settings'
                />
            </h3>
        );
    }

    renderSettings() {
        return (
            <UsersAndTeamsSettings
                enableUserCreation={this.state.enableUserCreation}
                enableTeamCreation={this.state.enableTeamCreation}
                maxUsersPerTeam={this.state.maxUsersPerTeam}
                restrictCreationToDomains={this.state.restrictCreationToDomains}
                restrictTeamNames={this.state.restrictTeamNames}
                onChange={this.handleChange}
            />
        );
    }
}

export class UsersAndTeamsSettings extends React.Component {
    static get propTypes() {
        return {
            enableUserCreation: React.PropTypes.bool.isRequired,
            enableTeamCreation: React.PropTypes.bool.isRequired,
            maxUsersPerTeam: React.PropTypes.oneOfType([
                React.PropTypes.number,
                React.PropTypes.string
            ]).isRequired,
            restrictCreationToDomains: React.PropTypes.string.isRequired,
            restrictTeamNames: React.PropTypes.bool.isRequired,
            onChange: React.PropTypes.func.isRequired
        };
    }

    render() {
        return (
            <SettingsGroup
                header={
                    <FormattedMessage
                        id='admin.general.usersAndTeams'
                        defaultMessage='Users and Teams'
                    />
                }
            >
                <BooleanSetting
                    id='enableUserCreation'
                    label={
                        <FormattedMessage
                            id='admin.team.userCreationTitle'
                            defaultMessage='Enable User Creation: '
                        />
                    }
                    helpText={
                        <FormattedMessage
                            id='admin.team.userCreationDescription'
                            defaultMessage='When false, the ability to create accounts is disabled. The create account button displays error when pressed.'
                        />
                    }
                    value={this.props.enableUserCreation}
                    onChange={this.props.onChange}
                />
                <BooleanSetting
                    id='enableTeamCreation'
                    label={
                        <FormattedMessage
                            id='admin.team.teamCreationTitle'
                            defaultMessage='Enable Team Creation: '
                        />
                    }
                    helpText={
                        <FormattedMessage
                            id='admin.team.teamCreationDescription'
                            defaultMessage='When false, the ability to create teams is disabled. The create team button displays error when pressed.'
                        />
                    }
                    value={this.props.enableTeamCreation}
                    onChange={this.props.onChange}
                />
                <TextSetting
                    id='maxUsersPerTeam'
                    label={
                        <FormattedMessage
                            id='admin.team.maxUsersTitle'
                            defaultMessage='Max Users Per Team:'
                        />
                    }
                    placeholder={Utils.localizeMessage('admin.team.maxUsersExample', 'Ex "25"')}
                    helpText={
                        <FormattedMessage
                            id='admin.team.maxUsersDescription'
                            defaultMessage='Maximum total number of users per team, including both active and inactive users.'
                        />
                    }
                    value={this.props.maxUsersPerTeam}
                    onChange={this.props.onChange}
                />
                <TextSetting
                    id='restrictCreationToDomains'
                    label={
                        <FormattedMessage
                            id='admin.team.restrictTitle'
                            defaultMessage='Restrict Creation To Domains:'
                        />
                    }
                    placeholder={Utils.localizeMessage('admin.team.restrictExample', 'Ex "corp.mattermost.com, mattermost.org"')}
                    helpText={
                        <FormattedMessage
                            id='admin.team.restrictDescription'
                            defaultMessage='Teams and user accounts can only be created from a specific domain (e.g. "mattermost.org") or list of comma-separated domains (e.g. "corp.mattermost.com, mattermost.org").'
                        />
                    }
                    value={this.props.restrictCreationToDomains}
                    onChange={this.props.onChange}
                />
                <BooleanSetting
                    id='restrictTeamNames'
                    label={
                        <FormattedMessage
                            id='admin.team.restrictNameTitle'
                            defaultMessage='Restrict Team Names: '
                        />
                    }
                    helpText={
                        <FormattedMessage
                            id='admin.team.restrictNameDesc'
                            defaultMessage='When true, You cannot create a team name with reserved words like www, admin, support, test, channel, etc'
                        />
                    }
                    value={this.props.restrictTeamNames}
                    onChange={this.props.onChange}
                />
            </SettingsGroup>
        );
    }
}
