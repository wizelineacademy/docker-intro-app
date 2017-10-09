package todo_handlers

import (
	"fmt"
	"net/http"
	"api/utils"
	"encoding/json"
	"api/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)


func GetTodos(w http.ResponseWriter, r *http.Request) {
	var mainResponse utils.MainResponse
	mainResponse.Status = 200
	session,_ :=utils.GetDB()
	c := session.DB("test").C("Todos")
	var results []models.Todo
	var rr []interface{}
	err := c.Find(bson.M{}).All(&results)
	c.Find(bson.M{}).All(&rr)
		if err != nil {
			mainResponse.Status = 500
			fmt.Println(err)
			mainResponse.Message = "Error"
			utils.WriteJson(w, mainResponse, mainResponse.Status)
			return
		}
	// if no results return an empty slice insted of nil
	if results == nil {
		results = make([]models.Todo, 0)
	}
	mainResponse.Response = results
	
	utils.WriteJson(w, mainResponse, mainResponse.Status)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var mainResponse utils.MainResponse
	mainResponse.Status = 200
	var newTodo models.Todo
	//Parse json into golang object
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newTodo)
	err = newTodo.ValidateStruct()
	if err != nil {
		fmt.Printf("Cannot decode json body %s", err.Error())
		mainResponse.Status = http.StatusBadRequest
		mainResponse.Message = "Cannot create todo relation -" + err.Error()
		utils.WriteJson(w, mainResponse ,mainResponse.Status)
		return
	}
	newTodo.ID = bson.NewObjectId()
	session,_ :=utils.GetDB()
	c := session.DB("test").C("Todos")
	// Insert Datas
	err = c.Insert(newTodo)
		if err != nil {
			mainResponse.Status = 500
			fmt.Println(err)
			mainResponse.Message = "Error"
			utils.WriteJson(w, mainResponse, mainResponse.Status)
			return
		}

	mainResponse.Response = newTodo
	
	utils.WriteJson(w, mainResponse, mainResponse.Status)
}


func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var mainResponse utils.MainResponse
	mainResponse.Status = 200
	//Get id from path params
	vars := mux.Vars(r)
	if vars["id"]== "" {
		fmt.Printf("No ID in path")
		mainResponse.Status = http.StatusBadRequest
		mainResponse.Message = "No ID in path"
		utils.WriteJson(w, mainResponse ,mainResponse.Status)
		return
	}

	session,_ :=utils.GetDB()
	c := session.DB("test").C("Todos")
	// Delete Record
	if !bson.IsObjectIdHex(vars["id"]) {
		mainResponse.Status = http.StatusBadRequest
		fmt.Println(vars["id"])
		mainResponse.Message = "ID is invalid - "+ vars["id"]
		utils.WriteJson(w, mainResponse, mainResponse.Status)
		return
	}

	err := c. RemoveId(bson.ObjectIdHex(vars["id"]))
		if err != nil {
			mainResponse.Status = 123
			fmt.Println(err,vars["id"])
			mainResponse.Message = err.Error()
			utils.WriteJson(w, mainResponse, mainResponse.Status)
			return
		}

	mainResponse.Response = "deleted " +  vars["id"]
		
	utils.WriteJson(w, mainResponse, mainResponse.Status)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var mainResponse utils.MainResponse
	mainResponse.Status = 200
	//Get id from path params
	vars := mux.Vars(r)
	if vars["id"]== "" {
		fmt.Printf("No ID in path")
		mainResponse.Status = http.StatusBadRequest
		mainResponse.Message = "No ID in path"
		utils.WriteJson(w, mainResponse ,mainResponse.Status)
		return
	}

	var newTodo models.Todo
	//Parse json into golang object
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newTodo)
	err = newTodo.ValidateStruct()
	if err != nil {
		fmt.Printf("Cannot decode json body %s", err.Error())
		mainResponse.Status = http.StatusBadRequest
		mainResponse.Message = "Cannot create todo relation - " + err.Error()
		utils.WriteJson(w, mainResponse ,mainResponse.Status)
		return
	}

	session,_ :=utils.GetDB()
	c := session.DB("test").C("Todos")
	// Update Record
	if bson.IsObjectIdHex(vars["id"]) {
		newTodo.ID = bson.ObjectIdHex(vars["id"])
	}else{
		mainResponse.Status = http.StatusBadRequest
		fmt.Println(err,vars["id"])
		mainResponse.Message = "ID is invalid - "+ vars["id"]
		utils.WriteJson(w, mainResponse, mainResponse.Status)
		return
	}

	colQuerier := bson.M{"_id": newTodo.ID  }
	change := bson.M{"$set": newTodo}
	err = c.Update(colQuerier, change)
		if err != nil {
			mainResponse.Status = 123
			fmt.Println(err,vars["id"])
			mainResponse.Message = err.Error()
			utils.WriteJson(w, mainResponse, mainResponse.Status)
			return
		}

	mainResponse.Message = "Updated " +  vars["id"]
	mainResponse.Response =  newTodo
	utils.WriteJson(w, mainResponse, mainResponse.Status)
}



func OptionsTodo(w http.ResponseWriter, r *http.Request) {
	var mainResponse utils.MainResponse
	mainResponse.Status = 200
	utils.WriteJson(w, mainResponse, mainResponse.Status)
}

