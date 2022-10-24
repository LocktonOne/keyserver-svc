package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/keyserver-svc/resources"
	"net/http"
)

func NewCreateWalletRequest(r *http.Request) (resources.CreateWalletRequest, error) {
	var request resources.CreateWalletRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}
