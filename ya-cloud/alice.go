package main

import (
	"context"
	"github.com/oreshkanet/gtc-ya-dialogs/alice"
	"github.com/oreshkanet/gtc-ya-dialogs/alice/config"
	"github.com/oreshkanet/gtc-ya-dialogs/logger"
	"go.uber.org/zap"
)

// aliceAppInstance - singleton-инстанс сервиса alice
var aliceAppInstance *alice.App

// initAliceApp - инициализация сервиса alice
func initAliceApp() (*alice.App, error) {
	ctx := context.Background()

	ctx, err := initLogging(ctx)
	if err != nil {
		return nil, err
	}
	logger.Info(ctx, "initializing alice app")

	gtc, err := initGtcApp(ctx)
	if err != nil {
		return nil, err
	}

	aliceAppInstance, err = alice.NewApp(ctx, config.LoadFromEnv(), gtc)
	if err != nil {
		return nil, err
	}

	return aliceAppInstance, nil
}

func initLogging(ctx context.Context) (context.Context, error) {
	instanceID := "GTC"

	zapConf := zap.NewProductionConfig()
	zapConf.Level.Enabled(zap.DebugLevel)
	zapConf.OutputPaths = []string{"stderr"}
	log, err := zapConf.Build(zap.AddCallerSkip(3))
	if err != nil {
		return nil, err
	}
	log = log.With(zap.String("instanceID", instanceID))
	return logger.CtxWithLogger(ctx, log), nil
}

func getAliceApp() (*alice.App, error) {
	if aliceAppInstance == nil {
		return initAliceApp()
	}
	return aliceAppInstance, nil
}
