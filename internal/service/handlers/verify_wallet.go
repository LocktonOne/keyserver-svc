package handlers

import (
	"errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/keyserver-svc/internal/service/requests"
	"net/http"
)

func VerifyWallet(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewVerifyWalletRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to unmarshal request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	wallet, err := WalletsQ(r).FilterByWalletID(request.WalletID).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get wallet")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if wallet == nil {
		Log(r).Error("wallet not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	verified, err := EmailTokensQ(r).Verify(wallet.Email, request.Data.Attributes.Token)
	if err != nil {
		Log(r).WithError(err).Error("failed to verify email token")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if !verified {
		Log(r).Debug("invalid token")
		ape.RenderErr(w, problems.BadRequest(errors.New("invalid token"))...)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
