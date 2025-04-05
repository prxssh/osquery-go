package repo

import (
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
