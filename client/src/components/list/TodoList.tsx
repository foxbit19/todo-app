import React from 'react'
import { List } from '@mui/material'
import Item from '../../models/item'
import { DragDropContext, Droppable } from 'react-beautiful-dnd'
import TodoItem from '../item/TodoItem';

interface Props {
    items: Item[]
    onItemClick?: (item: Item) => void
    onComplete?: (item: Item) => void
}

const TodoList = (props: Props) => {
    return (
        <DragDropContext onDragEnd={() => { }}>
            <Droppable droppableId='1'>
                {(provided) => <List data-testid="list" ref={provided.innerRef} {...provided.droppableProps}>
                    {props.items.map((item: Item, index: number) => (
                        <TodoItem index={index} item={item}
                            onClick={() => props.onItemClick && props.onItemClick(item)}
                            onComplete={() => props.onComplete && props.onComplete(item)} />
                    ))}
                    {provided.placeholder}
                </List >}
            </Droppable>
            {props.items.length === 0 ? <div>There are no todos to show</div> : ''}
        </DragDropContext>
    )
}

export default TodoList
