package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"connectrpc.com/connect"
	"github.com/esolveeg/cms-api/pkg/auth"
	"github.com/esolveeg/cms-api/pkg/redisclient"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	"github.com/iancoleman/strcase"
	"github.com/tangzero/inflector"
)

func (api *Api) authorizeRequestHeader(header http.Header) (*auth.Payload, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing metadata")
	}
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid authorization header format"))
	}

	authType := strings.ToLower(fields[0])
	if authType != "bearer" {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("unsupported authorization type: %s", authType))
	}

	accessToken := fields[1]
	payload, err := api.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}
	return payload, nil

}
func (api *Api) authorizedUserPermissions(ctx context.Context, payload *auth.Payload) (redisclient.PermissionsMap, error) {
	permissionsMap, err := api.redisClient.AuthSessionFind(ctx, payload.Username)
	if err != nil {
		_, permissionsMap, err = api.accountsUsecase.AppLogin(ctx, payload.Username)
		if err != nil {
			return nil, fmt.Errorf("can't load permissions: %s", err)

		}

	}
	return permissionsMap, nil

}

func (api *Api) checkForAccess(header http.Header, group string, permission string) (*redisclient.PermissionsMap, error) {
	userPayload, err := api.authorizeRequestHeader(header)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid access token: %w", err))
	}

	permissionMap, err := api.authorizedUserPermissions(context.Background(), userPayload)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get user permissions: %w", err))
	}

	groupPermissions, ok := permissionMap[group]
	if !ok {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("user does not have the required permissions"))
	}
	if permission != "list" && permission != "" {
		permissionName := strcase.ToCamel(fmt.Sprintf("%s_%s", group, permission))
		_, ok := groupPermissions[permissionName]
		if !ok {
			permissionName = strcase.ToCamel(fmt.Sprintf("%s_%s", inflector.Singularize(group), permission))
			_, ok = groupPermissions[permissionName]
			if !ok {
				return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("user does not have the required permissions"))
			}
		}
	}
	return &permissionMap, nil

}
func (api *Api) getAccessableActionsForGroup(permissionsMap redisclient.PermissionsMap, group string) (*devkitv1.ListDataOptions, error) {
	groupPermissions, ok := permissionsMap[group]
	if !ok {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("user don't have permission to this group"))
	}
	var (
		singularizedGroup            string = inflector.Singularize(group)
		findEndpoint                 string = strcase.ToLowerCamel(fmt.Sprintf("%s_find_for_update", singularizedGroup))
		findRequestProperty          string = "recordId"
		deleteRestoreRequestProperty string = "records"
		redirectRoute                string = fmt.Sprintf("%s_list", group)
		createUpdate                 string = fmt.Sprintf("%s_create_update", singularizedGroup)
		create                       string = fmt.Sprintf("%s_create", singularizedGroup)
		deleteRestore                string = fmt.Sprintf("%s_delete_restore", singularizedGroup)
		update                       string = fmt.Sprintf("%s_update", singularizedGroup)
	)
	response := devkitv1.ListDataOptions{
		Title:       fmt.Sprintf("%s_list", group),
		Description: fmt.Sprintf("%s_list", group),
	}
	_, ok = groupPermissions[strcase.ToCamel(create)]
	if ok {

		response.CreateHandler = &devkitv1.CreateHandler{
			RedirectRoute: redirectRoute,
			Title:         create,
			Endpoint:      strcase.ToLowerCamel(createUpdate),
			RouteName:     create,
		}
	}
	_, ok = groupPermissions[strcase.ToCamel(update)]
	if ok {
		response.UpdateHandler = &devkitv1.UpdateHandler{

			RedirectRoute:       redirectRoute,
			Title:               update,
			Endpoint:            strcase.ToLowerCamel(createUpdate),
			RouteName:           update,
			FindEndpoint:        findEndpoint,
			FindRequestProperty: findRequestProperty,
		}
	}
	_, ok = groupPermissions[strcase.ToCamel(deleteRestore)]
	if ok {
		response.DeleteRestoreHandler = &devkitv1.DeleteRestoreHandler{
			Endpoint:        strcase.ToLowerCamel(deleteRestore),
			RequestProperty: deleteRestoreRequestProperty,
		}
	}
	return &response, nil
}
