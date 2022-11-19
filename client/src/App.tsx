import React, { useEffect, useState } from 'react';
import TodoList from './components/list/TodoList';
import Title from './components/title/Title';
import Item from './models/item';
import ItemService from './services/itemService';

function App() {
    const [items, setItems] = useState<Item[]>([])
    const service = new ItemService()

    const getAllItems = async () => {
        setItems(await service.getAll())
    }

    useEffect(() => {
        getAllItems()
    }, [])


  return <div>
    <Title />
      <TodoList items={items} />
  </div>
}

export default App;
