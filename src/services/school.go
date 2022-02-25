package services

import (
	"my-school.com/my-school-api/src/entities"
	"my-school.com/my-school-api/src/repositories"
)

type SchoolService interface {
	FindAll() ([]entities.School, error)
	FindById(id int64) (entities.School, error)
	Save(school entities.School) (entities.School, error)
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

func (s *schoolService) Save(school entities.School) (entities.School, error) {
	return s.schoolRepository.Save(school)
}

func (s *schoolService) Delete(id int64) error {
	return s.schoolRepository.Delete(id)
}
