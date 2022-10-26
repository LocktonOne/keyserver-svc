package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/keyserver-svc/internal/data"
	"gitlab.com/tokene/keyserver-svc/internal/service/requests"
	"gitlab.com/tokene/keyserver-svc/resources"
	"net/http"
)

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateWalletRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	wallet, err := WalletsQ(r).FilterByEmail(request.Attributes.Email).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get wallet")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if wallet != nil {
		ape.RenderErr(w, problems.Conflict())
		return
	}

	id, err := WalletsQ(r).Create(data.Wallet{
		WalletId:     request.Attributes.WalletId,
		Email:        request.Attributes.Email,
		KeychainData: request.Attributes.KeychainData,
		Salt:         request.Attributes.Salt,
		Verified:     WalletsConfig(r).DisableConfirm,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to create wallet")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = EmailTokensQ(r).Create(data.EmailToken{
		Token:     TFAConfig(r).Token(),
		Confirmed: WalletsConfig(r).DisableConfirm,
		Email:     request.Attributes.Email,
	})
	if err != nil {
		Log(r).WithError(err).Warn("failed to create email token")
		return
	}

	response := resources.WalletResponse{
		Data: resources.Wallet{
			Key: resources.NewKeyInt64(id, resources.WALLET),
			Attributes: resources.WalletAttributes{
				WalletId:     request.Attributes.WalletId,
				Email:        request.Attributes.Email,
				KeychainData: request.Attributes.KeychainData,
				Salt:         request.Attributes.Salt,
				Verified:     WalletsConfig(r).DisableConfirm,
			},
		},
	}

	w.WriteHeader(http.StatusCreated)
	ape.Render(w, response)
}
