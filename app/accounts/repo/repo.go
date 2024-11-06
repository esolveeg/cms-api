package repo

import (
	"context"

	"github.com/esolveeg/cms-api/db"
)

type AccountsRepoInterface interface {
	UserFindNavigationBars(ctx context.Context, userId int32) ([]db.UserFindNavigationBarsRow, error)
	UserFind(ctx context.Context, req db.UserFindParams) (*db.AccountsSchemaUser, error)
	UsersDeleteRestore(ctx context.Context, req []int32) error
	UserDelete(ctx context.Context, userID int32) (*db.AccountsSchemaUser, error)
	UserPermissionsMap(ctx context.Context, userID int32) ([]byte, error)
	UserCreateUpdate(ctx context.Context, req db.UserCreateUpdateParams) (*db.AccountsSchemaUser, error)
	UsersList(ctx context.Context) ([]db.AccountsSchemaUser, error)
	RolesList(ctx context.Context) ([]db.AccountsSchemaRole, error)
	AuthUserIDFindByEmail(ctx context.Context, req string) (*string, error)
	RoleCreateUpdate(ctx context.Context, req db.RoleCreateUpdateParams) (*db.AccountsSchemaRole, error)
	RolesDeleteRestore(ctx context.Context, req []int32) error
}

type AccountsRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewAccountsRepo(store db.Store) AccountsRepoInterface {
	errorHandler := map[string]string{
		"roles_role_name_key":  "roleName",
		"users_user_name_key":  "userName",
		"users_user_email_key": "userEmail",
		"users_user_phone_key": "userPhone",
	}
	return &AccountsRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
