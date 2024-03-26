package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oreshkanet/gtc-ya-dialogs/internal/domain"
	"github.com/oreshkanet/gtc-ya-dialogs/internal/errors"
	"net/http"
)

type Service interface {
	GetUser(ctx context.Context, id domain.UserID) (*domain.User, errors.Err)
	GetUserForToken(ctx context.Context, token string) (*domain.User, error)
}

var _ Service = &service{}

type service struct {
	oauthClientID string
	oauthSecret   string
	domain        string
	httpClient    http.Client
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

func (s *service) GetUserForToken(ctx context.Context, token string) (*domain.User, error) {
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
		return nil, err
	}
	return user, nil
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
