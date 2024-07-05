package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empty")

	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	respone := recorder.Result()

	body, _ := io.ReadAll(respone.Body)
	fmt.Println(respone.StatusCode)

	fmt.Println(respone.Status)

	fmt.Println(string(body))
}


func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080/?name=Irfan", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	respone := recorder.Result()

	body, _ := io.ReadAll(respone.Body)
	fmt.Println(respone.StatusCode)

	fmt.Println(respone.Status)

	fmt.Println(string(body))
}
