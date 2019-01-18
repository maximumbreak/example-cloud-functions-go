package function

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Code    int
	Message string
}

type Body struct {
	Message string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	body := Body{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.Unmarshal(data, &body); err != nil {
		json.NewEncoder(w).Encode(Result{
			500,
			"Internal Server Error",
		})

	} else {
		json.NewEncoder(w).Encode(Result{
			200,
			body.Message,
		})
	}
}
