import { Alert, AlertColor, AlertTitle, createTheme, CssBaseline, Fab, ThemeProvider } from '@mui/material';
import React, { useEffect, useState } from 'react';
import TodoList from './components/list/TodoList';
import ShowTodo from './components/showTodo/ShowTodo';
import Title from './components/title/Title';
import Item from './models/item';
import ItemService from './services/itemService';
import UpdateTodo from './components/updateTodo/UpdateTodo';
import AddIcon from '@mui/icons-material/Add'
import NewTodo from './components/newTodo/NewTodo';

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
    },
});

function App() {
    const [items, setItems] = useState<Item[]>([])
    const [showNew, setShowNew] = useState<boolean>(false)
    const [showTodo, setShowTodo] = useState<boolean>(false)
    const [showUpdate, setShowUpdate] = useState<boolean>(false)
    const [item, setItem] = useState<Item>()
    const [alert, setAlert] = useState<JSX.Element>(<></>)
    const service = new ItemService()

    const buildAlert = (severity: AlertColor, title: string, message: string) => (
        <Alert severity={severity} onClose={() => { setAlert(<></>) }}>
            <AlertTitle>{title}</AlertTitle>
            {message}
        </Alert>
    )

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
        setShowNew(false)
    }

    const openNewModal = () => {
        setShowNew(true);
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
            setShowUpdate(false)
            setShowTodo(false)
        }
    }

    const handleComplete = async (item: Item) => {
        try {
            await service.delete(item.id)
            getAllItems()
        } catch (error: any) {
            console.error(error)
            setAlert(buildAlert('error', 'Item complete fails', 'This item cannot be completed'))
        } finally {
            setItem(undefined)
        }
    }

    const handleSave = async (description: string) => {
        try {
            await service.create(new Item(- 1, description, -1))
            setAlert(buildAlert('success', 'Item added', 'Your item was added into the todo list'))
            getAllItems()
        } catch (error: any) {
            console.error(error)
            setAlert(buildAlert('error', 'Item add fails', 'This item cannot be added'))
        } finally {
            setShowNew(false)
        }
    }

    const handleReorder = async (sourceIndex: number, targetIndex: number) => {
        try {
            await service.reorder(items[sourceIndex].id, items[targetIndex].id)
            getAllItems()
        } catch (error: any) {
            console.error(error)
            setAlert(buildAlert('error', 'Item reordering fails', 'This item cannot be reordered'))
        } finally {
            setShowNew(false)
        }
    }

    return (
        <ThemeProvider theme={darkTheme}>
            <CssBaseline />
            <Title />
            {alert}
            <Fab data-testid='new_button' color="primary" style={{ position: 'absolute', bottom: '2em', right: '2em' }} onClick={openNewModal}>
                <AddIcon />
            </Fab>
            <TodoList items={items} onItemClick={handleItemClick} onComplete={handleComplete} onReorder={handleReorder} />
            {<NewTodo open={showNew} onClose={handleClose} onSaveClick={handleSave} />}
            {item && <ShowTodo open={showTodo} item={item} onClose={handleClose} onUpdateClick={openUpdateModal} />}
            {item && <UpdateTodo open={showUpdate} item={item} onClose={handleClose} onUpdateClick={handleUpdate} />}
        </ThemeProvider>
    )
}

export default App;
