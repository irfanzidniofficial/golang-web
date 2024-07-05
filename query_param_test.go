package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Irfan&last_name=Zidni", nil)
	recorder:=httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	respone := recorder.Result()
	body, _:=io.ReadAll(respone.Body)

	fmt.Println(string(body))	

}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request){
	query := request.URL.Query()
	// bisa menangkap lebih dari 1 parameter, kalau Query Get hanya bisa menangkap 1 parameter
	names := query["name"]
	fmt.Println(writer, strings.Join(names, " "))
}

func TestMultipleQueryParameterValues(t *testing.T){
	request:= httptest.NewRequest(http.MethodGet, "ttp://localhost:8080/hello?name=Irfan&name=Zidni&name=Muhammad", nil)
	recorder:= httptest.NewRecorder()


	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))


}