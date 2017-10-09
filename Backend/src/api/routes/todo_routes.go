package routes

import (
	"api/handlers/todo"
)

var TodoRoutes = Routes{
	Route{
		"Read All Todos",
		"Get",
		"/todos",
		todo_handlers.GetTodos,
	},
	Route{
		"Create Todo in DB",
		"Put",
		"/todos",
		todo_handlers.CreateTodo,
	},
	Route{
		"Delete Todo from db by ID",
		"Delete",
		"/todos/{id}",
		todo_handlers.DeleteTodo,
	},
	Route{
		"Update a Todo in the db by OD",
		"POST",
		"/todos/{id}",
		todo_handlers.UpdateTodo,
	},
	Route{
		"Allow pre-flight",
		"OPTIONS",
		"/todos/{id}",
		todo_handlers.OptionsTodo,
	},
	Route{
		"Allow pre-flight",
		"OPTIONS",
		"/todos",
		todo_handlers.OptionsTodo,
	},
	
}
