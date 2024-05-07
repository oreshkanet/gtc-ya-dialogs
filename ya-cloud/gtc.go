package main

import (
	"context"
	"github.com/oreshkanet/gtc-ya-dialogs/gtc"
)

// TODO: это тоже должен быть singleton

func initGtcApp(ctx context.Context) (gtc.GTC, error) {
	return gtc.NewGtc()
}
