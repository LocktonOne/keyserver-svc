package handlers

import (
	"errors"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/keyserver-svc/internal/data"
	"net/http"
)

func RequestVerification(w http.ResponseWriter, r *http.Request) {
	walletID := chi.URLParam(r, "wallet-id")
	if walletID == "" {
		Log(r).Error("walletID not specified")
		ape.RenderErr(w, problems.BadRequest(errors.New("walletID not specified"))...)
		return
	}

	wallet, err := WalletsQ(r).FilterByWalletID(walletID).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get wallet")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if wallet == nil {
		Log(r).Debug("wallet not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	token, err := EmailTokensQ(r).Get(wallet.Email)
	if err != nil {
		Log(r).WithError(err).Error("failed to get email token")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if token == nil {
		err = EmailTokensQ(r).Create(data.EmailToken{
			Token:     TFAConfig(r).Token(),
			Confirmed: WalletsConfig(r).DisableConfirm,
			Email:     wallet.Email,
		})
		if err != nil {
			Log(r).WithError(err).Error("failed to create email token")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	}

	if token.Confirmed {
		Log(r).Debug("email already confirmed")
		ape.RenderErr(w, problems.BadRequest(errors.New("email already confirmed"))...)
		return
	}

	err = EmailTokensQ(r).MarkUnsent(token.Id)
	if err != nil {
		Log(r).WithError(err).Error("failed to mark email token unsent")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
