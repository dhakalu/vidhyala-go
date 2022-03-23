# vidyalaya-go

This is backend for vidyalaya, an application designed to provide IT services to different educational institutions.

## Tools/Languages used

- PostGres SQL for persistent storage
- Golang
  - gin (HTTP/API layer)
  - pgx (data access layer)
- AWS _(Not implemented)_
  - ECS/EKS to run the docker containers
  - RDS to host postgres sql
  - parameter-store to store environmental parameters

## Folder structure

- **micro-services**: each directory within will be a go module. Each of the module will be a separately deployable service.

- **stack**: tools for maintaining this mono repo.
- **pkg**: common resources/packages shared between micro services.

> **Todo:**
create cli tool to generate new micro-service project

## Documentation

This project utilizes [gin-swagger](https://github.com/swaggo/gin-swagger) library to autogenerate swagger documentation based on the comments on the routes.

To recreate the docs run command `go run github.com/swaggo/swag/cmd/swag init`

To see the documentation run the application and then visit <http://localhost:8080/api/v1/swagger/index.html#/>

## Running locally

To run this application locally run the command `go run .` from root of the micro service you want to run.
