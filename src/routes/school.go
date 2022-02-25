package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my-school.com/my-school-api/src/controllers"
	"my-school.com/my-school-api/src/services"
)

var (
	schoolsService    services.SchoolService        = services.NewSchoolService()
	schoolsController controllers.SchoolsController = controllers.NewSchoolsController(schoolsService)
)

func (r routes) addSchoolRoutes(rg *gin.RouterGroup) {
	schoolRoutes := rg.Group("/schools")

	schoolRoutes.GET("/", fetchAllSchool)
	schoolRoutes.GET("/:id", fetchSchoolById)
	schoolRoutes.DELETE("/:id", deleteSchoolById)
	schoolRoutes.POST("/", createSchool)
}

func fetchAllSchool(c *gin.Context) {
	schools, err := schoolsController.FindAll()
	if err != nil {
		log.Error(err.Error())
		c.JSON(400, gin.H{
			"message": "Some error occurred while fetching schools. Try again later.",
		})
	} else {
		c.JSON(200, schools)
	}
}

func fetchSchoolById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	school, err := schoolsController.FindById(int64(id))
	if err != nil {
		log.Error(err.Error())
		c.JSON(400, gin.H{
			"message": "Could not find school with id: " + c.Param("id"),
		})
	} else {
		c.JSON(200, school)
	}
}

func deleteSchoolById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := schoolsController.Delete(int64(id))
	if err != nil {
		log.Error(err.Error())
		c.JSON(400, gin.H{
			"message": "Could not delete school with id: " + c.Param("id"),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "School with id: " + c.Param("id") + " deleted successfully",
		})
	}
}

func createSchool(ctx *gin.Context) {
	record, err := schoolsController.Save(ctx)
	fmt.Printf("error is %v", err)
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(400, gin.H{
			"message": "Could not save record at this time. Please try again later.",
		})
	} else {
		ctx.JSON(200, record)
	}
}

// myServer.GET("/", func(conext *gin.Context) {
// 	conext.JSON(200, gin.H{
// 		"message": "Oops! None of the routes matched your request. Please check the documentation.",
// 	})
// })

// myServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
