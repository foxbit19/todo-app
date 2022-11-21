import React from 'react'
import { Grid, List, Typography } from '@mui/material'
import Item from '../../models/item'
import { DragDropContext, Droppable, DropResult } from '@hello-pangea/dnd'
import TodoItem from '../item/TodoItem';
import InfoIcon from '@mui/icons-material/Info'

interface Props {
    items: Item[]
    onUpdate?: (item: Item) => void
    onDelete?: (item: Item) => void
    onComplete?: (item: Item) => void
    onReorder?: (sourceIndex: number, targetIndex: number) => void
}

const TodoList = (props: Props) => {
    const handleDragEnd = (result: DropResult) => {
        if (props.onReorder && result.destination) {
            props.onReorder(result.source.index, result.destination.index)
        }
    }

    return <div style={{ marginTop: '2em' }}>
        <Grid container flexDirection='row' justifyContent='center' alignItems='center' spacing={2}>
            <Grid item>
                <InfoIcon />
            </Grid>
            <Grid item>
                <Typography variant='subtitle2' align='center'>
                    You can add, edit, delete and mark as complete an item. Drag an item over the list to change its priority.
                </Typography>
            </Grid>
        </Grid>
        {props.items.length === 0 ?
            <Typography align='center' margin={2} variant='h4'>There are no todos to show</Typography> : (
                <DragDropContext onDragEnd={handleDragEnd}>
                    <Droppable droppableId='Todo'>
                        {
                            (provided) => <List data-testid="list" ref={provided.innerRef} {...provided.droppableProps}>
                                {props.items.map((item: Item, index: number) => (
                                    <TodoItem key={index} index={index} item={item}
                                    onUpdate={props.onUpdate}
                                    onDelete={props.onDelete}
                                    onComplete={props.onComplete}
                                />
                            ))}
                                {provided.placeholder}
                            </List >
                        }
                    </Droppable>
                </DragDropContext>
            )
        }
    </div>

}

export default TodoList
