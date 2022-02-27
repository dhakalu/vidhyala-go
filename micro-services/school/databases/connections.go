package databases

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	host         = "localhost"
	port         = 5432
	user         = "postgres"
	password     = "your-password"
	databaseName = "baseapp"
)

func ConnectionString() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v", user, password, host, port, databaseName)
}

func CreateConnection() *pgxpool.Pool {

	fmt.Println("creating pools")

	pgSQLConnectionString := ConnectionString()

	var err error
	var pool *pgxpool.Pool
	pool, err = pgxpool.Connect(context.Background(), pgSQLConnectionString)

	if err != nil {
		panic(err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return pool
}
