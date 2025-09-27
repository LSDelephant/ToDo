package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

type Todo struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Done   bool   `json:"done"`
}

var todos = []Todo{
    {ID: 1, Title: "Вивчити Go", Done: false},
    {ID: 2, Title: "Зробити каву", Done: true},
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(todos)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handleTodoByID(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    for _, todo := range todos {
        if todo.ID == id {
            json.NewEncoder(w).Encode(todo)
            return
        }
    }

    http.Error(w, "Todo not found", http.StatusNotFound)
}
