package gtc

import (
	"context"
	"github.com/oreshkanet/gtc-ya-dialogs/gtc/domain"
)

func (g *gtc) GetDeviceList(ctx context.Context) []*domain.Device {
	return []*domain.Device{
		{Id: "1",
			Name:        "GTC_вентиляция",
			Description: "GTC_вентиляция (+-)",
		},
	}
}
