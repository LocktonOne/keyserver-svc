package service

import (
	"gitlab.com/tokene/keyserver-svc/internal/data/postgres"
	"gitlab.com/tokene/keyserver-svc/internal/service/handlers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxKDFQ(postgres.NewKDFQ(s.config.DB())),
			handlers.CtxWalletsQ(s.wallets),
			handlers.CtxWalletsConfig(s.config.WalletsConfig()),
			handlers.CtxEmailTokensQ(s.emailTokens),
			handlers.CtxTFAConfig(s.config.TFAConfig()),
		),
	)
	r.Route("/integrations/keyserver-svc", func(r chi.Router) {
		r.Route("/wallet", func(r chi.Router) {
			r.Post("/", handlers.CreateWallet)
			r.Route("/{wallet-id}", func(r chi.Router) {
				r.Get("/", handlers.GetWallet)

				r.Post("/verification", handlers.RequestVerification)
				r.Put("/verification", handlers.VerifyWallet)
			})

		})

		r.Get("/kdf", handlers.GetKDF)
	})

	return r
}
