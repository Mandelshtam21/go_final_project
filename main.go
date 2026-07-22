package main

import (
	"fmt"
	"net/http"
	"os"

	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
)

const webDir = "web"

func main() {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "./scheduler.db"
	}
	err := db.Init(dbFile)
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	api.Init()
	fmt.Println("Server is running on http://localhost:" + port + " (ノಠ益ಠ)ノ彡┻━┻")
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
