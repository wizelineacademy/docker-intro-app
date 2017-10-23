package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Title   string        `json:"title"`
	Project string        `json:"project"`
	Done    bool          `json:"done"`
}

func (todo *Todo) ValidateStruct() error {
	if todo.Title == "" {
		return errors.New("Body Missing title")
	}
	if todo.Project == "" {
		return errors.New("Body Missing project")
	}
	return nil
}
