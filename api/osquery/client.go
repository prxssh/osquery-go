package osquery

import "github.com/prxssh/osquery-go/internal/repo"

type OsqueryAPIService struct {
	repo *repo.Repo
}

func NewOsqueryAPIService(repo *repo.Repo) *OsqueryAPIService {
	return &OsqueryAPIService{
		repo: repo,
	}
}
