import React from 'react';
import TodoList from './components/list/TodoList';
import Title from './components/title/Title';

function App() {
  return <div>
    <Title />
    <TodoList items={[]} />
  </div>
}

export default App;
