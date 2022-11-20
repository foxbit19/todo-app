import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Grid, TextField, useMediaQuery, useTheme } from '@mui/material'
import React, { useState } from 'react'
import Item from '../../models/item'
import UpdateIcon from '@mui/icons-material/Update'
import { Box } from '@mui/system'

interface Props {
    open: boolean
    item: Item,
    onUpdateClick?: (updated: Item) => void
    onClose?: () => void
}

const UpdateTodo = (props: Props) => {
    const [inputValue, setInputValue] = useState<string>(props.item.description)

    const handleChange = (event: any) => {
        setInputValue(event.target.value)
    }

    const handleUpdateClick = () => {
        if (props.onUpdateClick) {
            props.onUpdateClick(new Item(props.item.id, inputValue, props.item.order))
        }
    }

    return <Dialog open={props.open} onClose={props.onClose} maxWidth={'lg'}>
        <DialogTitle>Item update</DialogTitle>
        <DialogContent>
            <DialogContentText>
                Enter the todo description in order to update it.
            </DialogContentText>
            <TextField required multiline fullWidth margin='dense' label='Description' data-testid='description' value={inputValue} onChange={handleChange} variant='standard' />
        </DialogContent>
        <DialogActions>
            <Button onClick={props.onClose}>Close</Button>
            <Button data-testid='update_button' variant='contained' onClick={handleUpdateClick}>Update</Button>
        </DialogActions>
    </Dialog >
}

export default UpdateTodo