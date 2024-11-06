package usecase

import (
	"context"

	"github.com/esolveeg/cms-api/db"
	"github.com/esolveeg/cms-api/pkg/redisclient"
	devkitv1 "github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	"github.com/rs/zerolog/log"
	"github.com/supabase-community/auth-go/types"
)

func (u *AccountsUsecase) userGenerateTokens(username string, userID int32) (*devkitv1.LoginInfo, error) {
	accessToken, accessPayload, err := u.tokenMaker.CreateToken(username, userID, u.tokenDuration)
	if err != nil {
		return nil, err
	}
	return &devkitv1.LoginInfo{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (u *AccountsUsecase) AppLogin(ctx context.Context, loginCode string) (*devkitv1.UserLoginResponse, redisclient.PermissionsMap, error) {
	user, err := u.repo.UserFind(ctx, db.UserFindParams{SearchKey: loginCode})
	if err != nil {
		return nil, nil, err
	}
	permissions, err := u.repo.UserPermissionsMap(ctx, user.UserID)
	if err != nil {
		return nil, nil, err
	}
	response := u.adapter.UserLoginGrpcFromSql(user)
	permissionsMap, err := u.redisClient.AuthSessionCreate(ctx, loginCode, permissions)
	if err != nil {
		return nil, nil, err
	}
	return response, permissionsMap, nil
}

func (u *AccountsUsecase) UserLogin(ctx context.Context, req *devkitv1.UserLoginRequest) (*devkitv1.UserLoginResponse, error) {
	userFindParams, supabaseRequest := u.adapter.UserLoginSqlFromGrpc(req)
	_, err := u.supaapi.AuthClient.Token(*supabaseRequest)
	if err != nil {
		log.Debug().Interface("supa err here", err).Msg("error")
		return nil, err
	}
	response, _, err := u.AppLogin(ctx, userFindParams.SearchKey)
	if err != nil {
		return nil, err
	}

	loginInfo, err := u.userGenerateTokens(req.LoginCode, response.User.UserId)
	if err != nil {
		return nil, err
	}
	response.LoginInfo = loginInfo

	if response.User.UserTypeId == 1 {
		navigationBar, err := u.repo.UserFindNavigationBars(ctx, response.User.UserId)
		if err != nil {
			return nil, err
		}
		navigations, err := u.adapter.UserFindNavigationBarsGrpcFromSql(&navigationBar)
		if err != nil {
			return nil, err
		}
		response.NavigationBar = *navigations
	}

	return response, nil
}

func (u *AccountsUsecase) UserLoginProvider(ctx context.Context, req *devkitv1.UserLoginProviderRequest) (*devkitv1.UserLoginProviderResponse, error) {
	resp, err := u.supaapi.ProviderLogin(types.Provider(req.Provider), req.RedirectUrl)
	if err != nil {
		return nil, err
	}
	return &devkitv1.UserLoginProviderResponse{Url: resp.AuthorizationURL}, nil
}

func (u *AccountsUsecase) UserInvite(ctx context.Context, req *devkitv1.UserInviteRequest) (*devkitv1.UserInviteResponse, error) {
	_, err := u.supaapi.AuthClient.Invite(types.InviteRequest{Email: req.UserEmail})
	if err != nil {
		return nil, err
	}
	return &devkitv1.UserInviteResponse{Message: "invitation sent"}, nil
}

func (u *AccountsUsecase) UserResetPassword(ctx context.Context, req *devkitv1.UserResetPasswordRequest) (*devkitv1.UserResetPasswordResponse, error) {
	if len(req.ResetToken) == 6 {
		resp, err := u.supaapi.AuthClient.VerifyForUser(*u.adapter.UserResetPasswordSupaFromGrpc(req))
		if err != nil {
			return nil, err
		}
		req.ResetToken = resp.AccessToken
	}
	user, err := u.supaapi.AuthClient.WithToken(req.ResetToken).GetUser()
	if err != nil {
		return nil, err
	}
	_, err = u.supaapi.AuthClient.AdminUpdateUser(types.AdminUpdateUserRequest{UserID: user.ID, Email: req.Email, Password: req.NewPassword})
	if err != nil {
		return nil, err
	}
	return &devkitv1.UserResetPasswordResponse{}, nil
}

func (u *AccountsUsecase) UserResetPasswordEmail(ctx context.Context, req *devkitv1.UserResetPasswordEmailRequest) (*devkitv1.UserResetPasswordEmailResponse, error) {
	err := u.supaapi.AuthClient.Recover(types.RecoverRequest{Email: req.Email})
	if err != nil {
		return nil, err
	}
	return &devkitv1.UserResetPasswordEmailResponse{}, nil
}

func (u *AccountsUsecase) UserLoginProviderCallback(ctx context.Context, req *devkitv1.UserLoginProviderCallbackRequest) (*devkitv1.UserLoginResponse, error) {
	user, err := u.supaapi.AuthClient.WithToken(req.AccessToken).GetUser()
	if err != nil {
		return nil, err
	}
	resp, _, err := u.AppLogin(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
