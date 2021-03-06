package vault

import "github.com/GrolimundSolutions/2vault/utils/errors"

type Vault struct {
	Text   string `json:"text,omitempty"`
	Secret string `json:"secret,omitempty"`
}

func (vault *Vault) Validate() *errors.RestErr {
	if vault.Text == "" {
		return errors.NewBadRequestError("Text must not be empty")
	}

	return nil
}
