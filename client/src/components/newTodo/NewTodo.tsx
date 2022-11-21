import React, { useState } from 'react'
import { Button, TextField } from '@mui/material'
import GenericDialog from '../dialog/GenericDialog'

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

    return (
        <GenericDialog open={props.open}
            title='New todo item'
            text={'Enter a description for your new todo item.'}
            button={<Button data-testid='todo_save' variant='contained' onClick={handleSave}>Save</Button>}
            onClose={props.onClose}>
            <TextField autoFocus required fullWidth margin='dense' label='Description' data-testid='todo_description' value={inputValue} onChange={handleChange} variant='standard' />
        </GenericDialog>
    )
}

export default NewTodo