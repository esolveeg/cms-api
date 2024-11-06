package repo

import (
	"context"

	"github.com/esolveeg/cms-api/db"
)

func (repo *AccountsRepo) RolesList(ctx context.Context) ([]db.AccountsSchemaRole, error) {
	resp, err := repo.store.RolesList(context.Background())
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return resp, nil
}
func (repo *AccountsRepo) RoleCreateUpdate(ctx context.Context, req db.RoleCreateUpdateParams) (*db.AccountsSchemaRole, error) {
	resp, err := repo.store.RoleCreateUpdate(context.Background(), req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}
func (repo *AccountsRepo) RolesDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.RolesDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)

	}
	return nil
}
