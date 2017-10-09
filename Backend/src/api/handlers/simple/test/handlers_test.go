package simple_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"api"
	"api/utils"
	"testing"
)

var (
	server *httptest.Server
	// reader   io.Reader //Ignore this for now
	// usersUrl string
)

func init() {
	utils.LoadConfig("../../../../config/config.yml")
	router := api.NewRouter()
	server = httptest.NewServer(router) //Creating new server

}

func TestIndex(t *testing.T) {
	utils.LoadConfig("../../../../config/config.yml")
	request, err := http.NewRequest("GET", server.URL+"/", nil)
	res, err := http.DefaultClient.Do(request)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	// Check the response body is what we expect.

	if result.(map[string]interface{})["I am working"] == nil {
		t.Errorf("handler returned unexpected body: got %v ", result)
	}
	expected := `I am awesome`
	if result.(map[string]interface{})["I am working"].(string) != expected {
		t.Errorf("handler returned unexpected value: got %v want %v", result.(map[string]interface{})["I am working"].(string), expected)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}

}
