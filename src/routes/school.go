package routes

import (
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

// @Summary Get all schools.
// @Description Retrieve list of all the available schools.
// @Tags school
// @Produce json
// @Success 200 {object} []entities.School
// @Failure 400 {object} routes.ErrorResponse
// @Router /api/v1/schools [get]
func fetchAllSchool(c *gin.Context) {
	schools, err := schoolsController.FindAll()
	if err != nil {
		// log.Error("Could not find schools")
		c.JSON(400, NewErrorResponse("Some error occurred while fetching schools. Try again later."))
	} else {
		c.JSON(200, schools)
	}
}

// @Summary Get single school
// @Description Retreive a school by its id. Id is unique for each school.
// @Tags school
// @Produce json
// @Success 200 {object} entities.School
// @Failure 400 {object} routes.ErrorResponse
// @Router /api/v1/schools/{id} [get]
// @Param id path int true "School ID"
func fetchSchoolById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	school, err := schoolsController.FindById(int64(id))
	if err != nil {
		log.Error(err.Error())
		c.JSON(400, NewErrorResponse("Could not find school with id: "+c.Param("id")))
	} else {
		c.JSON(200, school)
	}
}

// @Summary Delete a school
// @Description Removes a school by its id. Id is unique for each school.
// @Tags school
// @Produce json
// @Success 200 {object} entities.School
// @Failure 400 {object} routes.ErrorResponse
// @Router /api/v1/schools/{id} [delete]
// @Param id path int true "School ID"
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

// @Summary Create a school
// @Description Adds a school in database.
// @Tags school
// @Accept json
// @Produce json
// @Success 200 {object} entities.School
// @Failure 400 {object} routes.ErrorResponse
// @Router /api/v1/schools [post]
// @Param school body dtos.SchoolCreateRequest true "JSON object representing a school."
func createSchool(ctx *gin.Context) {
	record, err := schoolsController.Save(ctx)
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(400, NewErrorResponse(err.Error()))
	} else {
		ctx.JSON(200, record)
	}
}
