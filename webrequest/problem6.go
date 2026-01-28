package webrequest

import (
	"encoding/json"
	"net/http"
)

// LIMIT REQUEST BODY SIZE
func handler6(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Body too large or invalid", http.StatusBadRequest)
		return
	}
	w.Write([]byte("OK"))
}

func Problem6() {
	http.HandleFunc("/", handler6)
	http.ListenAndServe(":8080", nil)
}
