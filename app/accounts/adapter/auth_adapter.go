package adapter

import (
	"strings"

	"github.com/esolveeg/cms-api/db"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	"github.com/supabase-community/auth-go/types"
)

func (a *AccountsAdapter) UserLoginSqlFromGrpc(req *devkitv1.UserLoginRequest) (*db.UserFindParams, *types.TokenRequest) {
	isEmail := strings.Contains(req.LoginCode, "@") && strings.Contains(req.LoginCode, ".")
	supabseRequest := &types.TokenRequest{Password: req.UserPassword}
	if isEmail {
		supabseRequest.Email = req.LoginCode
	} else {
		supabseRequest.Phone = req.LoginCode
	}
	supabseRequest.GrantType = "password"
	return &db.UserFindParams{
		SearchKey: req.LoginCode,
	}, supabseRequest
}

func (a *AccountsAdapter) UserLoginGrpcFromSql(resp *db.AccountsSchemaUser) *devkitv1.UserLoginResponse {
	return &devkitv1.UserLoginResponse{
		User: a.UserEntityGrpcFromSql(resp),
	}
}
func (a *AccountsAdapter) UserResetPasswordSupaFromGrpc(req *devkitv1.UserResetPasswordRequest) *types.VerifyForUserRequest {
	return &types.VerifyForUserRequest{
		Type:       types.VerificationTypeRecovery,
		Token:      req.ResetToken,
		Email:      req.Email,
		RedirectTo: req.RedirectUrl,
	}
}
