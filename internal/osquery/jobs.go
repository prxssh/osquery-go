package osquery

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/prxssh/osquery-go/config/postgres"
	"github.com/prxssh/osquery-go/internal/repo"
	"github.com/rs/zerolog/log"
)

type JobsInfo struct {
	repo    *repo.Repo
	queries []*Query
}

type Query struct {
	Name    string
	SQL     string
	Every   time.Duration
	Handler func(repo *repo.Repo, results []map[string]any) error
}

func NewOsqueryJobs(dbClient *postgres.PostgresClient) *JobsInfo {
	return &JobsInfo{
		repo: repo.NewRepo(dbClient),
		queries: []*Query{
			{
				Name:    "OS Version",
				SQL:     "SELECT * FROM os_version;",
				Every:   10 * time.Second,
				Handler: handleOsqueryGetOsVersionJob,
			},
			{
				Name:    "Installed Applications",
				SQL:     "SELECT * FROM apps;",
				Every:   10 * time.Second,
				Handler: handleOsqueryGetAppJob,
			},
			{
				Name:    "Osquery Info",
				SQL:     "SELECT * FROM osquery_info;",
				Every:   10 * time.Second,
				Handler: handleOsqueryGetInfoJob,
			},
		},
	}
}

func (j *JobsInfo) ScheduleOsqueryJobs() chan struct{} {
	done := make(chan struct{})

	for _, query := range j.queries {
		ticker := time.NewTicker(query.Every)

		go func() {
			for {
				select {
				case <-ticker.C:
					executeAndProcessQuery(j.repo, query)
				case <-done:
					ticker.Stop()
					return
				}
			}
		}()
	}

	return done
}

func executeAndProcessQuery(repo *repo.Repo, query *Query) error {
	log.Info().
		Str("name", query.Name).
		Str("query", query.SQL).
		Msg("executing osquery")

	cmd := exec.Command("osqueryi", "--json", query.SQL)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("execute osquery failed: %w", err)
	}

	var results []map[string]any
	if err := json.Unmarshal(output, &results); err != nil {
		log.Error().Err(err).Msg("failed to parse osquery response")
		return fmt.Errorf("failed to parse osqueryi output: %w", err)
	}

	return query.Handler(repo, results)
}

func handleOsqueryGetOsVersionJob(
	repo *repo.Repo,
	data []map[string]any,
) error {
	_, err := repo.OsVersion.Upsert(context.Background(), data[0])
	if err != nil {
		log.Error().Err(err).Msg("failed to upsert osversion")
		return err
	}

	return nil
}

func handleOsqueryGetAppJob(repo *repo.Repo, data []map[string]any) error {
	if err := repo.Apps.UpsertWithTx(context.Background(), data); err != nil {
		log.Error().Err(err).Msg("failed to upsert apps details")
		return err
	}

	return nil
}

func handleOsqueryGetInfoJob(repo *repo.Repo, data []map[string]any) error {
	_, err := repo.OsqueryInfo.Upsert(context.Background(), data[0])
	if err != nil {
		log.Error().Err(err).Msg("failed to upsert osquery info")
		return err
	}

	return nil
}
