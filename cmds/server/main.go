package main

import (
	"fmt"
	"log"
	"net/http"
)

// All the server is responsible for here is serving our static assets, which is only needed during development.
// Nothing fancy and could be done with any other tech - doesn't have to be Go!
func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))

	fmt.Println("🚀 Server launching! Navigate to http://localhost:8080/ to see the magic ✨")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
