package main

import (
	"context"
	"github.com/oreshkanet/gtc-ya-dialogs/gtc"
)

func initGtcApp(ctx context.Context) (gtc.GTC, error) {
	return gtc.NewGtc()
}
