package webrequest

import (
	"fmt"
	"net/http"
)

// READ REQUEST METHOD & PATH
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
	w.Write([]byte("OK"))
}

func Problem1() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
