package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TEXT(w http.ResponseWriter, statusCode int, data string) {
	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte(data))

}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
