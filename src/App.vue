<template>
  <div id="app">
    <h1 class="ui dividing centered header">Vue.js Todo App</h1>
    <div class='ui three column centered grid'>
      <div class='column'>
        <todo-list v-bind:todos="todos"></todo-list>
        <create-todo v-on:create-todo="createItem"></create-todo>
      </div>
    </div>
  </div>
</template>

<script>
import sweetalert from 'sweetalert'
import TodoList from './components/TodoList'
import CreateTodo from './components/CreateTodo'
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'app',
  components: {
    TodoList,
    CreateTodo
  },
  computed: mapGetters({
    todos: 'getTodos'
  }),
  mounted: function () {
    this.getAllTodos()
  },
  methods: {
    ...mapActions([
      'createTodo',
      'getAllTodos'
    ]),
    // API put to create a new todo Item
    createItem (newTodo) {
      this.createTodo(newTodo)
      sweetalert('Success!', 'To-Do created!', 'success')
    }
  }
}
</script>
