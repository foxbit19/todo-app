import React from 'react'
import { Box, Link, Typography } from '@mui/material'
import LoveIcon from '@mui/icons-material/Favorite'

interface Props { }

const Footer = (props: Props) => {
    return (
        <Typography variant='subtitle2' marginTop={10} marginBottom={2} align='center'>
            Made with <LoveIcon fontSize='small' color='primary' /> by <Link href='https://github.com/foxbit19'>Mattia Peretti</Link>.
        </Typography>
    )
}

export default Footer