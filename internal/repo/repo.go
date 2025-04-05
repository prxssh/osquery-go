package repo

import (
	"context"
	"database/sql"

	"github.com/prxssh/osquery-go/config/postgres"
)

type Repo struct {
	Apps        IApps
	OsVersion   IOsVersion
	OsqueryInfo IOsqueryInfo
	client      *postgres.PostgresClient
}

func NewRepo(dbClient *postgres.PostgresClient) *Repo {
	return &Repo{
		client:      dbClient,
		Apps:        NewAppsRepo(dbClient),
		OsVersion:   NewOsVersionRepo(dbClient),
		OsqueryInfo: NewOsqueryInfoRepo(dbClient),
	}
}

func (r *Repo) BeginTx() (*sql.Tx, error) {
	return r.client.Begin()
}

func (r *Repo) ExecuteTransaction(
	ctx context.Context,
	operation func(txn *sql.Tx) (any, error),
) (any, error) {
	txn, err := r.client.Begin()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	res, err := operation(txn)
	if err != nil {
		return nil, err
	}

	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return res, nil
}
