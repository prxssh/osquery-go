package repo

import (
	"github.com/prxssh/osquery-go/config/postgres"
)

type Repo struct {
	apps        IApps
	osVersion   IOsVersion
	osqueryInfo IOsqueryInfo
	client      *postgres.PostgresClient
}

func NewRepo(dbClient *postgres.PostgresClient) *Repo {
	return &Repo{
		client:      dbClient,
		apps:        NewAppsRepo(dbClient),
		osVersion:   NewOsVersionRepo(dbClient),
		osqueryInfo: NewOsqueryInfoRepo(dbClient),
	}
}
