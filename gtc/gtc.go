// Интеграция с контроллером управления системой вентиляции

package gtc

import (
	"context"
	"github.com/oreshkanet/gtc-ya-dialogs/gtc/domain"
)

type GTC interface {
	GetDeviceList(ctx context.Context) []*domain.Device
}

type gtc struct {
}

func NewGtc() (GTC, error) {
	return &gtc{}, nil
}
