import { Grid, Typography } from '@mui/material'
import React from 'react'

interface Props { }

const Title = (props: Props) => {
    return (
        <Grid container direction='column' alignItems='center' spacing={2} minHeight={160} data-testid="title" >
            <Grid item xs>
                <Typography variant='h2'>ToDo App</Typography>
            </Grid>
            <Grid item xs>
                <Typography variant='h6'>The perfect solution for yours day to day ToDos</Typography>
            </Grid>
        </Grid>
    )
}

export default Title