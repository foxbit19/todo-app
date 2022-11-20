import { Box, createTheme, CssBaseline, Grid, Paper, styled, ThemeProvider } from '@mui/material';
import React, { useEffect, useState } from 'react';
import TodoList from './components/list/TodoList';
import ShowTodo from './components/showTodo/ShowTodo';
import Title from './components/title/Title';
import Item from './models/item';
import ItemService from './services/itemService';

function App() {
    const [items, setItems] = useState<Item[]>([])
    const [showTodo, setShowTodo] = useState<boolean>(false)
    const [item, setItem] = useState<Item>()
    const service = new ItemService()

    const getAllItems = async () => {
        setItems(await service.getAll())
    }

    useEffect(() => {
        getAllItems()
    }, [])

    const darkTheme = createTheme({
        palette: {
            mode: 'dark',
        },
    });

    const Item = styled(Paper)(({ theme }) => ({
        backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
        ...theme.typography.body2,
        padding: theme.spacing(1),
        margin: '10%',
        textAlign: 'center',
        color: theme.palette.text.secondary,
    }));

    const handleItemClick = (item: Item) => {
        setItem(item)
        setShowTodo(true)
    }

    const handleClose = () => {
        setShowTodo(false)
    }

    return (
        <ThemeProvider theme={darkTheme}>
            <CssBaseline />
            <Box sx={{ flexGrow: 1 }}>
                <Title />
                <Item><TodoList items={items} onItemClick={handleItemClick} /></Item>
            </Box>
            {item && <ShowTodo open={showTodo} item={item} onClose={handleClose} />}
        </ThemeProvider>
    )
}

export default App;
