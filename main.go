package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"my-school.com/my-school-api/src/controllers"
	server "my-school.com/my-school-api/src/server"
	"my-school.com/my-school-api/src/services"
)

var (
	schoolsService    services.SchoolService        = services.NewSchoolService()
	schoolsController controllers.SchoolsController = controllers.NewSchoolsController(schoolsService)
)

func main() {
	myServer := server.StartApplication()

	myServer.GET("/schools", func(c *gin.Context) {
		c.JSON(200, schoolsController.FindAll())
	})

	myServer.GET("/schools/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		school, err := schoolsController.FindById(int64(id))
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Could not find school with id: " + c.Param("id"),
			})
		} else {
			c.JSON(200, school)
		}
	})

	myServer.DELETE("/schools/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := schoolsController.Delete(int64(id))
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Could not delete school with id: " + c.Param("id"),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "School with id: " + c.Param("id") + " deleted successfully",
			})
		}
	})

	myServer.POST("/schools", func(ctx *gin.Context) {
		record, err := schoolsController.Save(ctx)
		fmt.Printf("error is %v", err)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "Could not save record at this time. Please try again later.",
			})
		} else {
			ctx.JSON(200, record)
		}
	})

	myServer.GET("/", func(conext *gin.Context) {
		conext.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	myServer.Run()
}
