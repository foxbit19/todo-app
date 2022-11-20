import { Grid, Typography } from '@mui/material'
import React from 'react'

interface Props { }

const Title = (props: Props) => {
    return (
        <Grid container spacing={2} minHeight={160} data-testid="title">
            <Grid xs={12} display="flex" justifyContent="center" alignItems="center">
                <Typography variant='h2'>ToDo App</Typography>
            </Grid>
            <Grid xs={12} display="flex" justifyContent="center" alignItems="center">
                <Typography variant='h6'>The perfect solution for yours day to day ToDos</Typography>
            </Grid>
        </Grid>
    )
}

export default Title