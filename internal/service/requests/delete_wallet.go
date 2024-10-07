package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/keyserver-svc/resources"
	"net/http"
)

type DeleteWalletRequest struct {
	Data resources.DeleteWalletRequest `json:"data"`
}

func NewDeleteWalletRequest(r *http.Request) (resources.DeleteWalletRequest, error) {
	var request DeleteWalletRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		return request.Data, errors.Wrap(err, "failed to unmarshal")
	}

	return request.Data, request.validate()
}

func (r DeleteWalletRequest) validate() error {
	return validation.Errors{
		"/data/attributes/password": validation.Validate(&r.Data.Attributes.Password, validation.Required),
		"/data/attributes/email":    validation.Validate(&r.Data.Attributes.Email, validation.Required),
	}.Filter()
}
