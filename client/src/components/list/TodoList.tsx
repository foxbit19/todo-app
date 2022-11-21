import React from 'react'
import { List } from '@mui/material'
import Item from '../../models/item'
import { DragDropContext, Droppable, DropResult } from '@hello-pangea/dnd'
import TodoItem from '../item/TodoItem';

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

    return <>
        {props.items.length === 0 ?
            <div>There are no todos to show</div> : (
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
    </>

}

export default TodoList
