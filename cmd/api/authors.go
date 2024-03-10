package main

import (
	"fmt"
	"net/http"
)

func (app *application) createAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Craindo Autor")
}
