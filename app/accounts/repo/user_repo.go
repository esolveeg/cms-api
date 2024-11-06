package repo

import (
	"context"

	"github.com/esolveeg/cms-api/db"
)

func (repo *AccountsRepo) UserFindNavigationBars(ctx context.Context, userId int32) ([]db.UserFindNavigationBarsRow, error) {
	resp, err := repo.store.UserFindNavigationBars(context.Background(), userId)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return resp, nil
}

func (repo *AccountsRepo) UserDelete(ctx context.Context, userID int32) (*db.AccountsSchemaUser, error) {
	resp, err := repo.store.UserDelete(context.Background(), userID)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}

func (repo *AccountsRepo) UsersList(ctx context.Context) ([]db.AccountsSchemaUser, error) {
	resp, err := repo.store.UsersList(context.Background())
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return resp, nil
}
func (repo *AccountsRepo) UserCreateUpdate(ctx context.Context, req db.UserCreateUpdateParams) (*db.AccountsSchemaUser, error) {
	resp, err := repo.store.UserCreateUpdate(context.Background(), req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}

func (repo *AccountsRepo) AuthUserIDFindByEmail(ctx context.Context, req string) (*string, error) {
	id, err := repo.store.AuthUserIDFindByEmail(context.Background(), req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &id, nil
}
func (repo *AccountsRepo) UsersDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.UsersDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return nil
}
func (repo *AccountsRepo) UserFind(ctx context.Context, req db.UserFindParams) (*db.AccountsSchemaUser, error) {
	resp, err := repo.store.UserFind(context.Background(), req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}
func (repo *AccountsRepo) UserPermissionsMap(ctx context.Context, userID int32) ([]byte, error) {
	resp, err := repo.store.UserPermissionsMap(context.Background(), userID)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return resp, nil
}
