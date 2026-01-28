package miniapi

import (
	"encoding/json"
	"net/http"
	"sync"
)

// defines the JSON structure
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Slice storage for Users(in-memory)
var users []User
var mu sync.Mutex

// Handler for /users
func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//http.MethodPost adds a new user
	case http.MethodPost:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		mu.Lock()
		users = append(users, user)
		defer mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "user added successfully.",
		})
		//http.MethodGet fetches all the users
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		mu.Lock()
		json.NewEncoder(w).Encode(users)
		defer mu.Unlock()
		//http.MethodPut updates a specific user
	case http.MethodPut:
		var updated User
		err := json.NewDecoder(r.Body).Decode(&updated)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		found := false
		for i, u := range users {
			if u.Name == updated.Name {
				users[i] = updated
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "user updated successful",
		})
		//http.MethodDelete deletes a specific user
	case http.MethodDelete:
		var req struct {
			Name string `json:"name"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		found := false
		for i, u := range users {
			if u.Name == req.Name {
				users = append(users[:i], users[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User deleted successfully.",
		})
		//default to handle invalid cases
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// UserAPI() acts as the main func()
func UsersAPI() {
	http.HandleFunc("/users", userHandler)
	http.ListenAndServe(":8080", nil)
}
