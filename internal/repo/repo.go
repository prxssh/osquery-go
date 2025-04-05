package repo

import (
	"github.com/prxssh/osquery-go/config/postgres"
)

type Repo struct {
	Apps     IApps
	Versions IVersions
	client   *postgres.PostgresClient
}

func NewRepo(dbClient *postgres.PostgresClient) *Repo {
	return &Repo{
		client:   dbClient,
		Apps:     NewAppsRepo(dbClient),
		Versions: NewVersionsRepo(dbClient),
	}
}
