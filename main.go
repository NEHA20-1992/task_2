package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type DataTable struct {
	ID            int `gorm:"primary_key;auto_increment" json:"id"`
	Name          string `json:"name"`
	Age           uint8  `json:"age"`
	Email         string `json:"email"`
	ContactNumber string `json:"contactNumber"`
	
}
 var db *gorm.DB
func Connect()error{
	var err error
	   db,err= gorm.Open("mysql","root"+":"+"root"+"@/"+"taskdata")
	// db,err= gorm.Open("mysql", "root:root@localhost:9090/taskdata")
	if err != nil{
		 log.Fatal(err)
		 log.Println("failed to connect database")
	}
	log.Println("Connection was successful!!")
return nil
} 

func main() {
 Connect()
 r := mux.NewRouter()
 r.HandleFunc("/data", createData).Methods("POST")
 r.HandleFunc("/fetchdata",getData).Methods("GET")
 r.HandleFunc("/fetchdata/{id}",getDataByID).Methods("GET")
 r.HandleFunc("/updatedata/{id}",updateDataByID).Methods("PUT")
 r.HandleFunc("/delete/{id}",DeleteDataById).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8909", r))
}



func createData(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var data DataTable
	json.Unmarshal(requestBody, &data)
	db.Create(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
func getData(w http.ResponseWriter, r *http.Request) {
	var data =[]DataTable{}
	db.Find(&data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getDataByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var data DataTable
	db.First(&data, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func updateDataByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var data DataTable
	json.Unmarshal(requestBody, &data)
	db.Save(&data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
	log.Println("Updated Successfully")
}

func DeleteDataById(w http.ResponseWriter,r *http.Request){
	vars :=mux.Vars(r)
	key :=vars["id"]
	var data DataTable
	id, _ := strconv.ParseInt(key, 10, 64)
	db.Where("id = ?", id).Delete(&data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	log.Println("data deleted Successfully")
}