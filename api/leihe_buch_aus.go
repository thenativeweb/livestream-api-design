package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type LeiheBuchAusRequestBody struct {
	ID uuid.UUID `json:"id"`
}

type LeiheBuchAusResponseBody struct {
}

func LeiheBuchAus() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		requestBodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody LeiheBuchAusRequestBody
		err = json.Unmarshal(requestBodyBytes, &requestBody)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("received leihe-buch-aus", requestBody.ID, requestBody)

		// Logik ...

		responseBody := LeiheBuchAusResponseBody{}

		responseBodyBytes, err := json.Marshal(responseBody)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(responseBodyBytes)
	}
}
