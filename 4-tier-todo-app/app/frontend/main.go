package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "bytes"
    "strconv"
    "io/ioutil"
)

type Task struct {
    ID       int    `json:"id"`
    TaskName string `json:"taskName"`
}

type TaskCount struct {
    Count int `json:"count"`
}

type Notification struct {
    Message string
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/tasks", tasksHandler)
    http.HandleFunc("/tasks/edit/", editHandler)
    http.HandleFunc("/tasks/delete/", deleteHandler)
    http.HandleFunc("/notification", getNotificationHandler)
    http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))

    
}

func getNotificationHandler(w http.ResponseWriter, r *http.Request) {
    // Fetch the message from the Rust microservice
    resp, err := http.Get("http://notification-svc:8083/notification")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Parse the HTML template
    tmpl, err := template.ParseFiles("templates/notification.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Execute the template with the message
    data := struct {
        Message string
    }{
        Message: string(body),
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}



func indexHandler(w http.ResponseWriter, r *http.Request) {
    tasks, err := getTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := struct {
        Tasks []Task
    }{
        Tasks: tasks,
    }

    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}




func tasksHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        taskName := r.FormValue("taskName")
        if taskName == "" {
            http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
            return
        }

        // Prepare the request body
        requestBody, err := json.Marshal(map[string]string{"taskName": taskName})
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Send a POST request to the Python 10.42.0.50
        // for kubernetes http://backend-svc:5000/tasks
        // for docker compose use the name in compose file
        resp, err := http.Post("http://crud-svc:5000/tasks", "application/json", bytes.NewBuffer(requestBody))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        // Check the response status code
        if resp.StatusCode != http.StatusOK {
            http.Error(w, "Failed to add task", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/", http.StatusFound)
        return
    }
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the task ID from the URL
    id := r.URL.Path[len("/tasks/edit/"):]
    if id == "" {
        http.Error(w, "Task ID is required", http.StatusBadRequest)
        return
    }

    // Parse the task ID as an integer
    taskId, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    if r.Method == http.MethodGet {
        // Retrieve the task from the database by ID
        task, err := getTaskById(taskId)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Render the edit form with the task data
        tmpl, err := template.ParseFiles("templates/edit.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        if err := tmpl.Execute(w, struct{ TaskID int; TaskName string }{TaskID: task.ID, TaskName: task.TaskName}); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    } else if r.Method == http.MethodPost {
        // Update the task in the database
        taskName := r.FormValue("taskName")

        // Prepare the request body
        requestBody, err := json.Marshal(map[string]string{"taskName": taskName})
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Send a PUT request to update the task in the Python 10.42.0.50
        req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://crud-svc:5000/tasks/edit/%d", taskId), bytes.NewBuffer(requestBody))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        req.Header.Set("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            http.Error(w, "Failed to update task", http.StatusInternalServerError)
            return
        }

        // Redirect to the home page on successful update
        http.Redirect(w, r, "/", http.StatusFound)
    }
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the task ID from the URL
    id := r.URL.Path[len("/tasks/delete/"):]
    if id == "" {
        http.Error(w, "Task ID is required", http.StatusBadRequest)
        return
    }

    // Parse the task ID as an integer
    taskId, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    // Send a DELETE request to delete the task in the Python 10.42.0.50
    req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://crud-svc:5000/tasks/delete/%d", taskId), nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        http.Error(w, "Failed to delete task", http.StatusInternalServerError)
        return
    }

    // Redirect to the home page on successful delete
    http.Redirect(w, r, "/", http.StatusFound)
}

func getTasks() ([]Task, error) {
    resp, err := http.Get("http://crud-svc:5000/tasks")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var tasks []Task
    if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
        return nil, err
    }

    return tasks, nil
}

func getTaskById(taskId int) (Task, error) {
    resp, err := http.Get(fmt.Sprintf("http://crud-svc:5000/tasks/%d", taskId))
    if err != nil {
        return Task{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return Task{}, fmt.Errorf("Failed to get task: %s", resp.Status)
    }

    var task Task
    if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
        return Task{}, err
    }

    return task, nil
}
