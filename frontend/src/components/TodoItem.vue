<template>
  <li>
    <div v-if="editing">
      <input type="text" v-model="editedTitle" @keyup.enter="updateTask" />
      <label @dblclick="editTitle">
        {{ task.title }}
      </label>
    </div>
    <div v-else>
      <label @dblclick="editTitle">
        <input type="checkbox" v-model="editedCompleted" @change="updateTask" />
        <span :class="{ completed: editedCompleted }">{{ editedTitle }}</span>
      </label>
    </div>
    <button @click="deleteTask">Delete</button>
  </li>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export type TaskType = {
  id: number
  title: string
  completed: boolean
}

export enum TodoItemEvents {
  DELETE_TASK = 'delete',
  UPDATE_TASK = 'update',
  EDIT_TITLE = 'editTitle'
}

export default defineComponent({
  name: 'TodoItem',
  props: {
    task: {
      type: Object as () => TaskType,
      required: true,
      default: () => ({
        title: 'task',
        completed: false
      })
    }
  },
  data() {
    return {
      editing: false as boolean,
      editedTitle: this.task.title as string,
      editedCompleted: this.task.completed as boolean
    }
  },
  emits: ['add', 'delete', 'update'],

  methods: {
    markCompleted: function () {
      this.editedCompleted = !this.task.completed
      this.updateTask()
    },

    deleteTask: function () {
      this.$emit(TodoItemEvents.DELETE_TASK, this.task.id)
    },

    editTitle() {
      this.editing = true
    },

    updateTask: function () {
      this.$emit(
        TodoItemEvents.UPDATE_TASK,
        this.task.id,
        this.editedTitle,
        this.editedCompleted
      )
      this.editing = false
    }
  }
})
</script>

<style>
.completed {
  text-decoration: line-through;
}
</style>
