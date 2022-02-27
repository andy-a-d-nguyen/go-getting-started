package main

import (
	"net/http"

	"go-getting-started.com/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}