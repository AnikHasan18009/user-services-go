package web

import (
	"fmt"
	"net/http"
	"user-service/web/middlewares"
)

func StartServer() *http.ServeMux {
	mux := http.NewServeMux()

	manager := middlewares.NewManager()
	InitRoutes(mux, manager)

	fmt.Println("Server Started")

	return mux
}
