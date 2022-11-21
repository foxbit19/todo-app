import { Alert, AlertColor, AlertTitle, createTheme, CssBaseline, Fab, ThemeProvider } from '@mui/material';
import React, { useEffect, useState } from 'react';
import TodoList from './components/list/TodoList';
import Item from './models/item';
import ItemService from './services/itemService';
import UpdateTodo from './components/updateTodo/UpdateTodo';
import AddIcon from '@mui/icons-material/Add'
import NewTodo from './components/newTodo/NewTodo';
import AppContainer from './components/container/AppContainer';
import DeleteTodo from './components/deleteTodo/DeleteTodo';
import Footer from './components/footer/Footer';
import Appbar from './components/appbar/Appbar';
import CompletedTodos from './components/completedTodos/CompletedTodos';
import CompletedItemService from './services/completedItemService';

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
        primary: {
            main: '#EF1B53'
        },
    },
});

function App() {
    const [items, setItems] = useState<Item[]>([])
    const [completedItems, setCompletedItems] = useState<Item[]>([])
    const [showNew, setShowNew] = useState<boolean>(false)
    const [showUpdate, setShowUpdate] = useState<boolean>(false)
    const [showDelete, setShowDelete] = useState<boolean>(false)
    const [showCompleted, setShowCompleted] = useState<boolean>(false)
    const [item, setItem] = useState<Item>()
    const [alert, setAlert] = useState<JSX.Element>(<></>)
    const itemService = new ItemService()
    const completedItemService = new CompletedItemService()

    const buildAlert = (severity: AlertColor, title: string, message: string) => (
        <Alert severity={severity} onClose={() => { setAlert(<></>) }}>
            <AlertTitle>{title}</AlertTitle>
            {message}
        </Alert>
    )

    const getAllItems = async () => {
        setItems(await itemService.getAll())
    }

    useEffect(() => {
        getAllItems()
    }, [])

    const handleClose = () => {
        setShowDelete(false)
        setShowUpdate(false)
        setShowNew(false)
        setShowCompleted(false)
    }

    const openNewModal = () => {
        setShowNew(true);
    }

    const openUpdateModal = async (item: Item) => {
        setItem(item)
        setShowUpdate(true)
    }

    const openDeleteModal = async (item: Item) => {
        setItem(item)
        setShowDelete(true)
    }

    const openCompletedModal = async () => {
        try {
            setCompletedItems(await completedItemService.getAll())
            setShowCompleted(true)
        } catch (error) {
            console.error(error)
            setAlert(buildAlert('error', 'Show completed fails', 'An error occurs during the download of the completed items'))
        }
    }

    const handleUpdate = async (item: Item) => {
        try {
            await itemService.update(item)
            setAlert(buildAlert('success', 'Item updated', 'Your item was updated successfully'))
            getAllItems()
        } catch (error: any) {
            console.error(error)
            setAlert(buildAlert('error', 'Item update fails', 'This item cannot be updated'))
        } finally {
            // clear the current item
            setItem(undefined)
            setShowUpdate(false)
        }
    }

    const handleDelete = async (item: Item) => {
        try {
            await itemService.delete(item.id)
            setAlert(buildAlert('success', 'Item deleted', 'Your item was deleted successfully'))
            getAllItems()
        } catch (error: any) {
            console.error(error)
            setAlert(buildAlert('error', 'Item delete fails', 'This item cannot be delete'))
        } finally {
            setItem(undefined)
            setShowDelete(false)
        }
    }

    const handleSave = async (description: string) => {
        try {
            await itemService.create(new Item(- 1, description, -1))
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
            await itemService.reorder(items[sourceIndex].id, items[targetIndex].id)
            getAllItems()
        } catch (error: any) {
            console.error(error)
            setAlert(buildAlert('error', 'Item reordering fails', 'This item cannot be reordered'))
        } finally {
            setShowNew(false)
        }
    }

    const handleComplete = async (item: Item) => {
        try {
            item.completed = true
            item.completedDate = new Date()
            await itemService.update(item)
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
            <Appbar onShowCompletedClick={openCompletedModal} />
            <AppContainer>
                <div>
                    {alert}
                    <Fab data-testid='new_button' color="primary" style={{ position: 'fixed', bottom: '2em', right: '2em' }} onClick={openNewModal}>
                        <AddIcon />
                    </Fab>
                    <TodoList
                        items={items}
                        onUpdate={openUpdateModal}
                        onDelete={openDeleteModal}
                        onComplete={handleComplete}
                        onReorder={handleReorder}
                    />
                    <Footer />
                    {<NewTodo open={showNew} onClose={handleClose} onSaveClick={handleSave} />}
                    {item && <UpdateTodo open={showUpdate} item={item} onClose={handleClose} onUpdateClick={handleUpdate} />}
                    {item && <DeleteTodo open={showDelete} item={item} onClose={handleClose} onDeleteClick={handleDelete} />}
                    <CompletedTodos items={completedItems} open={showCompleted} onClose={handleClose} />
                </div>
            </AppContainer>
        </ThemeProvider>
    )
}

export default App;
