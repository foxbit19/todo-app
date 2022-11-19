import React from 'react'
import Item from '../../models/item'

interface Props {
    items: Item[]
    onItemClick?: (item: Item) => void
}

const TodoList = (props: Props) => {
    const normalizeDescription = (description: string) => {
        return description.length > 50 ? `${description.substring(0, 50)}...` : description
    }

    return <div data-testid="list">
        {props.items.map((item: Item, index: number) => (
            <div key={index} data-testid={`todo_${item.id}`} onClick={() => props.onItemClick && props.onItemClick(item)}>{normalizeDescription(item.description)}</div>
        ))}
        {props.items.length === 0 ? <div>There are no todos to show</div> : ''}
    </div>
}

export default TodoList
