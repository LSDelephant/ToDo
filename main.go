package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/todos", handleTodos)
    http.HandleFunc("/todos/", handleTodoByID)

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
