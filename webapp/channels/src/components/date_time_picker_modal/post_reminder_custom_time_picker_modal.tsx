// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import type {Moment} from 'moment-timezone';
import React, {useCallback, useEffect, useState} from 'react';
import {useSelector} from 'react-redux';
import Constants from 'utils/constants';
import {isKeyPressed} from 'utils/keyboard';
import {getCurrentMomentForTimezone} from 'utils/timezone';

import {GenericModal} from '@mattermost/components';

import {getCurrentTimezone} from 'mattermost-redux/selectors/entities/timezone';

import DateTimeInput, {getRoundedTime} from 'components/custom_status/date_time_input';
import './style.scss';

type Props = {
    onExited: () => void;
    ariaLabel: string;
    header: string | React.ReactNode;
    onChange?: (datTime: Moment) => void;
    onCancel?: () => void;
    onConfirm?: (dateTime: Moment) => void;
    initialTime?: Moment;
    confirmButtonText?: string | React.ReactNode;
};

export default function DateTimePickerModal({onExited, ariaLabel, header, onConfirm, onCancel, initialTime, confirmButtonText, onChange}: Props) {
    const userTimezone = useSelector(getCurrentTimezone);
    const currentTime = getCurrentMomentForTimezone(userTimezone);
    const initialRoundedTime = getRoundedTime(currentTime);

    const [dateTime, setDateTime] = useState(initialTime || initialRoundedTime);

    const [isDatePickerOpen, setIsDatePickerOpen] = useState(false);

    useEffect(() => {
        function handleKeyDown(event: React.KeyboardEvent) {
            if (isKeyPressed(event, Constants.KeyCodes.ESCAPE) && !isDatePickerOpen) {
                onExited();
            }
        }

        document.addEventListener('keydown', handleKeyDown);

        return () => {
            document.removeEventListener('keydown', handleKeyDown);
        };
    }, [isDatePickerOpen, onExited]);

    const handleChange = useCallback((dateTime: Moment) => {
        setDateTime(dateTime);
        if (onChange) {
            onChange(dateTime);
        }
    }, [onChange]);

    const handleConfirm = useCallback(() => {
        if (onConfirm) {
            onConfirm(dateTime);
        }
    }, [dateTime, onConfirm]);

    return (
        <GenericModal
            id='DateTimePickerModal'
            ariaLabel={ariaLabel}
            onExited={onExited}
            modalHeaderText={header}
            confirmButtonText={confirmButtonText}
            handleConfirm={handleConfirm}
            handleCancel={onCancel}
            handleEnterKeyPress={handleConfirm}
            className={'date-time-picker-modal'}
            compassDesign={true}
            keyboardEscape={false}
        >
            <DateTimeInput
                time={dateTime}
                handleChange={handleChange}
                timezone={userTimezone}
                setIsDatePickerOpen={setIsDatePickerOpen}
            />
        </GenericModal>
    );
}
