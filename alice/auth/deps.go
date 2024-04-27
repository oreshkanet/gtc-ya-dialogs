package auth

import "github.com/oreshkanet/gtc-ya-dialogs/alice/config"

type Deps interface {
	GetConfig() *config.Config
	//GetSecureConfig() *secure.Config
	//GetRepository() db.Repository
	//GetTxManager() db.TxManager
}
