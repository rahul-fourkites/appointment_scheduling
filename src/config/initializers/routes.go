package initializer

import (
	"../../../pkg/mux"
	"../../../src/app/controllers/api/v1"
	"fmt"
	"net/http"
)

func InitializeRoutes() {
	fmt.Print("Initializing routes...\n")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/test/home", controller.Home)
	http.Handle("/", r)
}
