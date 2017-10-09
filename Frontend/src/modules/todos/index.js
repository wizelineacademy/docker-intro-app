import * as actions from './actions'
import * as getters from './getters'
import Vue from 'vue'
import sweetalert from 'sweetalert'

import {
  DELETE_TODO,
  CREATE_TODO,
  UPDATE_TODO,
  GET_TODOS,
  COMPLETE_TODO
} from './mutation-types'

// initial state
const initialState = {
  todos: []
}

// mutations
const mutations = {
  [GET_TODOS] (state, allTodos) {
    console.log('received', allTodos)
    state.todos = allTodos
  },

  [CREATE_TODO] (state, newTodo) {
    state.todos.push(newTodo)
    console.log('hit api at', process.env.API_URL)
    sweetalert('Success!', 'To-Do completed!', 'success')
  },

  [COMPLETE_TODO] (state, todo) {
    const org = state.todos.filter(x => x.ID === todo.ID)
    const todoIndex = state.todos.indexOf(org[0])
    Vue.set(state.todos, todoIndex, todo)
  },

  [UPDATE_TODO] (state, updatedtodo) {
    const org = state.todos.filter(x => x.ID === updatedtodo.ID)
    const todoIndex = state.todos.indexOf(org[0])
    console.log('updated!', state.todos.indexOf(org[0]))
    Vue.set(state.todos, todoIndex, updatedtodo)
  },

  [DELETE_TODO] (state, todo) {
    const org = state.todos.filter(x => x.ID === todo.ID)
    const todoIndex = state.todos.indexOf(org[0])
    sweetalert('Deleted!', 'Your To-Do has been deleted.', 'success')
    state.todos.splice(todoIndex, 1)
  }
}

export default {
  state: { ...initialState },
  // OBS! Don't forget to export your actions from the products module as well.
  actions,
  getters,
  mutations
}
