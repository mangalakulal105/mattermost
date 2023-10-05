import React from 'react';
import { AllowedIPRange } from '@mattermost/types/config';
import { Button, Modal } from 'react-bootstrap';
import { useIntl } from 'react-intl';

import './delete_confirmation.scss';

type Props = {
    onClose?: () => void;
    onConfirm?: (filter: AllowedIPRange) => void;
    filterToDelete?: AllowedIPRange;
}

export default function DeleteConfirmationModal({ onClose, onConfirm, filterToDelete }: Props) {
    const {formatMessage} = useIntl();
    return (
        <Modal
            className={'DeleteConfirmationModal'}
            dialogClassName={'DeleteConfirmationModal__dialog'}
            show={true}
            onHide={() => onClose?.()}
        >
            <Modal.Header closeButton={true}>
                <div className='title'>
                    {formatMessage({id: 'admin.ip_filtering.delete_confirmation_title', defaultMessage: 'Delete IP Filter'})}
                </div>
            </Modal.Header>
            <Modal.Body>
                {formatMessage({
                    id: 'admin.ip_filtering.delete_confirmation_body',
                    defaultMessage: 'Are you sure you want to delete IP filter {filter}? Users with IP addresses outside of this range won\'t be able to access the workspace when IP Filtering is enabled'
                },
                    { filter: (<strong>{filterToDelete?.Description}</strong>) }
                )}
            </Modal.Body>
            <Modal.Footer>
                <Button
                    type="button"
                    className="btn-cancel"
                    onClick={() => onClose?.()}
                >
                    {formatMessage({id: 'admin.ip_filtering.cancel', defaultMessage: 'Cancel'})}
                </Button>
                <Button
                    type="button"
                    className="btn-delete"
                    onClick={() => onConfirm?.(filterToDelete!)}
                >
                    {formatMessage({id: 'admin.ip_filtering.delete_filter', defaultMessage: 'Delete filter'})}
                </Button>
            </Modal.Footer>
        </Modal>
    )
}