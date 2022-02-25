package controllers

import (
	"github.com/gin-gonic/gin"
	"my-school.com/my-school-api/src/entities"
	"my-school.com/my-school-api/src/services"
)

type SchoolsController interface {
	FindAll() []entities.School
	FindById(id int64) (entities.School, error)
	Save(ctx *gin.Context) (entities.School, error)
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

func (c controller) FindAll() []entities.School {
	return c.service.FindAll()
}

func (c controller) FindById(id int64) (entities.School, error) {
	return c.service.FindById(id)
}

func (c controller) Save(ctx *gin.Context) (entities.School, error) {
	var school entities.School
	ctx.BindJSON(&school)
	return c.service.Save(school)
}

func (c controller) Delete(id int64) error {
	return c.service.Delete(id)
}
