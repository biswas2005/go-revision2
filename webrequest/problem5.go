package webrequest

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// READ JSON REQUEST BODY
func handler5(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Error in method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User Created",
		"name":    user.Name,
	})
}

func Problem5() {
	http.HandleFunc("/", handler5)
	http.ListenAndServe(":8080", nil)
}
