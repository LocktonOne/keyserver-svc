package handlers

import (
	"encoding/base64"
	"fmt"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/keyserver-svc/internal/service/helpers"
	"gitlab.com/tokene/keyserver-svc/internal/service/requests"
	"net/http"
)

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteWalletRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to unmarshal request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	wallet, err := WalletsQ(r).FilterByEmail(request.Attributes.Email).Get()
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

	kdf, err := KDFQ(r).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get kdf-params")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	salt, err := base64.StdEncoding.DecodeString(wallet.Salt)
	if err != nil {
		Log(r).WithError(err).Error("failed to decode salt")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	walletID, err := helpers.GenerateWalletID(
		[]byte(request.Attributes.Email),
		[]byte(request.Attributes.Password),
		salt,
		int(kdf.N),
		int(kdf.R),
		int(kdf.P),
		int(kdf.Bits/8),
	)

	if err != nil {
		Log(r).WithError(err).Error("failed to get old wallet id")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if walletID != wallet.WalletId {
		fmt.Println(walletID)
		Log(r).Error("wallet id mismatch")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	if !wallet.Verified {
		problem := problems.Forbidden()
		problem.Code = "verification_required"
		ape.RenderErr(w, problem)
		return
	}

	if err = WalletsQ(r).Delete(walletID); err != nil {
		Log(r).WithError(err).Error("failed to delete wallet")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
