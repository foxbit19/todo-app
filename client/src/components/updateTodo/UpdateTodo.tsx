import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, TextField } from '@mui/material'
import React, { useEffect, useState } from 'react'
import Item from '../../models/item'
import GenericDialog from '../dialog/GenericDialog'

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

    useEffect(() => {
        if (props.item) {
            setInputValue(props.item.description)
        }
    }, [props.item])


    return (
        <GenericDialog open={props.open}
            title='Item update'
            text={'Edit the todo description in order to update it.'}
            button={<Button data-testid='update_button' variant='contained' onClick={handleUpdateClick}>Update</Button>}
            onClose={props.onClose}>
            <TextField autoFocus required fullWidth margin='dense' label='Description' data-testid='description' value={inputValue} onChange={handleChange} variant='standard' />
        </GenericDialog>
    )
}

export default UpdateTodo