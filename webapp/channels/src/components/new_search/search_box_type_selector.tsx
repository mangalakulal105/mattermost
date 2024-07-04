// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {useCallback} from 'react';
import {FormattedMessage} from 'react-intl';
import {useSelector} from 'react-redux';
import styled from 'styled-components';

import {getLicense} from 'mattermost-redux/selectors/entities/general';

import ErrorBoundary from 'plugins/pluggable/error_boundary';

import type {GlobalState} from 'types/store';

const SearchTypeSelectorContainer = styled.div`
    margin: 24px 32px 0px 24px;
    display: flex;
    align-items: center;
    padding: 4px;
    background-color: var(--center-channel-bg);
    border-radius: var(--radius-m);
    border: var(--border-default);
    width: fit-content;
    gap: 4px;
`;

type SearchTypeItemProps = {
    selected: boolean;
};

const SearchTypeItem = styled.button<SearchTypeItemProps>`
    display: flex;
    cursor: pointer;
    padding: 4px 10px;
    background-color: ${(props) => (props.selected ? 'rgba(var(--button-bg-rgb), 0.08)' : 'transparent')};
    color: ${(props) => (props.selected ? 'var(--button-bg)' : 'rgba(var(--center-channel-color-rgb), 0.75)')};
    border-radius: 4px;
    font-size: 12px;
    line-height: 16px;
    font-weight: 600;
    border: none;
    &:hover {
        color: rgba(var(--center-channel-color-rgb), 0.88);
        background: rgba(var(--center-channel-color-rgb), 0.08);
    }
`;

type Props = {
    searchType: string;
    setSearchType: (searchType: string) => void;
}

const SearchTypeSelector = ({searchType, setSearchType}: Props) => {
    const license = useSelector(getLicense);
    const setMessagesSearchType = useCallback(() => setSearchType('messages'), [setSearchType]);
    const setFilesSearchType = useCallback(() => setSearchType('messages'), [setSearchType]);

    let SearchPluginButtons = useSelector((state: GlobalState) => state.plugins.components.SearchButtons) || [];
    if (license.IsLicensed !== 'true') {
        SearchPluginButtons = [];
    }
    return (
        <SearchTypeSelectorContainer>
            <SearchTypeItem
                selected={searchType === 'messages'}
                onClick={setMessagesSearchType}
            >
                <FormattedMessage
                    id='search_bar.usage.search_type_messages'
                    defaultMessage='Messages'
                />
            </SearchTypeItem>
            <SearchTypeItem
                selected={searchType === 'files'}
                onClick={setFilesSearchType}
            >
                <FormattedMessage
                    id='search_bar.usage.search_type_files'
                    defaultMessage='Files'
                />
            </SearchTypeItem>
            {SearchPluginButtons.map(({component, pluginId}: any) => {
                const Component = component as React.ComponentType;
                return (
                    <SearchTypeItem
                        key={pluginId}
                        selected={searchType === pluginId}
                        onClick={() => setSearchType(pluginId)}
                    >
                        <ErrorBoundary>
                            <Component/>
                        </ErrorBoundary>
                    </SearchTypeItem>
                );
            })}
        </SearchTypeSelectorContainer>
    );
};

export default SearchTypeSelector;
