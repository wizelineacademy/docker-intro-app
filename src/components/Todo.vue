<template>
  <div class='ui centered card'>
    <div class="content" v-show="!isEditing">
      <div class='header'>
          {{ todo.title }}
      </div>
      <div class='meta'>
          {{ todo.project }}
      </div>
      <div class='extra content'>
          <span class='right floated edit icon' v-on:click="showForm">
          <i class='edit icon'></i>
        </span>
        <span class='right floated trash icon' v-on:click="deleteItem(todo)">
          <i class='trash icon'></i>
        </span>
      </div>
    </div>
    <div class="content" v-show="isEditing">
      <div class='ui form'>
        <div class='field'>
          <label>Title</label>
          <input type='text' v-model="title" >
        </div>
        <div class='field'>
          <label>Project</label>
          <input type='text' v-model="project" >
        </div>
        <div class='ui fluid buttons'>
          <button class='ui basic red button' v-on:click="hideForm">
            Cancel
          </button>
          <div class="or"></div>
          <button class='ui basic green button' v-on:click="saveItem(todo,title, project)">
            Save 
          </button>
        </div>
      </div>
    </div>
    <div class='ui bottom attached green basic button' v-show="!isEditing &&todo.done" disabled>
        Completed
    </div>
    <div class='ui bottom attached red basic button' v-on:click="completeItem(todo)" v-show="!isEditing && !todo.done">
        Pending
    </div>
  </div>
</template>

<script type="text/javascript">
  export default {
    props: ['todo'],
    data () {
      console.log(this)
      return {
        isEditing: false,
        title: this.todo.title,
        project: this.todo.project
      }
    },
    methods: {
      completeItem (todo) {
        this.$emit('complete-todo', todo)
      },
      deleteItem (todo) {
        this.$emit('delete-todo', todo)
      },
      saveItem (todo, title, project) {
        this.isEditing = false
        const updateTodo = {ID: todo.ID, title, project, done: todo.done}
        this.$emit('update-todo', updateTodo)
      },
      showForm () {
        this.isEditing = true
      },
      hideForm () {
        console.log('cancel ')
        this.isEditing = false
      }
    }
  }
</script>
