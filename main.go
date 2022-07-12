package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID          string `json:"id"`
	TasksName   string `json:"task_name"`
	TaskDetails string `json:"task_details"`
	Date        string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task1 := Tasks{
		ID:          "1",
		TasksName:   "New Project",
		TaskDetails: "You must have to lead the project and complete it within the given time",
		Date:        "08/07/2022",
	}
	tasks = append(tasks, task1)

	task2 := Tasks{
		ID:          "2",
		TasksName:   "HCI Project",
		TaskDetails: "We have to finish the HCI project as soon as possible.",
		Date:        "08/07/2022",
	}
	tasks = append(tasks, task2)

	fmt.Println("Your tasks are: ", tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I m a home page.")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if taskId["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I m a home page.")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I m a home page.")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I m a home page.")
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")

	// port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	allTasks()
	handleRoutes()

}
