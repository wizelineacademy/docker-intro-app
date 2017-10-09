import {
  DELETE_TODO,
  CREATE_TODO,
  UPDATE_TODO,
  GET_TODOS,
  COMPLETE_TODO
} from './mutation-types'

import api from '../../services/api'

export function getAllTodos ({ commit, state }) {
  api.get('/todos')
  .then((res) => { commit(GET_TODOS, res.data.Response) })
  .catch((err) => {
    console.log('trigger error alert', err)
  })
}

export function createTodo ({ commit, state }, newTodo) {
  api.put('/todos', newTodo)
  .then((res) => { commit(CREATE_TODO, res.data.Response) })
  .catch((err) => {
    console.log('trigger error alert', err)
  })
}

export function updateTodo ({ commit, state }, updatedTodo) {
  api.post(`/todos/${updatedTodo.ID}`, updatedTodo)
  .then((res) => { commit(UPDATE_TODO, updatedTodo) })
  .catch((err) => {
    console.log('trigger error alert', err)
  })
}

export function completeTodo ({ commit, state }, todo) {
  api.post(`/todos/${todo.ID}`, todo)
  .then((res) => { commit(COMPLETE_TODO, todo); console.log(res.data.Response) })
  .catch((err) => {
    console.log('trigger error alert', err)
  })
}

export function deleteTodo ({ commit, state }, todo) {
  api.delete(`/todos/${todo.ID}`, todo)
  .then((res) => { commit(DELETE_TODO, todo); console.log(res.data.Response) })
  .catch((err) => {
    console.log('trigger error alert', err)
  })
}
