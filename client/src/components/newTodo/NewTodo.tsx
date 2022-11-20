import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, TextField } from '@mui/material'
import React, { useState } from 'react'

interface Props {
    open: boolean,
    onSaveClick?: (description: string) => void
    onClose?: () => void
}

const NewTodo = (props: Props) => {
    const [inputValue, setInputValue] = useState<string>('')

    const handleChange = (event: any) => {
        setInputValue(event.target.value)
    }

    const handleSave = () => {
        if (inputValue.length === 0) {
            return;
        }

        if (props.onSaveClick) {
            props.onSaveClick(inputValue)
        }
        setInputValue('')
    }

    return <Dialog open={props.open} onClose={props.onClose} maxWidth={'lg'}>
        <DialogTitle>New todo</DialogTitle>
        <DialogContent>
            <DialogContentText>
                Enter a description for your new todo item.
            </DialogContentText>
            <TextField autoFocus required fullWidth margin='dense' label='Description' data-testid='todo_description' value={inputValue} onChange={handleChange} variant='standard' />
        </DialogContent>
        <DialogActions>
            <Button onClick={props.onClose}>Close</Button>
            <Button data-testid='todo_save' variant='contained' onClick={handleSave}>Save</Button>
        </DialogActions>
    </Dialog>
}

export default NewTodo