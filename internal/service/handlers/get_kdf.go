package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/keyserver-svc/resources"
	"net/http"
	"strconv"
)

func GetKDF(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	kdf, err := KDFQ(r).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get kdf-params")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if email != "" { // load wallet KDF
		wallet, err := WalletsQ(r).FilterByEmail(email).Get()
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

		kdf.Salt = &wallet.Salt
	}

	response := resources.KdfResponse{
		Data: resources.Kdf{
			Key: resources.Key{
				Type: resources.KDF,
				ID:   strconv.Itoa(kdf.Version),
			},
			Attributes: resources.KdfAttributes{
				Algorithm: kdf.Algorithm,
				Bits:      int32(kdf.Bits),
				N:         int32(kdf.N),
				R:         int32(kdf.R),
				P:         int32(kdf.P),
				Salt:      kdf.Salt,
			},
		},
	}

	ape.Render(w, response)
}
