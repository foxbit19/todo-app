import React from 'react'
import { List, ListItem, ListItemText, Typography } from '@mui/material';
import GenericDialog from '../dialog/GenericDialog';
import Item from '../../models/item';
import dayjs from 'dayjs'
import LocalizedFormat from 'dayjs/plugin/localizedFormat'

dayjs.extend(LocalizedFormat)

interface Props {
    open: boolean
    items: Item[]
    onClose?: () => void
}

const CompletedTodos = (props: Props) => {
    return (
        <GenericDialog
            open={props.open}
            onClose={props.onClose}
            title='Completed items'
            text='These are your completed items reverse ordered by completed date.'
        >
            <>
                {!props.items && (
                    <Typography variant='body2'>
                        There are no completed items.
                    </Typography>
                )}
                {props.items && (
                    <List>
                        {props.items.map((item: Item) => (
                            <ListItem divider>
                                <ListItemText primary={item.description} secondary={`Completed on ${dayjs(item.completedDate).format('LLLL')}`} />
                            </ListItem>
                        ))}
                    </List>
                )}
            </>
        </GenericDialog>
    )
}

export default CompletedTodos