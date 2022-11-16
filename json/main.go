package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var odd = make(chan int)
var even = make(chan int)

type EmployeeDetails struct {
	Name string `json:"name"`
	//	Age  int    `json:"age"`
	//	Position string `json:"position"`
	//	Salary   uint64 `json:"salary"`
}

type APIHandler struct {
	odd  chan (int)
	even chan (int)
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{odd: make(chan int), even: make(chan int)}
}

func (a *APIHandler) SayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := r.URL.Query()
	w.Write([]byte(fmt.Sprintf("<h1>%s from %s</h1>", vars["name"], value["location"])))
}

func (a *APIHandler) Print(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%s", vars["intValue"])
	val, _ := strconv.Atoi(vars["intValue"])
	switch val % 2 {
	case 0:
		even <- val
	case 1:
		odd <- val
	default:
	}

}

func (a *APIHandler) ReadfromChannel() {
	for {

		select {
		case <-odd:
			fmt.Println("given number is odd")
		case <-even:
			fmt.Println("given number is even")
		}
	}

}

func (e *EmployeeDetails) PrintEmployeeDetails(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	//var val EmployeeDetails
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	//fmt.Printf("--------> %s", string(bodyBytes))
	var emp EmployeeDetails
	er := json.Unmarshal(bodyBytes, &emp)
	//fmt.Println(res)
	//err := json.Unmarshal(Data, &country1)

	if er != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	// printing details of
	// decoded data
	fmt.Println("Struct is:", emp)
	json.NewEncoder(w).Encode(emp)

	//	res := json.Unmarshal(bodyBytes, &val)
	//	fmt.Print(res)

}
func main() {
	a := NewAPIHandler()
	s := EmployeeDetails{}
	go a.ReadfromChannel()

	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", a.SayHello)
	r.HandleFunc("/number/{intValue}", a.Print)

	r.HandleFunc("/employee", s.PrintEmployeeDetails)

	fmt.Printf("Server started at :8080....")
	http.ListenAndServe(":8080", r)

}
