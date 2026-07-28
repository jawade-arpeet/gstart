package client

import (
	"gstart/internal/config"

	"github.com/resend/resend-go/v3"
)

type ResendClient struct {
	rsnd *resend.Client
}

func newResendClient(cfg *config.ResendConfig) *ResendClient {
	return &ResendClient{
		rsnd: resend.NewClient(cfg.ApiKey),
	}
}
