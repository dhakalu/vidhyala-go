package services

import (
	"time"

	"github.com/ulule/deepcopier"
	"vidyalaya.com/micro-services/school/dtos"
	"vidyalaya.com/micro-services/school/entities"
	"vidyalaya.com/micro-services/school/repositories"
)

type SchoolService interface {
	FindAll() ([]entities.School, error)
	FindById(id int64) (entities.School, error)
	Save(school dtos.SchoolCreateRequest) (*entities.School, error)
	Delete(id int64) error
}

type schoolService struct {
	schoolRepository repositories.SchoolRepository
}

func NewSchoolService() SchoolService {
	return &schoolService{
		schoolRepository: repositories.NewSchoolRepository(),
	}
}

func (s *schoolService) FindAll() ([]entities.School, error) {
	return s.schoolRepository.FindAll()
}

func (s *schoolService) FindById(id int64) (entities.School, error) {
	return s.schoolRepository.FindById(id)
}

func (s *schoolService) Save(createRequest dtos.SchoolCreateRequest) (*entities.School, error) {
	schoolEntity := entities.School{}
	// todo is this right way to do this?
	deepcopier.Copy(&createRequest).To(&schoolEntity)
	schoolEntity.CreatedAt = time.Now()
	schoolEntity.UpdatedAt = time.Now()
	return s.schoolRepository.Save(schoolEntity)
}

func (s *schoolService) Delete(id int64) error {
	return s.schoolRepository.Delete(id)
}
