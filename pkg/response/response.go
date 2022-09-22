package response

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, httpStatus int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")

	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpStatus)
	w.Write(bytes)
}
