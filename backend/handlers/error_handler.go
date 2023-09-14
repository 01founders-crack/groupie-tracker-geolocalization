package handlers

import "net/http"

func Handle500(w http.ResponseWriter, r *http.Request) {
	data := struct{}{}
	renderTemplate(w, "500", data)
}
