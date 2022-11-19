import React from 'react'
import Item from '../../model/item'

interface Props {
    item: Item
    onUpdateClick?: (item: Item) => void
    onDeleteClick?: (item: Item) => void
}

const ShowTodo = (props: Props) => {

    const handleUpdateClick = () => {
        if (props.onUpdateClick) {
            props.onUpdateClick(props.item)
        }
    }

    const handleDeleteClick = () => {
        if (props.onDeleteClick) {
            props.onDeleteClick(props.item)
        }
    }

    return <div>
        Item details
        <div>{props.item.description}</div>
        <button data-testid='update_button' onClick={handleUpdateClick}></button>
        <button data-testid='delete_button' onClick={handleDeleteClick}></button>
    </div>
}

export default ShowTodo