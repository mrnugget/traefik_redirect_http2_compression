package main

import (
	"fmt"
	"net/http"
	"os"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	url := "https://redirectwebapp.docker.localhost/end"
	w.Header().Set("Location", url)
	w.WriteHeader(302)

	fmt.Fprintf(w, "<html><body>You are being <a href=\""+url+"\">redirected</a>.</body></html>")
}

func end(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<html><body>
	<p><a href="/redirect">/redirect</a></p>
	</body></html>
	`)
	fmt.Fprintf(w, "Done!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/end", end)
	http.ListenAndServe(":"+port, nil)
}
