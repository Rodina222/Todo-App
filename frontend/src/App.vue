<template>
  <div id="app">
    <HelloToDo msg="✍️ ToDo Application" />
    <TaskList
      :tasks="tasks"
      @add="addTask"
      @delete="deleteTask"
      @update="updateTask"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import axios from 'axios'

import HelloToDo from './components/Header.vue'
import TaskList from './components/TodoList.vue'
import { TaskType } from './components/TodoItem.vue'

export default defineComponent({
  name: 'App',
  components: {
    HelloToDo,
    TaskList
  },

  data() {
    return {
      tasks: [] as TaskType[],
      error: '' as string,
      //baseurl: 'http://localhost:8096' as string
      baseurl: this.$API_BASE_URL
    }
  },
  beforeMount() {
    this.getTasks()
  },

  methods: {
    async addTask(title: string) {
      const newTask = { title: title, completed: false }

      try {
        const response = await axios.post(this.baseurl + '/todos', newTask, {
          headers: {
            'Content-Type': 'application/json; charset=utf-8',
            Accept: 'application/json'
          }
        })

        const data = response.data
        console.debug(data)
        this.tasks.push({
          id: data.id,
          title: data.title,
          completed: data.completed
        })
      } catch (error) {
        console.error('Error adding task:', error)
      }
    },

    async updateTask(id: number, title: string, completed: boolean) {
      const updatedTodo = { id: id, title: title, completed: completed }

      try {
        const response = await axios.put(
          this.baseurl + `/todos/${id}`,
          updatedTodo,
          {
            headers: {
              'Content-Type': 'application/json; charset=utf-8',
              Accept: 'application/json'
            }
          }
        )
        const data = response.data
        console.debug(data)
      } catch (error) {
        console.error('Error updating task:', error)
      }

      this.getTasks()
      console.log('todos', this.tasks)
    },

    async deleteTask(id: number) {
      const todoIndex = this.tasks.findIndex((todo) => todo.id === id)
      if (todoIndex >= 0) {
        this.tasks.splice(todoIndex, 1)

        try {
          const response = await axios.delete(this.baseurl + `/todos/${id}`)
          const data = response.data
          console.debug(data)
        } catch (error) {
          console.error('Error deleting todo:', error)
        }
      }
    },

    async getTasks() {
      console.log('baseurl', this.baseurl)
      try {
        const response = await axios.get(this.baseurl + '/todos')
        this.tasks = response.data
      } catch (error) {
        console.error('Error fetching todos:', error)
      }
    },

    async getTaskById(id: number) {
      try {
        const response = await axios.get(this.baseurl + `/todos/${id}`)
        const data = response.data
        console.debug(data)
      } catch (error) {
        console.error(`Error fetching todo with ID ${id}:`, error)
      }
    }
  },

  mounted() {
    this.getTasks()
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
