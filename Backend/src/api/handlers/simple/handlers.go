package simple

import (
	"fmt"
	"net/http"
	"api/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var mainResponse utils.MainResponse
	mainResponse.Status = 200
	body := map[string]string{"I am working": "I am awesome"}
	utils.WriteJson(w, body, mainResponse.Status)

}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "Welcome!\n")
}
