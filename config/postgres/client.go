package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresClient struct {
	*sql.DB
}

func Init() (*PostgresClient, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s "+"port=%s "+"user=%s "+"password=%s "+"dbname=%s "+"sslmode=disable",
		config.Env.Postgres.Host,
		config.Env.Postgres.Port,
		config.Env.Postgres.User,
		config.Env.Postgres.Password,
		config.Env.Postgres.Dbname,
	)

	client, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return &PostgresClient{client}, nil
}
