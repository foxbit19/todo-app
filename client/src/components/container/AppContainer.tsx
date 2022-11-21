import React, { ReactElement } from 'react'
import Container from '@mui/material/Container'

interface Props {
    children?: ReactElement
}

const AppContainer = (props: Props) => {
    return (
        <Container fixed>{props.children}</Container>
    )
}

export default AppContainer
