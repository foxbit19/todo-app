import React, { useState } from 'react'

type Props = {}

const NewTodo = (props: Props) => {
    const [inputValue, setInputValue] = useState<string | undefined>('')
    const [item, setItem] = useState<string | undefined>('')

    const handleChange = (event: any) => {
        setInputValue(event.target.value)
    }

    const handleSave = () => {
        setItem(inputValue)
        setInputValue('')
    }

    return <div>
        <input data-testid='todo_description' value={inputValue} onChange={handleChange} />
        <button data-testid='todo_save' onClick={handleSave}>Save</button>
        <div>{item}</div>
    </div>
}

export default NewTodo