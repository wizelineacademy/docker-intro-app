package main

import (
	"net/http"
	"api"
  "api/utils"
	"log"
)

func handleCount(w http.ResponseWriter, r *http.Request) {

}

func main() {
	log.Println("Starting Server")
  utils.LoadConfig("config/config.yml")
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":"+utils.CONFIG.Port, router))
}
