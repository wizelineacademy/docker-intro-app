package handlers

import (
	"fmt"
	"net/http"
)

func Read(w http.ResponseWriter, r *http.Request) {

	fmt.Println("I'm reading")
	w.WriteHeader(http.StatusCreated)
}
