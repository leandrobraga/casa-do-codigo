package main

import (
	"cdc/internal/data"
	"encoding/json"
	"net/http"
)

func (app *application) createAuthor(w http.ResponseWriter, r *http.Request) {
	var input data.AuthorDtoInput
	json.NewDecoder(r.Body).Decode(&input)

	w.Header().Set("Content-Type", "applciation/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Criando Autor!\n"))
}
