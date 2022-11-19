import React, { useState } from 'react'
import Item from '../../models/item'

interface Props {
    item: Item,
    onUpdateClick?: (updated: Item) => void
}

const UpdateTodo = (props: Props) => {
    const [inputValue, setInputValue] = useState<string>(props.item.description)

    const handleChange = (event: any) => {
        setInputValue(event.target.value)
    }

    const handleUpdateClick = () => {
        if (props.onUpdateClick) {
            props.onUpdateClick(new Item(props.item.id, inputValue, props.item.order))
        }
    }

    return <div>
        <div>Item update</div>
        <input data-testid='description' value={inputValue} onChange={handleChange} />
        <button data-testid='update_button' onClick={handleUpdateClick}>Update</button>
    </div>
}

export default UpdateTodo