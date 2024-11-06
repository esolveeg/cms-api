package api

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	apiv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
)

func (api *Api) RolesList(ctx context.Context, req *connect.Request[apiv1.RolesListRequest]) (*connect.Response[apiv1.RolesListResponse], error) {
	permissionGroupName := "roles"
	permissionsMap, err := api.checkForAccess(req.Header(), permissionGroupName, "list")
	if err != nil {
		return nil, err
	}
	response, err := api.accountsUsecase.RolesList(ctx)
	if err != nil {
		return nil, err
	}
	options, err := api.getAccessableActionsForGroup(*permissionsMap, permissionGroupName)
	if err != nil {
		return nil, err
	}
	response.Options = options
	return connect.NewResponse(response), nil
}

func (api *Api) RoleCreateUpdate(ctx context.Context, req *connect.Request[apiv1.RoleCreateUpdateRequest]) (*connect.Response[apiv1.RoleCreateUpdateResponse], error) {
	err := api.validator.Validate(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	if req.Msg.GetRoleId() == 0 && req.Msg.GetRoleName() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("role name is required if role ID not passed (create scenario)"))
	}
	_, err = api.checkForAccess(req.Header(), "roles", "create_update")
	if err != nil {
		return nil, err
	}
	response, err := api.accountsUsecase.RoleCreateUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(response), nil
}

func (api *Api) RolesDeleteRestore(ctx context.Context, req *connect.Request[apiv1.RolesDeleteRestoreRequest]) (*connect.Response[apiv1.RolesDeleteRestoreResponse], error) {
	_, err := api.checkForAccess(req.Header(), "roles", "delete_restore")
	if err != nil {
		return nil, err
	}
	_, err = api.accountsUsecase.RolesDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&apiv1.RolesDeleteRestoreResponse{}), nil
}
