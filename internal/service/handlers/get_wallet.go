package handlers

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/keyserver-svc/resources"
	"net/http"
)

func GetWallet(w http.ResponseWriter, r *http.Request) {
	walletID := chi.URLParam(r, "wallet-id")

	wallet, err := WalletsQ(r).FilterByWalletID(walletID).Get()
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

	if !wallet.Verified {
		problem := problems.Forbidden()
		problem.Code = "verification_required"
		ape.RenderErr(w, problem)
		return
	}

	response := resources.WalletResponse{
		Data: resources.Wallet{
			Key: resources.NewKeyInt64(wallet.Id, resources.WALLET),
			Attributes: resources.WalletAttributes{
				WalletId:     wallet.WalletId,
				Email:        wallet.Email,
				KeychainData: wallet.KeychainData,
				Salt:         wallet.Salt,
				Verified:     wallet.Verified,
			},
		},
	}

	ape.Render(w, response)
}
