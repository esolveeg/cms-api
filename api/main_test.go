package api

import (
	"context"
	"os"
	"testing"

	"github.com/esolveeg/cms-api/config"
	"github.com/esolveeg/cms-api/db"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1/devkitv1connect"
	"github.com/rs/zerolog/log"
)

var realDbApi devkitv1connect.DevkitServiceHandler
var testConfig config.Config

func newTestApi(store db.Store) devkitv1connect.DevkitServiceHandler {
	api, err := NewApi(testConfig, store)
	if err != nil {
		log.Fatal().Err(err).Msg("canot start the api")
	}
	return api
}
func TestMain(m *testing.M) {
	_, err := config.LoadState("../config")
	if err != nil {
		panic(err)
	}
	testConfig, err = config.LoadConfig("../config", "test")
	if err != nil {
		panic(err)
	}

	store, _, err := db.InitDB(context.Background(), testConfig.DBSource, false)
	if err != nil {
		log.Fatal().Err(err).Msg("canot start the api")
	}
	realDbApi, err = NewApi(testConfig, store)
	if err != nil {
		log.Fatal().Err(err).Msg("canot start the api")
	}

	os.Exit(m.Run())
}
