package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/keyserver-svc/resources"
	"net/http"
)

type CreateWalletRequest struct {
	Data resources.CreateWalletRequest `json:"data"`
}

func NewCreateWalletRequest(r *http.Request) (resources.CreateWalletRequest, error) {
	var request CreateWalletRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		return request.Data, errors.Wrap(err, "failed to unmarshal")
	}

	return request.Data, request.validate()
}

func (r CreateWalletRequest) validate() error {
	return validation.Errors{
		"/data/attributes/wallet_id": validation.Validate(&r.Data.Attributes.WalletId, validation.Required),
		"/data/attributes/email":     validation.Validate(&r.Data.Attributes.Email, validation.Required),
	}.Filter()
}
