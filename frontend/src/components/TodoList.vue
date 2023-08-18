<template>
  <div class="TodoList">
    <input
      type="text"
      class="custom-input"
      placeholder="Add a new todo"
      v-model="title"
      @keyup.enter="addTask"
    />
    <button class="btn btn-primary" type="button" @click="addTask">Add</button>
    <ul>
      <todo-item
        v-for="task in tasks"
        :key="task.id"
        :task="task"
        class="todoItem"
        @delete="deleteTask"
        @update="updateTask"
      ></todo-item>
    </ul>
    <h2 v-show="tasks.length === 0">No Todos Here😞</h2>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import TodoItem, { TaskType } from './TodoItem.vue'

export default defineComponent({
  name: 'TodoList',
  components: {
    TodoItem
  },

  props: {
    tasks: {
      type: Array as () => TaskType[],
      required: true,
      default: () => [] // Default value should be an empty array, not an object
    },
    baseurl: String
  },
  data() {
    return {
      title: '' as string,
      editingIndex: -1,
      editText: ''
    }
  },
  emits: ['add', 'delete', 'markCompleted', 'update'],
  methods: {
    addTask() {
      /*  this.todos.push({ title: this.newTodo, completed: false })*/

      this.$emit('add', this.title)
      this.title = ''
    },
    deleteTask: function (id: number) {
      this.$emit('delete', id)
    },
    markCompleted(index: number) {
      this.$emit('markCompleted', index)
    },
    updateTask(id: number, title: string, completed: boolean) {
      this.$emit('update', id, title, completed)
    }
  }
})
</script>

<style>
.custom-input {
  height: 25px;
  width: 300px;
  font-size: 20px;
  padding: 10px;
  margin-bottom: 5px;
  border: none;
  background-color: white;
  margin-left: 25px;
}

.btn-primary {
  background-color: yellowgreen;
  padding: 10px 20px;
  height: 48px;
  font-weight: bold;
  margin-bottom: 20px;
  margin-top: 1rem;
}

.todoList {
  max-width: 600px;
  margin: auto;
  padding-top: 50px;
}

.todoCheckbox {
  margin-right: 10px;
  transform: scale(1.5);
}

ul {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  margin: 0;
  padding: 0;
  list-style: none;
}

.todoItem {
  border: 1px solid #ccc;
  align-items: flex-start;
  padding: 10px; /* Decrease the padding to reduce the height */
  margin-bottom: 10px;
  margin-top: 8px;
  background-color: darkgray;
  max-width: 60%;
  text-align: center;
  width: 100%;
  color: black;
  font-weight: bold;
  margin-right: 30px;
  display: flex; /* Add display:flex to align the grey box with the checkbox and delete button */
}

.todoItem button {
  margin-left: auto; /* Use margin-left:auto to align the Delete button to the right */
  background-color: orangered;
  color: white;
}

.todoCheckbox:checked ~ .todoText {
  text-decoration: line-through;
}

.todoText.editing {
  display: none;
}

input[type='text'].form-control {
  width: 80%;
  height: auto;
  resize: none;
  box-sizing: border-box;
}
</style>
