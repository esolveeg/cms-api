package usecase

import (
	"context"
	apiv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (u *AccountsUsecase) RolesDeleteRestore(ctx context.Context, req *apiv1.RolesDeleteRestoreRequest) (*apiv1.RolesDeleteRestoreResponse, error) {
	err := u.repo.RolesDeleteRestore(ctx, req.Records)
	if err != nil {
		return nil, err
	}
	return &apiv1.RolesDeleteRestoreResponse{}, nil
}
func (u *AccountsUsecase) RolesList(ctx context.Context) (*apiv1.RolesListResponse, error) {
	roles, err := u.repo.RolesList(ctx)
	if err != nil {
		return nil, err
	}
	response := u.adapter.RolesListGrpcFromSql(roles)
	return response, nil
}
func (u *AccountsUsecase) RoleCreateUpdate(ctx context.Context, req *apiv1.RoleCreateUpdateRequest) (*apiv1.RoleCreateUpdateResponse, error) {

	roleCreateParams := u.adapter.RoleCreateUpdateSqlFromGrpc(req)

	role, err := u.repo.RoleCreateUpdate(ctx, *roleCreateParams)

	if err != nil {
		return nil, err
	}
	response := u.adapter.RoleCreateUpdateGrpcFromSql(role)
	return response, nil
}
