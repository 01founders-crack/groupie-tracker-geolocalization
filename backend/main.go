// backend/main.go
package main

import (
	"fmt"
	"groupie-tracker/backend/handlers"

	"net/http"
)

type RelationsData struct {
	ID             int
	DatesLocations map[string][]string
}

func main() {
	// Serve static files and set up routes
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/styles"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/robots.txt")
	})
	http.HandleFunc("/", handlers.HandleRoot)
	http.HandleFunc("/group", handlers.HandleGroup)
	http.HandleFunc("/500", handlers.Handle500)
	// if you need to add new page $create a new handler$ for that in handlers folder
	// If you added new components (more than one html file for one page) go to $common.go$ in the handlers $add else if function$

	port := "443"
	println("Server listening on port https://localhost:" + port)

	// Serve the application over HTTPS with HTTP/2 support
	err := http.ListenAndServeTLS(":"+port, "certificates/server.crt", "certificates/server.key", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
