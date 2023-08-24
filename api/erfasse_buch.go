package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type ErfasseBuchRequestBody struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type ErfasseBuchResponseBody struct {
	ID uuid.UUID `json:"id"`
}

func ErfasseBuch() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		requestBodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody ErfasseBuchRequestBody
		err = json.Unmarshal(requestBodyBytes, &requestBody)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		id := uuid.New()
		log.Println("received erfasse-buch", id, requestBody)

		// Logik ...

		responseBody := ErfasseBuchResponseBody{
			ID: id,
		}

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
