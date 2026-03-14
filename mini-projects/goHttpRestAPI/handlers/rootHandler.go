package handlers

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// w is interface to response writer and r is for http Request
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Requested URL not found i.e. 404 \n"))
		return
	}
	w.WriteHeader(http.StatusOK)                       //instead of code use constant
	w.Write([]byte("API Version v1 : Hello World \n")) //Simple conversion
}
