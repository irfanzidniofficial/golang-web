package golang_web

import (
	"fmt"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {

	contentType:= request.Header.Get("content-type") 
	fmt.Fprintln(writer, contentType)
}

func TestRequestHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("content-type", "application/json")

	recorder:=httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request){
	writer.Header().Add("X-Powered-By","Irfan Zidni")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("content-type", "application/json")

	recorder:=httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}