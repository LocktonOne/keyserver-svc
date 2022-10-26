package handlers

import (
	"encoding/base64"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/keyserver-svc/internal/data"
	"gitlab.com/tokene/keyserver-svc/internal/service/helpers"
	"gitlab.com/tokene/keyserver-svc/internal/service/requests"
	"net/http"
)

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewChangePasswordRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to unmarshal request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	kdf, err := KDFQ(r).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get kdf-params")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	oldWallet, err := WalletsQ(r).FilterByEmail(request.Attributes.Email).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get wallet")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if oldWallet == nil {
		Log(r).Error("wallet not found by email")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	oldSalt, err := base64.StdEncoding.DecodeString(oldWallet.Salt)
	if err != nil {
		Log(r).WithError(err).Error("failed to decode salt")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	oldWalletID, err := helpers.GenerateWalletID(
		[]byte(request.Attributes.Email),
		[]byte(request.Attributes.OldPassword),
		oldSalt,
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

	if oldWalletID != oldWallet.WalletId {
		Log(r).Error("wallet id mismatch")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	newWallet := data.Wallet{
		Id:           oldWallet.Id,
		WalletId:     request.Attributes.WalletId,
		Email:        oldWallet.Email,
		KeychainData: request.Attributes.KeychainData,
		Salt:         request.Attributes.Salt,
		Verified:     oldWallet.Verified,
	}

	err = WalletsQ(r).Update(newWallet)
	if err != nil {
		Log(r).WithError(err).Error("failed to update wallet")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
