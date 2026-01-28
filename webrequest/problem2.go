package webrequest

import "net/http"

//HANDLE GET & POST DIFFERENTLY
func handler1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("This is method GET"))
	case http.MethodPut:
		w.Write([]byte("This is method PUT"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Problem2() {
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}
