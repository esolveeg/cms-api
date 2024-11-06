package usecase

import (
	"context"
	"time"

	"github.com/esolveeg/cms-api/app/accounts/adapter"
	"github.com/esolveeg/cms-api/app/accounts/repo"
	"github.com/esolveeg/cms-api/db"
	"github.com/esolveeg/cms-api/pkg/auth"
	"github.com/esolveeg/cms-api/pkg/redisclient"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	supaapigo "github.com/darwishdev/supaapi-go"
)

type AccountsUsecaseInterface interface {
	UserLogin(ctx context.Context, req *devkitv1.UserLoginRequest) (*devkitv1.UserLoginResponse, error)
	UsersDeleteRestore(ctx context.Context, req *devkitv1.UsersDeleteRestoreRequest) (*devkitv1.UsersDeleteRestoreResponse, error)
	UsersList(ctx context.Context) (*devkitv1.UsersListResponse, error)
	UserCreateUpdate(ctx context.Context, req *devkitv1.UserCreateUpdateRequest) (*devkitv1.UserCreateUpdateResponse, error)
	UserLoginProvider(ctx context.Context, req *devkitv1.UserLoginProviderRequest) (*devkitv1.UserLoginProviderResponse, error)
	UserDelete(ctx context.Context, userID int32) (*devkitv1.AccountsSchemaUser, error)
	UserInvite(ctx context.Context, req *devkitv1.UserInviteRequest) (*devkitv1.UserInviteResponse, error)
	RolesDeleteRestore(ctx context.Context, req *devkitv1.RolesDeleteRestoreRequest) (*devkitv1.RolesDeleteRestoreResponse, error)
	RolesList(ctx context.Context) (*devkitv1.RolesListResponse, error)
	AppLogin(ctx context.Context, loginCode string) (*devkitv1.UserLoginResponse, redisclient.PermissionsMap, error)
	UserResetPassword(ctx context.Context, req *devkitv1.UserResetPasswordRequest) (*devkitv1.UserResetPasswordResponse, error)
	UserResetPasswordEmail(ctx context.Context, req *devkitv1.UserResetPasswordEmailRequest) (*devkitv1.UserResetPasswordEmailResponse, error)
	RoleCreateUpdate(ctx context.Context, req *devkitv1.RoleCreateUpdateRequest) (*devkitv1.RoleCreateUpdateResponse, error)
}

type AccountsUsecase struct {
	store         db.Store
	adapter       adapter.AccountsAdapterInterface
	tokenMaker    auth.Maker
	tokenDuration time.Duration
	supaapi       supaapigo.Supaapi
	redisClient   redisclient.RedisClientInterface
	repo          repo.AccountsRepoInterface
}

func NewAccountsUsecase(store db.Store, supaapi supaapigo.Supaapi, redisClient redisclient.RedisClientInterface, tokenMaker auth.Maker, tokenDuration time.Duration) AccountsUsecaseInterface {
	return &AccountsUsecase{
		supaapi:       supaapi,
		tokenMaker:    tokenMaker,
		redisClient:   redisClient,
		tokenDuration: tokenDuration,
		store:         store,
		adapter:       adapter.NewAccountsAdapter(),
		repo:          repo.NewAccountsRepo(store),
	}
}
