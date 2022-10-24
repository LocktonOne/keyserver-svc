package requests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/keyserver-svc/resources"
	"net/http"
)

type VerifyWalletRequest struct {
	WalletID string                        `json:"-"`
	Data     resources.VerifyWalletRequest `json:"data"`
}

func NewVerifyWalletRequest(r *http.Request) (VerifyWalletRequest, error) {
	var request VerifyWalletRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		return request, errors.Wrap(err, "failed to unmarshal")
	}

	request.WalletID = chi.URLParam(r, "wallet-id")

	return request, nil
}
