package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)


type todo struct {
	Id   string
	Title string
}

var todos = make([]todo, 5)

func main() {
    todos = append(todos, todo{Id: "1", Title: "Master go"})
    todos = append(todos, todo{Id: "2", Title: "Master htmx"})

	mux := http.NewServeMux()
    mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/todos", handleTodos)

    http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tn := "home.html"
	t, err := template.New(tn).ParseFiles(filepath.Join("pages", tn))

	if err != nil {
		panic(err)
	}

	err = t.Execute(w, nil)
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
    method := r.Method
    if method == http.MethodPost {
        createTodo(w, r)
        return
    } else if method == http.MethodDelete {
        deleteTodo(w, r)
        return
    }
    geTodos(w, r)
}

func geTodos(w http.ResponseWriter, r *http.Request) {
	tn := "todos.html"
	t, err := template.New(tn).ParseFiles(filepath.Join("components", tn))

	if err != nil {
		panic(err)
	}

	err = t.Execute(w, todos)

	if err != nil {
		panic(err)
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	tn := "todos.html"
	t, err := template.New(tn).ParseFiles(filepath.Join("components", tn))

    if err != nil {
        panic(err)
    }

    todoId  := strconv.Itoa(int(time.Now().UnixMilli()))
    newTodo := todo{
        Id: todoId,
        Title: r.FormValue("todo"),
    }
    todos = append(todos, newTodo)

	err = t.Execute(w, todos)

	if err != nil {
		panic(err)
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	tn := "todos.html"
	t, err := template.New(tn).ParseFiles(filepath.Join("components", tn))

    if err != nil {
        panic(err)
    }

    id := r.FormValue("id")
    filtered := make([]todo, 0)

    for _, v := range todos {
        if v.Id == id {
            continue
        }
        filtered = append(filtered, v)
    }

    todos = filtered
	err = t.Execute(w, todos)

	if err != nil {
		panic(err)
	}
}
