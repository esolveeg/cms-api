// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	RoleCreateUpdate(ctx context.Context, arg RoleCreateUpdateParams) (AccountsSchemaRole, error)
	RolesDeleteRestore(ctx context.Context, records []int32) error
	RolesList(ctx context.Context) ([]AccountsSchemaRole, error)
	UserCreateUpdate(ctx context.Context, arg UserCreateUpdateParams) (AccountsSchemaUser, error)
	UserFind(ctx context.Context, arg UserFindParams) (AccountsSchemaUser, error)
	UserPermissionsMap(ctx context.Context, userID int32) ([]byte, error)
	UsersDeleteRestore(ctx context.Context, records []int32) error
	UsersList(ctx context.Context) ([]AccountsSchemaUser, error)
}

var _ Querier = (*Queries)(nil)
