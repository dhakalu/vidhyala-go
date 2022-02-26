package routes

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

// Creates/Initializes routes for all of the application
func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}
	v1 := r.router.Group("/api/v1")
	r.addSwagerRoutes(v1)
	r.addSchoolRoutes(v1)
	return r
}

func (r routes) Run(addr ...string) error {
	gin.SetMode(gin.ReleaseMode)
	return r.router.Run()
}
