import React, { ReactElement } from 'react'
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Typography } from '@mui/material'

interface Props {
    open: boolean,
    title: string,
    text: string,
    children?: ReactElement,
    button?: ReactElement,
    onClose?: () => void
}

const GenericDialog = (props: Props) => {
    return <Dialog open={props.open} onClose={props.onClose} maxWidth={'lg'}>
        <DialogTitle>{props.title}</DialogTitle>
        <DialogContent>
            <DialogContentText>
                {props.text}
            </DialogContentText>
            {props.children}
        </DialogContent>
        <DialogActions>
            <Button data-testid='close_dialog_button' onClick={props.onClose}>Close</Button>
            {props.button}
        </DialogActions>
    </Dialog>
}

export default GenericDialog