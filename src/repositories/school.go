package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"my-school.com/my-school-api/src/databases"
	"my-school.com/my-school-api/src/entities"
)

type SchoolRepository interface {
	FindAll() ([]entities.School, error)
	FindById(id int64) (entities.School, error)
	Save(school entities.School) (*entities.School, error)
	Delete(id int64) error
}

type schoolRepository struct {
	db pgxpool.Pool
}

func NewSchoolRepository() SchoolRepository {
	return &schoolRepository{
		db: *databases.CreateConnection(),
	}
}

func (r *schoolRepository) FindAll() ([]entities.School, error) {
	// todo implement
	var schools []entities.School

	rows, err := r.db.Query(context.Background(), `SELECT * FROM schools`)
	if err != nil {
		return schools, err
	} else {
		for rows.Next() {
			var school entities.School
			err := rows.Scan(
				&school.Id,
				&school.Name,
				&school.Address,
				&school.City,
				&school.Providence,
				&school.Zip,
				&school.Phone,
				&school.Fax,
				&school.Email,
				&school.Website,
				&school.CreatedAt,
				&school.UpdatedAt,
			)
			if err != nil {
				fmt.Println(err)
				// todo should you end and return an error?
				continue
			}
			schools = append(schools, school)
		}
	}
	return schools, nil
}

func (r *schoolRepository) FindById(id int64) (entities.School, error) {
	var school entities.School

	row := r.db.QueryRow(context.Background(), `SELECT * FROM schools WHERE id = $1`, id)

	err := row.Scan(
		&school.Id,
		&school.Name,
		&school.Address,
		&school.City,
		&school.Providence,
		&school.Zip,
		&school.Phone,
		&school.Fax,
		&school.Email,
		&school.Website,
		&school.CreatedAt,
		&school.UpdatedAt,
	)
	if err != nil {
		return entities.School{}, err
	}
	return school, nil
}

func (r *schoolRepository) Save(school entities.School) (*entities.School, error) {

	_, err := r.FindById(school.Id)
	if err != nil {
		sql := `INSERT INTO schools(
		school_name,
		street_address,
		city,
		providence,
		zip,
		phone,
		fax,
		email,
		website,
		created_at,
		updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`
		_, err := r.db.Exec(context.Background(), sql,
			school.Name,
			school.Address,
			school.City,
			school.Providence,
			school.Zip,
			school.Phone,
			school.Fax,
			school.Email,
			school.Website,
			school.CreatedAt,
			school.UpdatedAt,
		)
		if err != nil {
			fmt.Printf("%v \n", err)
			return &entities.School{}, err
		}
		return &school, nil
	} else {
		return nil, err
	}
}

func (r *schoolRepository) Delete(id int64) error {
	sql := `DELETE FROM schools WHERE id = $1`
	_, err := r.db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}
	return nil
}
