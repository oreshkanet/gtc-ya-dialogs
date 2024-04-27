package main

import (
	"context"
	"github.com/oreshkanet/gtc-ya-dialogs/alice"
	"log"
	"testing"
)

func TestStartTask(t *testing.T) {
	req := alice.NewRequest()

	ctx := context.Background()
	res, err := Handler(ctx, req)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}
}
