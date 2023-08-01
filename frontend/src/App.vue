<template>
  <div id="app">
    <HelloToDo msg="✍️ ToDo Application" />
    <TodoList
      :todos="todos"
      @add="addTodo"
      @delete="deleteTodo"
      @update="updateTodo"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

import HelloToDo from './components/Header.vue'
import TodoList from './components/TodoList.vue'

export default defineComponent({
  name: 'App',
  components: {
    HelloToDo,
    TodoList
  },
  data() {
    return {
      todos: [] as { id: number; title: string; completed: boolean }[],
      error: '' as string,
      baseurl: 'http://localhost:8096' as string
    }
  },
  beforeMount() {
    this.getTodos()
  },

  methods: {
    addTodo(title: string) {
      console.log('new todo', title)
      const newTodo = { title: title, completed: false }

      console.log('todos', this.todos)

      // Make a POST request to the backend API to add the new todo item
      fetch(this.baseurl + '/todos', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json; charset=utf-8',
          Accept: 'application/json'
        },
        body: JSON.stringify(newTodo)
      })
        .then((response) => response.json())
        .then((data) => {
          // Add the new todo item to the local state
          this.todos.push({
            id: data.id,
            title: data.title,
            completed: data.completed
          })
        })
        .catch((error) => {
          console.error('Error adding todo:', error)
        })
    },

    updateTodo(id: number, title: string, completed: boolean) {
      console.log('updated todo', title, completed)

      const updatedTodo = { id: id, title: title, completed: completed }

      // Make a POST request to the backend API to add the new todo item
      fetch(this.baseurl + `/todos/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json; charset=utf-8',
          Accept: 'application/json'
        },
        body: JSON.stringify(updatedTodo)
      })
        .then((response) => response.json())

        .catch((error) => {
          console.error('Error updating todo:', error)
        })

      this.getTodos()

      console.log('todos', this.todos)
    },
    deleteTodo(id: number) {
      const todoIndex = this.todos.findIndex((todo) => todo.id === id)
      if (todoIndex >= 0) {
        this.todos.splice(todoIndex, 1)
        fetch(this.baseurl + `/todos/${id}`, {
          method: 'DELETE'
        })
          .then((response) => response.json())
          .then((data) => {
            console.log('Todo item deleted:', data)
          })
          .catch((error) => {
            console.error('Error deleting todo:', error)
          })
      }
    },

    getTodos() {
      fetch(this.baseurl + '/todos')
        .then((response) => response.json())
        .then((data) => {
          this.todos = data
        })
        .catch((error) => {
          console.error('Error fetching todos:', error)
        })
    },

    getTodoById(id: number) {
      fetch(this.baseurl + `/todos/${id}`)
        .then((response) => response.json())
        .then((data) => {
          console.log('Todo item fetched:', data)
          // Do something with the fetched todo item, such as updating a component's data property
        })
        .catch((error) => {
          console.error(`Error fetching todo with ID ${id}:`, error)
        })
    }
  },

  mounted() {
    this.getTodos()
  }
})
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  max-width: 2000px;
  margin: 0 auto;
  padding: 2rem;
  color-scheme: light dark;
  color: rgba(255, 255, 255, 0.87);
  background-color: #242424;
  height: 100vh;
}

TodoList {
  border: 1px solid #ccc;
  align-items: flex-start;
  padding: 20px;
  margin-bottom: 10px;
  margin-top: 8px;
  background-color: darkgray;
  max-width: 60%;
  text-align: center;
  width: 100%;
  color: black;
  font-weight: bold;
  margin-right: 30px;
}
</style>
