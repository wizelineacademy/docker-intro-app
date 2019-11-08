<template>
  <div>
    <p class="tasks">Completed Tasks: {{todos.filter(todo => {return todo.done === true}).length}}</p>
    <p class="tasks">Pending Tasks: {{todos.filter(todo => {return todo.done === false}).length}}</p>
    <todo v-on:delete-todo="deleteItem" v-on:complete-todo="completeItem" v-on:update-todo="updateItem" v-for="todo in todos" :todo.sync="todo" :key="todo.ID"></todo>
  </div>
</template>

<script type = "text/javascript" >
import sweetalert from 'sweetalert'
import Todo from './Todo'
import { mapActions } from 'vuex'

export default {
  props: ['todos'],
  components: {
    Todo
  },
  methods: {
    ...mapActions([
      'deleteTodo',
      'updateTodo',
      'completeTodo',
      'getAllTodos'
    ]),
    deleteItem (todo) {
      sweetalert({
        title: 'Are you sure?',
        text: 'This To-Do will be permanently deleted!',
        type: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#DD6B55',
        confirmButtonText: 'Yes, delete it!',
        closeOnConfirm: false
      })
        .then((willDelete) => {
        // have API do a post to create the todo
          if (willDelete) {
            this.deleteTodo(todo)
          }
        })
    },
    updateItem (updatedtodo) {
      this.updateTodo(updatedtodo)
    },
    completeItem (todo) {
      // have API do a post to complete the todo
      this.completeTodo(todo)
    }
  }
}
</script>

<style scoped>
p.tasks {
  text-align: center;
}
</style>

