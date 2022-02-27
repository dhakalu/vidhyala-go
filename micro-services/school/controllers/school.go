package controllers

import (
	"github.com/gin-gonic/gin"
	"vidhyalayer.com/micro-services/school/dtos"
	"vidhyalayer.com/micro-services/school/entities"
	"vidhyalayer.com/micro-services/school/services"
)

type SchoolsController interface {
	FindAll() ([]entities.School, error)
	FindById(id int64) (entities.School, error)
	Save(ctx *gin.Context) (*entities.School, error)
	Delete(id int64) error
}

type controller struct {
	service services.SchoolService
}

func NewSchoolsController(service services.SchoolService) SchoolsController {
	return controller{
		service: service,
	}
}

func (c controller) FindAll() ([]entities.School, error) {
	return c.service.FindAll()
}

func (c controller) FindById(id int64) (entities.School, error) {
	return c.service.FindById(id)
}

func (c controller) Save(ctx *gin.Context) (*entities.School, error) {
	var school dtos.SchoolCreateRequest
	err := ctx.BindJSON(&school)
	if err != nil {
		// todo return a custom error so it can be handled in the route properly
		return nil, err
	}
	return c.service.Save(school)
}

func (c controller) Delete(id int64) error {
	return c.service.Delete(id)
}
