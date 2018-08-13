package initializer

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	fmt.Print("Starting server...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
