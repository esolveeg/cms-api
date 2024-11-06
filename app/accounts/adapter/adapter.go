package adapter

import (
	"github.com/esolveeg/cms-api/db"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	"github.com/supabase-community/auth-go/types"
)

type AccountsAdapterInterface interface {
	UserLoginSqlFromGrpc(req *devkitv1.UserLoginRequest) (*db.UserFindParams, *types.TokenRequest)
	UserResetPasswordSupaFromGrpc(req *devkitv1.UserResetPasswordRequest) *types.VerifyForUserRequest
	UserLoginGrpcFromSql(resp *db.AccountsSchemaUser) *devkitv1.UserLoginResponse
	UserFindNavigationBarsGrpcFromSql(resp *[]db.UserFindNavigationBarsRow) (*[]*devkitv1.NavigationBarItem, error)
	UserCreateUpdateGrpcFromSql(resp *db.AccountsSchemaUser) *devkitv1.UserCreateUpdateResponse
	UsersListGrpcFromSql(resp []db.AccountsSchemaUser) *devkitv1.UsersListResponse
	UserCreateUpdateSqlFromGrpc(req *devkitv1.UserCreateUpdateRequest) *db.UserCreateUpdateParams
	UserEntityGrpcFromSql(resp *db.AccountsSchemaUser) *devkitv1.AccountsSchemaUser
	RolesListGrpcFromSql(resp []db.AccountsSchemaRole) *devkitv1.RolesListResponse
	RoleEntityGrpcFromSql(resp *db.AccountsSchemaRole) *devkitv1.AccountsSchemaRole
	RoleCreateUpdateSqlFromGrpc(req *devkitv1.RoleCreateUpdateRequest) *db.RoleCreateUpdateParams
	RoleCreateUpdateGrpcFromSql(resp *db.AccountsSchemaRole) *devkitv1.RoleCreateUpdateResponse
}

type AccountsAdapter struct {
}

func NewAccountsAdapter() AccountsAdapterInterface {
	return &AccountsAdapter{}
}
