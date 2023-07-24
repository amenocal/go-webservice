package main

import (
	"fmt"
	"net/http"

	"github.com/amenocal/go-webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	fmt.Println("Starting Server...")
	http.ListenAndServe(":3000", nil) //second parameter is the ServeMux or Serve multiplexer, object that will handle all request s that are coming in.

}
