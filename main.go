package main

import (
	"fmt"
	"net/http"
	"os"
)

const webDir = "web"

func main() {
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}
	fmt.Println("Server is running on http://localhost:" + port + " (ノಠ益ಠ)ノ彡┻━┻")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
