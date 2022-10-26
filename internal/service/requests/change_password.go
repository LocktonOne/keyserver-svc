package requests

import (
	"encoding/json"
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

	return request.Data, nil
}
