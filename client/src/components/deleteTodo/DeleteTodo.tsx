import { Button } from '@mui/material'
import React from 'react'
import Item from '../../models/item'
import GenericDialog from '../dialog/GenericDialog'

interface Props {
    open: boolean
    item: Item,
    onDeleteClick?: (item: Item) => void
    onClose?: () => void
}

const DeleteTodo = (props: Props) => {
    return (
        <GenericDialog open={props.open}
            title='Are you sure you want to delete this item?'
            text={props.item.description}
            button={<Button onClick={() => props.onDeleteClick && props.onDeleteClick(props.item)}>Yes</Button>}
            onClose={props.onClose}
        />
    )
}

export default DeleteTodo