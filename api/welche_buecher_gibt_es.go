package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type Book struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
}

type WelcheBuecherGibtEsResponseBody struct {
	Books []Book `json:"books"`
}

func WelcheBuecherGibtEs() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		// Logik ...

		responseBody := WelcheBuecherGibtEsResponseBody{
			Books: []Book{
				{
					ID:     uuid.New(),
					Title:  "The Hitchhiker's Guide to the Galaxy",
					Author: "Douglas Adams",
				},
			},
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
