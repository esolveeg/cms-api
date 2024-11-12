package api

import (
	// USECASE_IMPORTS
	"github.com/bufbuild/protovalidate-go"
	accountsUsecase "github.com/darwishdev/devkit-api/app/accounts/usecase"
	publicUsecase "github.com/darwishdev/devkit-api/app/public/usecase"
	"github.com/darwishdev/devkit-api/config"
	"github.com/darwishdev/devkit-api/db"
	"github.com/darwishdev/devkit-api/pkg/auth"
	"github.com/darwishdev/devkit-api/pkg/redisclient"
	"github.com/darwishdev/devkit-api/pkg/resend"
	"github.com/darwishdev/devkit-api/proto_gen/devkit/v1/devkitv1connect"
	"github.com/darwishdev/sqlseeder"
	supaapigo "github.com/darwishdev/supaapi-go"
	"golang.org/x/crypto/bcrypt"
)

type Api struct {
	devkitv1connect.UnimplementedDevkitServiceHandler
	accountsUsecase accountsUsecase.AccountsUsecaseInterface
	config          config.Config
	validator       *protovalidate.Validator
	tokenMaker      auth.Maker
	sqlSeeder       sqlseeder.SeederInterface
	publicUsecase   publicUsecase.PublicUsecaseInterface
	// USECASE_FIELDS
	supaapi     supaapigo.Supaapi
	redisClient redisclient.RedisClientInterface
	store       db.Store
}

func HashFunc(req string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req), bcrypt.DefaultCost)
	return string(hashedPassword)
}
func NewApi(config config.Config, store db.Store, tokenMaker auth.Maker, redisClient redisclient.RedisClientInterface, validator *protovalidate.Validator) (devkitv1connect.DevkitServiceHandler, error) {
	resendClient, err := resend.NewResendService(config.ResendApiKey, config.ClientBaseUrl)
	if err != nil {
		return nil, err
	}
	supaapi := supaapigo.NewSupaapi(supaapigo.SupaapiConfig{
		ProjectRef:     config.DBProjectREF,
		Env:            supaapigo.DEV,
		Port:           config.SupaAPIPort,
		ServiceRoleKey: config.SupabaseServiceRoleKey,
		ApiKey:         config.SupabaseApiKey,
	})
	sqlSeeder := sqlseeder.NewSeeder(sqlseeder.SeederConfig{HashFunc: HashFunc})
	// USECASE_INSTANTIATIONS
	accountsUsecase := accountsUsecase.NewAccountsUsecase(store, supaapi, redisClient, tokenMaker, config.AccessTokenDuration)
	publicUsecase := publicUsecase.NewPublicUsecase(store, supaapi, redisClient, resendClient)
	return &Api{
		// USECASE_INJECTIONS
		accountsUsecase: accountsUsecase,
		store:           store,
		redisClient:     redisClient,
		tokenMaker:      tokenMaker,
		supaapi:         supaapi,
		config:          config,
		sqlSeeder:       sqlSeeder,
		publicUsecase:   publicUsecase,
		validator:       validator,
	}, nil
}
