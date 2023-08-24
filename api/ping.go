package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Ping() httprouter.Handle {
	return func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		writer.WriteHeader(http.StatusOK)
	}
}
