import React from 'react'
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Slide } from '@mui/material'
import Item from '../../models/item'
import UpdateIcon from '@mui/icons-material/Update'
import DoneIcon from '@mui/icons-material/Done'
import { TransitionProps } from '@mui/material/transitions'

interface Props {
    open: boolean
    item: Item
    onUpdateClick?: (item: Item) => void
    onDeleteClick?: (item: Item) => void
    onClose?: () => void
}

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement<any, any>;
    },
    ref: React.Ref<unknown>,
) {
    return <Slide direction="up" ref={ref} {...props} />;
});

const ShowTodo = (props: Props) => {

    const handleUpdateClick = () => {
        if (props.onUpdateClick) {
            props.onUpdateClick(props.item)
        }
    }

    const handleDeleteClick = () => {
        if (props.onDeleteClick) {
            props.onDeleteClick(props.item)
        }
    }

    return <Dialog open={props.open} onClose={props.onClose} TransitionComponent={Transition}>
        <DialogTitle>Item details</DialogTitle>
        <DialogContent>
            <DialogContentText>
                {props.item.description}
            </DialogContentText>
        </DialogContent>
        <DialogActions>
            <Button data-testid='update_button' onClick={handleUpdateClick} startIcon={<UpdateIcon />}>Update</Button>
            <Button data-testid='delete_button' onClick={handleDeleteClick} startIcon={<DoneIcon />}>Complete</Button>
        </DialogActions>
    </Dialog>
}

export default ShowTodo