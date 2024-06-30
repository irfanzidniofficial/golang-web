package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request){
	name := request.URL.Query().Get("name")
	if name == ""{
		fmt.Fprintln(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)

	}
}

func TestQueryParameter(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Irfan", nil)
	recorder:=httptest.NewRecorder()

	SayHello(recorder, request)

	respone := recorder.Result()
	body, _:=io.ReadAll(respone.Body)

	fmt.Println(string(body))
	
}