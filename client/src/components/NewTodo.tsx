import React, { useState } from 'react'

type Props = {}

const NewTodo = (props: Props) => {
    const [value, setValue] = useState<string>('')

    const handleSave = () => {
        setValue('')
    }

    return <div>
        <input data-testid='todo_description' value={value} />
        <button data-testid='todo_save' onClick={handleSave}>Save</button>
    </div>
}

export default NewTodo