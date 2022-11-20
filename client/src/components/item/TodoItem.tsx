import React from 'react'
import Item from '../../models/item'
import { Draggable } from 'react-beautiful-dnd'
import { Box, Checkbox, ListItem, ListItemButton, ListItemIcon, ListItemText, Paper, styled, Typography } from '@mui/material'
import CircleIcon from '@mui/icons-material/Circle'
import CheckIcon from '@mui/icons-material/CheckCircle'

interface Props {
    index: number
    item: Item
    onClick?: () => void
    onComplete?: () => void
}

const CustomPaper: any = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(1),
    margin: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));


const TodoItem = (props: Props) => {
    const normalizeDescription = (description: string) => {
        return description.length > 50 ? `${description.substring(0, 50)}...` : description
    }

    return (
        <Draggable draggableId={props.item.id.toString()} index={props.index}>
            {(provided) => (
                <CustomPaper data-testid={`todo_${props.item.id}`} ref={provided.innerRef}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}>
                    <ListItem>
                        <Checkbox icon={<CircleIcon />} checkedIcon={<CheckIcon />} onClick={props.onComplete} />
                        <Typography onClick={props.onClick}>{normalizeDescription(props.item.description)}</Typography>
                        <Typography fontWeight={'bold'}></Typography>
                    </ListItem>
                </CustomPaper>
            )}
        </Draggable>

    )
}

export default TodoItem