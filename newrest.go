package main

import (
	"fmt"
	"net/http"

	"github.com/connect2naga/Training/pkg"

	"github.com/gorilla/mux"
)
type APIHandler struct {
	odd chan(int)
	even chan(int)
}

func (a *APIHandler) SayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := r.URL.Query()
	w.Write([]byte(fmt.Sprintf("<h1>%s from %s</h1>",vars["name"], value["location"])))
}


func (a *APIHandler) Print(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%s",vars["intValue"])
	//val  := 0 
	val, _ := strconv.Atoi(vars["intValue"])
    fmt.Printf("CONV STRING:%v",val)
	switch val %2 {
	case 0:
		 a.even <- val
	case 1:
		a.odd <- val
	default :
	}

}


func main() {
	a := pkg.APIHandler{}

	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", a.SayHello)
	r.HandleFunc("/number/{intValue}", a.Print)
	fmt.Printf("Server started at :8080....")
	http.ListenAndServe(":8080", r)

}