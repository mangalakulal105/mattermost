// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useIntl} from 'react-intl';

import {InformationOutlineIcon} from '@mattermost/compass-icons/components';

import ExternalLink from 'components/external_link';
import OverlayTrigger from 'components/overlay_trigger';

import Constants from 'utils/constants';

import TitleTooltip from './title_tooltip';

const Title = () => {
    const intl = useIntl();
    return (
        <OverlayTrigger
            delayShow={Constants.OVERLAY_TIME_DELAY}
            placement='top'
            overlay={<TitleTooltip/>}
        >
            <span>
                <ExternalLink href='https://docs.mattermost.com/configure/reporting-configuration-settings.html#site-statistics'>
                    {intl.formatMessage({id: 'analytics.team.totalUsers', defaultMessage: 'Total Activated Users'})}
                    <InformationOutlineIcon size='16'/>
                </ExternalLink>
            </span>
        </OverlayTrigger>
    );
};

export default Title;
