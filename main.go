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
	if global == nil || len(*global) == 0 {
		result := map[string]interface{}{
			"status":  http.StatusOK,
			"data":    nil,
			"message": "No Data Available",
		}
		json.NewEncoder(rw).Encode(result)
	} else {
		result := map[string]interface{}{
			"status":  http.StatusOK,
			"data":    global,
			"message": "Successfully Get Data",
		}
		json.NewEncoder(rw).Encode(result)
	}

}

func Post(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	var (
		input       Person
		todos, data []Person
	)

	json.NewDecoder(r.Body).Decode(&input)
	if global == nil {
		input.ID = 1
		todos = append(todos, input)
		global = &todos
		result := map[string]interface{}{
			"status":  http.StatusOK,
			"data":    global,
			"message": "Successfully Posted Data",
		}
		json.NewEncoder(rw).Encode(result)
	} else {
		data = *global
		input.ID = len(*global) + 1
		data = append(data, input)
		global = &data
		result := map[string]interface{}{
			"status":  http.StatusOK,
			"data":    global,
			"message": "Successfully Posted Data",
		}
		json.NewEncoder(rw).Encode(result)
	}
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
			result := map[string]interface{}{
				"status":  http.StatusOK,
				"data":    global,
				"message": "Successfully Deleted Data",
			}
			json.NewEncoder(rw).Encode(result)
		}
	}

}

func Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	var (
		data  []Person
		input Person
	)
	data = *global

	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	if data == nil || len(data) == 0 {
		result := map[string]interface{}{
			"status":  http.StatusBadRequest,
			"data":    global,
			"message": "Invalid Request",
		}
		json.NewEncoder(rw).Encode(result)
	}

	json.NewDecoder(r.Body).Decode(&input)
	for index, curData := range data {
		if curData.ID == id {
			data[index].FirstName = input.FirstName
			data[index].LastName = input.LastName
			global = &data
			result := map[string]interface{}{
				"status":  http.StatusOK,
				"data":    global,
				"message": "Successfully Updated Data",
			}
			json.NewEncoder(rw).Encode(result)
		}
	}
}
