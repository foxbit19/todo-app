import React from 'react'
import Item from '../../models/item'
import { Draggable } from 'react-beautiful-dnd'
import { Box, Checkbox, ListItemButton, ListItemIcon, ListItemText } from '@mui/material'
import CircleIcon from '@mui/icons-material/Circle'
import CheckIcon from '@mui/icons-material/CheckCircle'

interface Props {
    index: number
    item: Item
    onClick?: () => void
    onComplete?: () => void
}

const TodoItem = (props: Props) => {
    const normalizeDescription = (description: string) => {
        return description.length > 50 ? `${description.substring(0, 50)}...` : description
    }

    return (
        <Draggable draggableId={props.item.id.toString()} index={props.index}>
            {(provided) => (
                <ListItemButton data-testid={`todo_${props.item.id}`} ref={provided.innerRef}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}>
                    <Checkbox icon={<CircleIcon />} checkedIcon={<CheckIcon />} onClick={props.onComplete} />
                    <ListItemText onClick={props.onClick}>{normalizeDescription(props.item.description)}</ListItemText>
                </ListItemButton>
            )}
        </Draggable>
    )
}

export default TodoItem