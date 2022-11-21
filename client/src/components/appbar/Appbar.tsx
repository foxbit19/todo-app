import React from 'react'
import { AppBar, Box, Button, Container, Grid, Toolbar, Typography } from '@mui/material'

interface Props { }

const Appbar = (props: Props) => {
    return (
        <AppBar position="static" enableColorOnDark>
            <Toolbar>
                <Typography variant="h5" style={{ color: 'black', fontFamily: 'Satisfy' }} fontWeight='bold' flexGrow={1}>
                    Easy to do
                </Typography>
                <Button sx={{ color: 'black', marginLeft: 5 }}>
                    show completed
                </Button>
            </Toolbar>
        </AppBar>
    )
}

export default Appbar