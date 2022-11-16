package main
import (
    "fmt"
    "log"
    "net/http"
)
type Car struct {
    Name    string `json:"name"`
    Color string `json:"color"`
    Model string `json:"model"`
}
type Carinfo []Car
func view(w http.ResponseWriter, r *http.Request) {
    a := Carinfo{
        Car{Name: "BMW", Color: "Black", Model: "xyz"},
    }
    fmt.Println("Endpoint hit:All articles endpoint")
    fmt.Fprintf(w, "%s", a)
}
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome")
}
func handleRequest() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/names", view)
    log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
    handleRequest()
}