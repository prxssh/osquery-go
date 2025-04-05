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
				Name: "Version",
				SQL: `
                SELECT version AS version_value, 'OS' AS version_type
                FROM os_version
                UNION ALL
                SELECT version AS version_value, 'OSQuery' AS version_type
                FROM osquery_info;`,
				Every:   10 * time.Second,
				Handler: handleOsqueryGetVersionJob,
			},
			{
				Name:    "Installed Applications",
				SQL:     "SELECT * FROM apps;",
				Every:   10 * time.Second,
				Handler: handleOsqueryGetAppJob,
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

func handleOsqueryGetVersionJob(
	repo *repo.Repo,
	data []map[string]any,
) error {
	if err := repo.Versions.Upsert(context.Background(), data); err != nil {
		log.Error().Err(err).Msg("failed to upsert versions")
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
