package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oreshkanet/gtc-ya-dialogs/alice/api"
	"github.com/oreshkanet/gtc-ya-dialogs/alice/domain"
	"github.com/oreshkanet/gtc-ya-dialogs/alice/errors"
	"net/http"
)

type Service interface {
	GetUser(ctx context.Context, id domain.UserID) (*domain.User, errors.Err)
	GetUserForToken(ctx context.Context, token string) (*domain.User, errors.Err)
	AuthenticateAlice(ctx context.Context, req *api.Request) (context.Context, errors.Err)
}

var _ Service = &service{}

type service struct {
	oauthClientID string
	oauthSecret   string
	domain        string
	httpClient    http.Client
}

func NewService(deps Deps) (Service, errors.Err) {
	conf := deps.GetConfig()
	//secConf := deps.GetSecureConfig()
	return &service{
		oauthClientID: conf.OAuthClientID,
		//oauthSecret:   secConf.OAuthSecret,
		domain: conf.Domain,
		//repo:          deps.GetRepository(),
		//txMgr:         deps.GetTxManager(),
	}, nil
}

func (s *service) GetUser(ctx context.Context, id domain.UserID) (*domain.User, errors.Err) {
	var user *domain.User
	/*
		err := s.txMgr.InTx(ctx, db.TxRO()).Do(func(ctx context.Context) error {
			var err error
			user, err = s.repo.GetUser(ctx, id)
			return err
		})
		if err != nil {
			return nil, err
		}
	*/
	if user == nil {
		return nil, errors.NewUnauthenticated()
	}
	return user, nil
}

func (s *service) GetUserForToken(ctx context.Context, token string) (*domain.User, errors.Err) {
	user, err := s.getYandexUser(ctx, token)
	/*
		if err != nil {
			return nil, err
		}
		err = s.txMgr.InTx(ctx).Do(func(ctx context.Context) error {
			return s.repo.SaveUser(ctx, user)
		})
	*/
	if err != nil {
		return nil, errors.NewInternal(err)
	}
	return user, nil
}

func (s *service) AuthenticateAlice(ctx context.Context, req *api.Request) (context.Context, errors.Err) {
	if req.Session.User == nil {
		return ctx, errors.NewUnauthenticated()
	}
	if req.Session.User.Token == "" {
		return ctx, errors.NewUnauthenticated()
	}
	user, err := s.GetUserForToken(ctx, req.Session.User.Token)
	if err != nil {
		return ctx, err
	}
	return s.ctxWithUser(ctx, user), nil
}

func (s *service) getYandexUser(ctx context.Context, oauthToken string) (*domain.User, error) {
	req, err := http.NewRequest(http.MethodGet, "https://login.yandex.ru/info?format=json", nil)
	if err != nil {
		return nil, errors.NewInternal(err)
	}
	req = req.WithContext(ctx)
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", oauthToken))
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewInternal(err) //TODO
	}
	if resp.StatusCode != 200 {
		return nil, errors.NewInternal(fmt.Errorf("passport bad status: %d", resp.StatusCode)) //TODO
	}

	var res yandexAuthInfoDTO
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, errors.NewInternal(fmt.Errorf("parsing passport response: %w", err))
	}
	if len(res.Login) == 0 {
		return nil, errors.NewInternal(fmt.Errorf("login not found in passport response"))
	}
	return &domain.User{
		ID:    domain.UserID(res.Login),
		Name:  res.Login,
		Email: res.DefaultEmail,
		//Phone: res.DefaultPhone.Number,
		Avatar: res.Avatar,
	}, nil
}

func (s *service) ctxWithUser(ctx context.Context, user *domain.User) context.Context {
	return context.WithValue(ctx, "alice_user", user)
}
