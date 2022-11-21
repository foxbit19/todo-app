import React from 'react'
import Item from '../../models/item'
import { Draggable } from '@hello-pangea/dnd'
import { Checkbox, Grid, IconButton, ListItem, Paper, styled, Typography } from '@mui/material'
import CircleIcon from '@mui/icons-material/Circle'
import CheckIcon from '@mui/icons-material/CheckCircle'
import UpdateIcon from '@mui/icons-material/Edit'
import DeleteIcon from '@mui/icons-material/Delete'

interface Props {
    index: number
    item: Item
    onUpdate?: (item: Item) => void
    onDelete?: (item: Item) => void
    onComplete?: (item: Item) => void
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
    const handleComplete = () => {
        if (props.onComplete) {
            props.onComplete(props.item)
        }
    }

    const handleUpdate = () => {
        if (props.onUpdate) {
            props.onUpdate(props.item)
        }
    }

    const handleDelete = () => {
        if (props.onDelete) {
            props.onDelete(props.item)
        }
    }

    return (
        <Draggable key={props.item.id} draggableId={`${props.item.id}${props.index}`} index={props.index} >
            {(provided) => (
                <CustomPaper data-testid={`todo_${props.item.id}`} ref={provided.innerRef}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}>
                    <ListItem>
                        <Grid container direction='row' alignItems="center" wrap='nowrap'>
                            <Grid item>
                                <Checkbox icon={<CircleIcon />} checkedIcon={<CheckIcon />} onClick={handleComplete} />
                            </Grid>
                            <Grid item xs={8} md={11}>
                                <Typography padding={1}>{props.item.description}</Typography>
                            </Grid>
                            <Grid item>
                                <IconButton color='primary' title='Edit this item' onClick={handleUpdate}><UpdateIcon /></IconButton>
                            </Grid>
                            <Grid item>
                                <IconButton color='primary' title='Delete this item' onClick={handleDelete}><DeleteIcon /></IconButton>
                            </Grid>
                        </Grid>
                    </ListItem>
                </CustomPaper>
            )}
        </Draggable>

    )
}

export default TodoItem