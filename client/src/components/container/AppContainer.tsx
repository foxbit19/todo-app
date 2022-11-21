import React, { ReactElement } from 'react'
import Container from '@mui/material/Container'
import Appbar from '../appbar/Appbar'

interface Props {
    children?: ReactElement
}

const AppContainer = (props: Props) => {
    return (
        <Container fixed style={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
            {props.children}
        </Container>
    )
}

export default AppContainer
