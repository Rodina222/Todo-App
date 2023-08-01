<template>
  <li>
    <div v-if="editing">
      <input type="text" v-model="editedTitle" @keyup.enter="updateTodo" />
      <label @dblclick="editTitle()">
        {{ todo.title }}
      </label>
    </div>
    <div v-else>
      <label @dblclick="editTitle()">
        <input type="checkbox" v-model="editedCompleted" @change="updateTodo" />
        {{ editedTitle }}
      </label>
    </div>
    <button @click="deleteTodo()">Delete</button>
  </li>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'TodoItem',
  props: {
    todo: {
      type: Object,
      required: true,
      default: () => ({
        title: 'todo',
        completed: false
      })
    }
  },
  data() {
    return {
      editing: false as boolean,
      editedTitle: this.todo.title as string,
      editedCompleted: this.todo.completed as boolean
    }
  },
  emits: ['add', 'delete', 'update'],

  methods: {
    markCompleted: function () {
      this.editedCompleted = !this.todo.completed
      this.updateTodo()
    },

    deleteTodo: function () {
      this.$emit('delete', this.todo.id)
    },

    editTitle() {
      this.editing = true
    },

    updateTodo: function () {
      this.$emit('update', this.todo.id, this.editedTitle, this.editedCompleted)
      this.editing = false
    }
  }
})
</script>

<style></style>
