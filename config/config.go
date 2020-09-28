package config

import (
	"context"

	"github.com/10xLabs/chandler/config"
	"github.com/10xLabs/chandler/envdecoder"
	cstore "github.com/10xLabs/chandler/store"
	"github.com/10xLabs/log"

	"github.com/10xLabs/mixpanel-projector/store"
)

type cfg struct {
	Environment string `env:"ENVIRONMENT,required"`
	FileStore   struct {
		Dir string `env:"FILE_STORE_DIR,required"`
	}
	MixpanelToken string `env:"MIXPANEL_TOKEN,required"`
	Store         cstore.Store
	Log           log.Config
}

// App ...
var App cfg

// Setup ...
func Setup() {
	if err := envdecoder.Decode(context.TODO(), &App); err != nil {
		panic(err)
	}
	App.Store = store.NewFileStore()
	App.Log.DebugEnabled = App.Environment != config.ProductionEnvironment
	App.Log.PrettyPrint = App.Environment == config.DevelopmentEnvironment
}
