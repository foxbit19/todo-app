import React from 'react'

interface Props { }

const Title = (props: Props) => {
    return <div data-testid="title">
        <div>ToDo App</div>
        <div>The perfect solution for yours day to day ToDos</div>
    </div>
}

export default Title