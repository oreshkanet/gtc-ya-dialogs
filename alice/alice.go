package alice

import (
	"context"
	"github.com/oreshkanet/gtc-ya-dialogs/alice/api"
	"github.com/oreshkanet/gtc-ya-dialogs/alice/auth"
	"github.com/oreshkanet/gtc-ya-dialogs/alice/config"
	"github.com/oreshkanet/gtc-ya-dialogs/gtc"
	"github.com/oreshkanet/gtc-ya-dialogs/logger"
	"go.uber.org/zap"
)

type Request = api.Request
type Response = api.Response

type App struct {
	ctx    context.Context
	cfg    *config.Config
	logger *zap.Logger

	auth auth.Service
	gtc  gtc.GTC
}

func NewApp(ctx context.Context, cfg *config.Config, gtc gtc.GTC) (*App, error) {
	var err error

	app := &App{
		ctx:    ctx,
		cfg:    cfg,
		gtc:    gtc,
		logger: logger.FromCtx(ctx),
	}

	app.auth, err = auth.NewService(app)
	if err != nil {
		return nil, err
	}

	return app, err
}

func (a *App) GetConfig() *config.Config {
	assertInitialized(a.cfg, "config")
	return a.cfg
}

func (a *App) GetLogger() *zap.Logger {
	assertInitialized(a.logger, "logger")
	return a.logger
}

func (a *App) GetContext() context.Context {
	assertInitialized(a.ctx, "ctx")
	return a.ctx
}
