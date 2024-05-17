// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useSelector} from 'react-redux';
import {Link} from 'react-router-dom';

import {Client4} from 'mattermost-redux/client';
import {getConfig, getLicense} from 'mattermost-redux/selectors/entities/general';

import BackButton from 'components/common/back_button';
import Logo from 'components/common/svg_images_components/logo_dark_blue_svg';

import './header.scss';

export type HeaderProps = {
    alternateLink?: React.ReactElement;
    backButtonURL?: string;
    onBackButtonClick?: React.EventHandler<React.MouseEvent>;
}

const Header = ({alternateLink, backButtonURL, onBackButtonClick}: HeaderProps) => {
    const license = useSelector(getLicense);

    const {
        EnableCustomBrand,
        SiteName,
        CustomBrandHasLogo,
    } = useSelector(getConfig);

    const ariaLabel = SiteName || 'Mattermost';

    let freeBanner = null;
    if (license.IsLicensed === 'false') {
        freeBanner = <><Logo/><span className='freeBadge'>{'FREE EDITION'}</span></>;
    }

    let title: React.ReactNode = SiteName;
    if (title === 'Mattermost') {
        if (freeBanner) {
            title = '';
        } else {
            title = <Logo/>;
        }
    }

    if (EnableCustomBrand === 'true' && CustomBrandHasLogo === 'true') {
        title = <img src={Client4.getCustomLightLogoUrl('0')}/>
    }

    return (
        <div className='hfroute-header'>
            <div className='header-main'>
                <div>
                    {freeBanner &&
                        <Link
                            className='header-logo-link'
                            to='/'
                            aria-label={ariaLabel}
                        >
                            {freeBanner}
                        </Link>
                    }
                    {title &&
                        <Link
                            className='header-logo-link'
                            to='/'
                            aria-label={ariaLabel}
                        >
                            {title}
                        </Link>
                    }
                </div>
                {alternateLink}
            </div>
            {onBackButtonClick && (
                <BackButton
                    className='header-back-button'
                    url={backButtonURL}
                    onClick={onBackButtonClick}
                />
            )}
        </div>
    );
};

export default Header;
