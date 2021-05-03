package projects

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"time"

	"github.com/devpies/devpie-client-core/projects/platform/database"
)

var (
	ErrNotFound  = errors.New("project not found")
	ErrInvalidID = errors.New("id provided was not a valid UUID")
)

func Retrieve(ctx context.Context, repo *database.Repository, pid string) (Project, error) {
	var p Project

	if _, err := uuid.Parse(pid); err != nil {
		return p, ErrInvalidID
	}

	// TODO: Check ownership at Team and User level before granting access
	// TODO: Also allow site Admins (where role==admin in Auth0)

	stmt := repo.SQ.Select(
		"project_id",
		"name",
		"prefix",
		"description",
		"team_id",
		"user_id",
		"active",
		"public",
		"column_order",
		"updated_at",
		"created_at",
	).From(
		"projects",
	).Where(sq.Eq{"project_id": "?"})

	q, args, err := stmt.ToSql()
	if err != nil {
		return p, errors.Wrapf(err, "building query: %v", args)
	}

	row := repo.DB.QueryRowContext(ctx, q, pid)
	err = row.Scan(&p.ID, &p.Name, &p.Prefix, &p.Description, &p.TeamID, &p.UserID, &p.Active, &p.Public, (*pq.StringArray)(&p.ColumnOrder), &p.UpdatedAt, &p.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, ErrNotFound
		}
		return p, err
	}

	return p, nil
}

func List(ctx context.Context, repo *database.Repository, uid string) ([]Project, error) {
	var p Project
	var ps = make([]Project, 0)

	stmt := repo.SQ.Select(
		"project_id",
		"name",
		"prefix",
		"description",
		"team_id",
		"user_id",
		"active",
		"public",
		"column_order",
		"updated_at",
		"created_at",
	).From("projects").Where(sq.Eq{"user_id": "?"})
	q, args, err := stmt.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "building query: %v", args)
	}

	rows, err := repo.DB.QueryContext(ctx, q, uid)
	if err != nil {
		return nil, errors.Wrap(err, "selecting projects")
	}
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Name, &p.Prefix, &p.Description, &p.TeamID, &p.UserID, &p.Active, &p.Public, (*pq.StringArray)(&p.ColumnOrder), &p.UpdatedAt, &p.CreatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "scanning row into Struct")
		}
		ps = append(ps, p)
	}

	return ps, nil
}

func Create(ctx context.Context, repo *database.Repository, np NewProject, uid string, now time.Time) (Project, error) {
	p := Project{
		ID:          uuid.New().String(),
		Name:        np.Name,
		Prefix:      fmt.Sprintf("%s-", np.Name[:3]),
		Active:      true,
		UserID:      uid,
		TeamID:      np.TeamID,
		ColumnOrder: []string{"column-1", "column-2", "column-3", "column-4"},
		UpdatedAt: now.UTC(),
		CreatedAt: now.UTC(),
	}

	stmt := repo.SQ.Insert(
		"projects",
	).SetMap(map[string]interface{}{
		"project_id":   p.ID,
		"name":         p.Name,
		"prefix":       p.Prefix,
		"team_id":      p.TeamID,
		"description":  "",
		"user_id":      p.UserID,
		"column_order": pq.Array(p.ColumnOrder),
		"updated_at":   p.UpdatedAt,
		"created_at":   p.CreatedAt,
	})

	if _, err := stmt.ExecContext(ctx); err != nil {
		return p, errors.Wrapf(err, "inserting project: %v", p)
	}

	return p, nil
}

func Update(ctx context.Context, repo *database.Repository, pid string, update UpdateProject) (Project, error) {
	p, err := Retrieve(ctx, repo, pid)
	if err != nil {
		return p, err
	}

	if update.Name != nil {
		p.Name = *update.Name
	}
	if update.Description != nil {
		p.Description = *update.Description
	}
	if update.Active != nil {
		p.Active = *update.Active
	}
	if update.Public != nil {
		p.Public = *update.Public
	}
	if update.ColumnOrder != nil {
		p.ColumnOrder = update.ColumnOrder
	}
	if update.TeamID != nil {
		p.TeamID = *update.TeamID
	}

	stmt := repo.SQ.Update(
		"projects",
	).SetMap(map[string]interface{}{
		"name":         p.Name,
		"description":  p.Description,
		"active":       p.Active,
		"public":       p.Public,
		"column_order": pq.Array(p.ColumnOrder),
		"team_id":      p.TeamID,
		"updated_at":   update.UpdatedAt,
	}).Where(sq.Eq{"project_id": pid})

	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return p, errors.Wrap(err, "updating project")
	}

	return p, nil
}

func Delete(ctx context.Context, repo *database.Repository, pid, uid string) error {
	if _, err := uuid.Parse(pid); err != nil {
		return ErrInvalidID
	}

	stmt := repo.SQ.Delete(
		"projects",
	).Where(sq.Eq{"project_id": pid, "user_id": uid})

	if _, err := stmt.ExecContext(ctx); err != nil {
		return errors.Wrapf(err, "deleting project %s", pid)
	}

	return nil
}