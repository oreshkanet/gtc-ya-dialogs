package alice

import (
	"context"
	"github.com/oreshkanet/gtc-ya-dialogs/alice/api"
	"github.com/oreshkanet/gtc-ya-dialogs/alice/errors"
	"github.com/oreshkanet/gtc-ya-dialogs/logger"
	"go.uber.org/zap"
)

func (a *App) Handle(ctx context.Context, req *api.Request) (*api.Response, error) {
	sessionID := req.Session.SessionID
	ctx = logger.CtxWithLogger(ctx, a.logger.With(zap.String("sessionID", string(sessionID))))
	//ctx = cache.ContextWithCache(ctx)
	ctx, err := a.auth.AuthenticateAlice(ctx, req)
	if err != nil {
		if err.GetCode() == errors.CodeUnauthenticated {
			return &api.Response{
				Version:             req.Version,
				StartAccountLinking: &api.EmptyObj{},
			}, nil
		}
		return a.reportError(ctx, err)
	}
	resp, err := a.handle(ctx, req)
	if err != nil {
		return a.reportError(ctx, err)
	}
	resp.Version = req.Version
	return resp, nil
}

func (a *App) handle(ctx context.Context, req *api.Request) (*api.Response, errors.Err) {
	//req.Request.Command

	if req.Session.New || req.AccountLinkingComplete != nil {
		return &api.Response{Response: &api.Resp{
			Text: "Давайте я помогу вам со списками!",
		}}, nil
	}
	/*
		if state := req.State.Session; state.State != api.StateInit {
			intents := req.Request.NLU.Intents
			if req.Request.Type == api.RequestTypeSimple && intents.Cancel != nil || intents.Reject != nil {
				return &api.Response{
					Response: &api.Resp{Text: "Чем я могу помочь?"},
				}, nil
			}
			scenario, ok := a.stateScenarios[state.State]
			if ok {
				resp, err := scenario(ctx, req)
				if err != nil {
					return nil, err
				}
				if resp != nil {
					return resp, nil
				}
			}
		}
		for _, s := range a.scratchScenarios {
			resp, err := s(ctx, req)
			if err != nil {
				return nil, err
			}
			if resp != nil {
				return resp, err
			}
		}
	*/
	return &api.Response{Response: &api.Resp{
		Text: "Я вас не поняла",
	}}, nil
}

func (a *App) reportError(ctx context.Context, err errors.Err) (*api.Response, error) {
	//errors.Log(ctx, err)
	a.logger.Error(err.Error())
	return nil, err
}

func NewRequest() *Request {
	return &Request{
		Session: api.Session{
			User: &api.User{},
		},
	}
}
