package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/get", Get)
	http.HandleFunc("/post", Post)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/update", Update)
	http.ListenAndServe(":8000", nil)
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

var global *[]Person

func Get(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(global)
}

func Post(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	var (
		todo  Person
		todos []Person
	)

	json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	json.NewEncoder(rw).Encode(todo)
	global = &todos
	rw.Write([]byte(`{"message": "Successfully Posted Data"}`))
}

func Delete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	var data []Person
	data = *global

	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	for index, todo := range data {
		if todo.ID == id {
			data = append(data[:index], data[index+1:]...)
			global = &data
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(`{"message": "Success Deleting data"}`))
		}
	}
}

func Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	var (
		dataPerson []Person
		updateData Person
	)
	dataPerson = *global

	json.NewDecoder(r.Body).Decode(&updateData)
	for index, data := range dataPerson {
		if data.ID == updateData.ID {
			dataPerson = append(dataPerson[:index], dataPerson[index+1:]...)
			dataPerson = append(dataPerson, updateData)
			global = &dataPerson
		}
	}
	json.NewEncoder(rw).Encode(global)
	rw.Write([]byte(`{"message": "Successfully Updated Data"}`))
}
