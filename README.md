# vidyalaya-go

This is backend for vidyalaya, an application designed to provide IT services to different educational institutions.

## Tools/Languages used

- PostGres SQL for persistent storage
- Golang
  - gin (HTTP/API layer)
  - pgx (data access layer)

## Folder structure

## Documentation

This project utilizes [gin-swagger](https://github.com/swaggo/gin-swagger) library to autogenerate swagger documentation based on the comments on the routes.

To recreate the docs run command `go run github.com/swaggo/swag/cmd/swag init`

To see the documentation run the application and then visit <http://localhost:8080/api/v1/swagger/index.html#/>

## Running locally

To run this application locally run the command `go run .` from root of the project.
