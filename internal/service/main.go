package service

import (
	"gitlab.com/tokene/keyserver-svc/internal/data"
	"gitlab.com/tokene/keyserver-svc/internal/data/postgres"
	"net"
	"net/http"

	"gitlab.com/tokene/keyserver-svc/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log         *logan.Entry
	copus       types.Copus
	listener    net.Listener
	config      config.Config
	emailTokens data.EmailTokensQ
	wallets     data.WalletsQ
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	go s.runVerificationSender()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:         cfg.Log(),
		copus:       cfg.Copus(),
		listener:    cfg.Listener(),
		config:      cfg,
		emailTokens: postgres.NewEmailTokensQ(cfg.DB()),
		wallets:     postgres.NewWalletsQ(cfg.DB()),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
