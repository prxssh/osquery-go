// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package models

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.getOSDetailsStmt, err = db.PrepareContext(ctx, getOSDetails); err != nil {
		return nil, fmt.Errorf("error preparing query GetOSDetails: %w", err)
	}
	if q.getOsqueryInfoStmt, err = db.PrepareContext(ctx, getOsqueryInfo); err != nil {
		return nil, fmt.Errorf("error preparing query GetOsqueryInfo: %w", err)
	}
	if q.listAppsStmt, err = db.PrepareContext(ctx, listApps); err != nil {
		return nil, fmt.Errorf("error preparing query ListApps: %w", err)
	}
	if q.upsertStmt, err = db.PrepareContext(ctx, upsert); err != nil {
		return nil, fmt.Errorf("error preparing query Upsert: %w", err)
	}
	if q.upsertAppStmt, err = db.PrepareContext(ctx, upsertApp); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertApp: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.getOSDetailsStmt != nil {
		if cerr := q.getOSDetailsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOSDetailsStmt: %w", cerr)
		}
	}
	if q.getOsqueryInfoStmt != nil {
		if cerr := q.getOsqueryInfoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOsqueryInfoStmt: %w", cerr)
		}
	}
	if q.listAppsStmt != nil {
		if cerr := q.listAppsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAppsStmt: %w", cerr)
		}
	}
	if q.upsertStmt != nil {
		if cerr := q.upsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertStmt: %w", cerr)
		}
	}
	if q.upsertAppStmt != nil {
		if cerr := q.upsertAppStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertAppStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                 DBTX
	tx                 *sql.Tx
	getOSDetailsStmt   *sql.Stmt
	getOsqueryInfoStmt *sql.Stmt
	listAppsStmt       *sql.Stmt
	upsertStmt         *sql.Stmt
	upsertAppStmt      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                 tx,
		tx:                 tx,
		getOSDetailsStmt:   q.getOSDetailsStmt,
		getOsqueryInfoStmt: q.getOsqueryInfoStmt,
		listAppsStmt:       q.listAppsStmt,
		upsertStmt:         q.upsertStmt,
		upsertAppStmt:      q.upsertAppStmt,
	}
}
