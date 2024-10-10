// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts_role.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const roleCreate = `-- name: RoleCreate :one
INSERT INTO accounts_schema.roles(role_name, role_description)
    VALUES ($1, $2)
RETURNING
    role_id, role_name, role_description, created_at, updated_at, deleted_at
`

type RoleCreateParams struct {
	RoleName        string      `json:"role_name"`
	RoleDescription pgtype.Text `json:"role_description"`
}

func (q *Queries) RoleCreate(ctx context.Context, arg RoleCreateParams) (AccountsSchemaRole, error) {
	row := q.db.QueryRow(ctx, roleCreate, arg.RoleName, arg.RoleDescription)
	var i AccountsSchemaRole
	err := row.Scan(
		&i.RoleID,
		&i.RoleName,
		&i.RoleDescription,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
