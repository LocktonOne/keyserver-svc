package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/keyserver-svc/resources"
	"net/http"
)

type ChangePasswordRequest struct {
	Data resources.ChangePasswordRequest `json:"data"`
}

func NewChangePasswordRequest(r *http.Request) (resources.ChangePasswordRequest, error) {
	var request ChangePasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		return request.Data, errors.Wrap(err, "failed to unmarshal")
	}

	return request.Data, request.validate()
}

func (r ChangePasswordRequest) validate() error {
	return validation.Errors{
		"/data/attributes/wallet_id": validation.Validate(&r.Data.Attributes.WalletId, validation.Required),
		"/data/attributes/email":     validation.Validate(&r.Data.Attributes.Email, validation.Required),
	}.Filter()
}
