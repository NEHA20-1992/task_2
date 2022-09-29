package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	
)
type modelTodo struct{
	ID  string `json:"id"`
	Todos string `json:"todos"`
}
var container []modelTodo


func allTodo(w http.ResponseWriter,r *http.Request){
	fmt.Println("hit the end point")
	json.NewEncoder(w).Encode(container)
}
 
func postTodo(w http.ResponseWriter,r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
	var todos modelTodo
	json.Unmarshal(reqBody, &todos)
	container=append(container,todos)
	fmt.Println("Test post endpoint worked")
	 json.NewEncoder(w).Encode(container)
}
 
func getTodoByID(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	for  _,todo := range container{
		if todo.ID ==key{
			json.NewEncoder(w).Encode(todo)
		}
	}
}

func updateTodoByID(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var updatetodo modelTodo
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatetodo)
	for _,todo := range container{
		if todo.ID ==id{
			todo.Todos=updatetodo.Todos
			json.NewEncoder(w).Encode(todo)
		}
	}
}

func deleteTodo(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var updatetodo modelTodo
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatetodo)
	for i,todo := range container{
		if todo.ID ==id{
			container = append(container[:i], container[i+1:]...)
		}
	}
}

func main() {
            log.Println(" Listening and serving HTTP on :8909")

				handleRequests()
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/todo", allTodo).Methods("GET")
	r.HandleFunc("/posttodo",postTodo).Methods("POST")
	r.HandleFunc("/todo/{id}",getTodoByID).Methods("GET")
    r.HandleFunc("/updatetodo/{id}",updateTodoByID).Methods("PUT")
     r.HandleFunc("/deletetodo/{id}",deleteTodo).Methods("DELETE")
log.Fatal(http.ListenAndServe(":8909", r))
}


