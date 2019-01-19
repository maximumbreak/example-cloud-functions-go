package function

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Result struct {
	Code    int
	Message string
	Token   string
}

type Body struct {
	Message string
}

const (
	mySigningKey = "BeforeSecond-Golang"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	body := Body{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.Unmarshal(data, &body); err != nil {
		json.NewEncoder(w).Encode(Result{
			500,
			"Internal Server Error",
			"",
		})
		return

	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"Message": body.Message,
			"exp":     time.Now().Add(time.Hour * 72).Unix(),
		})
		tokenString, err := token.SignedString([]byte(mySigningKey))

		if err != nil {
			json.NewEncoder(w).Encode(Result{
				500,
				err.Error(),
				"",
			})
			return
		}

		json.NewEncoder(w).Encode(Result{
			200,
			body.Message,
			tokenString,
		})
		return
	}
}
