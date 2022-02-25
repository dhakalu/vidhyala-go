package main

import (
	routes "my-school.com/my-school-api/src/routes"
)

func main() {
	allRoutes := routes.NewRoutes()
	allRoutes.Run()
}
