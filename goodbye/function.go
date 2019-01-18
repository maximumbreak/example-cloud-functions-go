package function

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Code    int
	Message string
}

func GoodBye(w http.ResponseWriter, r *http.Request) {
	data := Result{
		200,
		"Goodbye BeforeSecond",
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
