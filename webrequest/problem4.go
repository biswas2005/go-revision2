package webrequest

import (
	"net/http"
)

// READ HEADERS
func handler4(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Authorized"))
}

func Problem4() {
	http.HandleFunc("/", handler4)
	http.ListenAndServe(":8080", nil)
}
