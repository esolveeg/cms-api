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

func (api *Api) authorizeRequestHeader(ctx context.Context, header http.Header) (*auth.Payload, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing metadata")
	}
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != "bearer" {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := api.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, err
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

func (api *Api) getAccessableActionsForGroup(ctx context.Context, header http.Header, group string) (*devkitv1.ListDataOptions, error) {
	payload, err := api.authorizeRequestHeader(ctx, header)
	if err != nil {
		return nil, err
	}
	permissionsMap, err := api.authorizedUserPermissions(ctx, payload)
	if err != nil {
		return nil, err
	}

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
