import React from 'react'
import { List, ListItemButton, ListItemIcon, ListItemText } from '@mui/material'
import Item from '../../models/item'
import DoneIcon from '@mui/icons-material/Done'
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd'
interface Props {
    items: Item[]
    onItemClick?: (item: Item) => void
}

const TodoList = (props: Props) => {
    const normalizeDescription = (description: string) => {
        return description.length > 50 ? `${description.substring(0, 50)}...` : description
    }

    return (
        <DragDropContext onDragEnd={() => { }}>
            <Droppable droppableId='1'>
                {(provided) => <List data-testid="list" ref={provided.innerRef} {...provided.droppableProps}>
                    {props.items.map((item: Item, index: number) => (
                        <Draggable draggableId={item.id.toString()} index={index}>
                            {(provided) => (
                                <ListItemButton key={index} data-testid={`todo_${item.id}`} onClick={() => props.onItemClick && props.onItemClick(item)} ref={provided.innerRef}
                                    {...provided.draggableProps}
                                    {...provided.dragHandleProps}>
                                    <ListItemIcon><DoneIcon /></ListItemIcon>
                                    <ListItemText>{normalizeDescription(item.description)}</ListItemText>
                                </ListItemButton>
                            )}
                        </Draggable>
                    ))}
                    {props.items.length === 0 ? <div>There are no todos to show</div> : ''}
                    {provided.placeholder}
                </List >}
            </Droppable>
        </DragDropContext>
    )
}

export default TodoList
