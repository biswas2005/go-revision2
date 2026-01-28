package webrequest

import (
	"fmt"
	"net/http"
)

// READ QUERY PARAMETERS
func handler3(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(w, "Hello %s , age %s", name, age)
}
func Problem3() {
	http.HandleFunc("/", handler3)
	http.ListenAndServe(":8080", nil)
}
