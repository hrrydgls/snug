package handlers

import (
	"net/http"
	"fmt"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the About Page!")
}
