import { Alert, AlertColor, AlertTitle, Box, Button, createTheme, CssBaseline, Grid, Paper, styled, ThemeProvider } from '@mui/material';
import React, { useEffect, useState } from 'react';
import TodoList from './components/list/TodoList';
import ShowTodo from './components/showTodo/ShowTodo';
import Title from './components/title/Title';
import Item from './models/item';
import ItemService from './services/itemService';
import UpdateTodo from './components/updateTodo/UpdateTodo';
import AddIcon from '@mui/icons-material/Add'

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
    },
});

const CustomPaper = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(1),
    margin: '10%',
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));

const buildAlert = (severity: AlertColor, title: string, message: string) => (
    <Alert severity={severity} onClose={() => { }}>
        <AlertTitle>{title}</AlertTitle>
        {message}
    </Alert>
)

function App() {
    const [items, setItems] = useState<Item[]>([])
    const [showTodo, setShowTodo] = useState<boolean>(false)
    const [showUpdate, setShowUpdate] = useState<boolean>(false)
    const [item, setItem] = useState<Item>()
    const [alert, setAlert] = useState<JSX.Element>(<></>)
    const service = new ItemService()

    const getAllItems = async () => {
        setItems(await service.getAll())
    }

    useEffect(() => {
        getAllItems()
    }, [])

    const handleItemClick = (item: Item) => {
        setItem(item)
        setShowTodo(true)
    }

    const handleClose = () => {
        setShowTodo(false)
        setShowUpdate(false)
    }

    const openUpdateModal = async (item: Item) => {
        setShowUpdate(true)
        setShowTodo(false)
    }

    const handleUpdate = async (item: Item) => {
        try {
            await service.update(item)
            setAlert(buildAlert('success', 'Item updated', 'Your item was updated successfully'))
            getAllItems()
        } catch (error: any) {
            console.error(error)
            setAlert(buildAlert('error', 'Item update fails', 'This item cannot be updated'))
        } finally {
            // clear the current item
            setItem(undefined)
        }
    }

    const handleComplete = async (item: Item) => {
        try {
            await service.delete(item.id)
            setAlert(buildAlert('success', 'Item completed', 'Your item is now complete'))
            getAllItems()
        } catch (error: any) {
            console.error(error)
            setAlert(buildAlert('error', 'Item complete fails', 'This item cannot be completed'))
        } finally {
            setItem(undefined)
        }

    }

    return (
        <ThemeProvider theme={darkTheme}>
            <CssBaseline />
            <Box sx={{ flexGrow: 1 }}>
                <Title />
                {alert}
                <Button data-testid='new_button' variant='contained' startIcon={<AddIcon />}>Add new</Button>
                <CustomPaper><TodoList items={items} onItemClick={handleItemClick} onComplete={handleComplete} /></CustomPaper>
            </Box>
            {item && <ShowTodo open={showTodo} item={item} onClose={handleClose} onUpdateClick={openUpdateModal} />}
            {item && <UpdateTodo open={showUpdate} item={item} onClose={handleClose} onUpdateClick={handleUpdate} />}
        </ThemeProvider>
    )
}

export default App;
