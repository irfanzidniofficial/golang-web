package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)


func TestServer(t *testing.T){
	server:= http.Server{
		Addr: "localhost:8080",

	}

	err:= server.ListenAndServe()
	if err!=nil{
		panic(err)
	}
}


func TestRequest(t *testing.T) {
	var handler http.HandlerFunc= func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, r.Method)
		fmt.Println(w, r.RequestURI)
	}

	server:=http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}
	err:=server.ListenAndServe()
	if err!=nil{
		panic(err.Error)
	}
}
