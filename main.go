package main

import (
	"net/http"

	"github.com/jsulz/go-algo/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
