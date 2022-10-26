package service

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/clienturl"
	"gitlab.com/tokene/keyserver-svc/internal/notificator"
	"time"
)

func (s *service) runVerificationSender() {
	log := s.config.Log().WithField("service", "verification-sender")
	if s.config.Notificator().IsDisabled() {
		log.Warn("disabled")
		return
	}
	go func() {
		log.Debug("starting")
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for ; ; <-ticker.C {
			err := sendVerifications(s, log)
			if err != nil {
				log.WithError(err).Error("Failed to send verifications")
			}
		}
	}()
}

func sendVerifications(app *service, log *logan.Entry) error {
	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("sendVerifications panicked")
		}
	}()

	tokensQ := app.emailTokens
	tokens, err := tokensQ.GetUnsent()
	if err != nil {
		return errors.Wrap(err, "failed to get unsent verifications")
	}

	for _, token := range tokens {
		wallet, err := app.wallets.FilterByEmail(token.Email).Get()
		if err != nil {
			log.WithError(err).WithField("email", token.Email).Warn("wallet with this email does not exist")
			continue
		}

		payload := clienturl.NewPayload(
			notificator.RedirectTypeEmailVerification,
			map[string]interface{}{
				"wallet_id": wallet.WalletId,
				"token":     token,
			})

		err = app.config.Notificator().SendVerificationLink(token.Email, payload)
		if err != nil {
			log.WithError(err).WithField("email", token.Email).Warn("failed to send verification link")
			continue
		}

		if err := tokensQ.MarkSent(token.Id); err != nil {
			return errors.Wrap(err, "failed to mark notification as sent")
		}
	}

	return nil
}