package api

import (
	// USECASE_IMPORTS
	companiesUsecase "github.com/esolveeg/cms-api/app/companies/usecase"

	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/sqlseeder"
	supaapigo "github.com/darwishdev/supaapi-go"
	accountsUsecase "github.com/esolveeg/cms-api/app/accounts/usecase"
	publicUsecase "github.com/esolveeg/cms-api/app/public/usecase"
	"github.com/esolveeg/cms-api/config"
	"github.com/esolveeg/cms-api/db"
	"github.com/esolveeg/cms-api/pkg/auth"
	"github.com/esolveeg/cms-api/pkg/redisclient"
	"github.com/esolveeg/cms-api/pkg/resend"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1/devkitv1connect"
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
	companiesUsecase companiesUsecase.CompaniesUsecaseInterface

	supaapi     supaapigo.Supaapi
	redisClient redisclient.RedisClientInterface
	store       db.Store
}

func HashFunc(req string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req), bcrypt.DefaultCost)
	return string(hashedPassword)
}
func NewApi(config config.Config, store db.Store) (devkitv1connect.DevkitServiceHandler, error) {
	validator, err := protovalidate.New()
	if err != nil {
		return nil, err
	}
	resendClient, err := resend.NewResendService(config.ResendApiKey, config.ClientBaseUrl)

	supaapi := supaapigo.NewSupaapi(supaapigo.SupaapiConfig{
		ProjectRef:     config.DBProjectREF,
		Env:            supaapigo.DEV,
		Port:           config.SupaAPIPort,
		ServiceRoleKey: config.SupabaseServiceRoleKey,
		ApiKey:         config.SupabaseApiKey,
	})
	tokenMaker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		panic("cann't create paset maker in gapi/api.go")
	}
	sqlSeeder := sqlseeder.NewSeeder(sqlseeder.SeederConfig{HashFunc: HashFunc})
	redisClient := redisclient.NewRedisClient(config.RedisHost, config.RedisPort, config.RedisPassword, config.RedisDatabase)
	// USECASE_INSTANTIATIONS
	companiesUsecase := companiesUsecase.NewCompaniesUsecase(store)
	accountsUsecase := accountsUsecase.NewAccountsUsecase(store, supaapi, redisClient, tokenMaker, config.AccessTokenDuration)
	publicUsecase := publicUsecase.NewPublicUsecase(store, supaapi, redisClient, resendClient)
	return &Api{
		// USECASE_INJECTIONS
		companiesUsecase: companiesUsecase,

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
