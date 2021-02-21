package services

import (
	"fmt"
	"github.com/GrolimundSolutions/2vault/domain/vault"
	"github.com/GrolimundSolutions/2vault/utils/errors"
	ansible "github.com/sosedoff/ansible-vault-go"
)

func CreateVault(vault vault.Vault) (*vault.Vault, *errors.RestErr) {
	if err := vault.Validate(); err != nil {
		return nil, err
	}
	// Encrypt secret data
	// TODO: Get Masterpassword from outside (Openshift Secret)
	str, err := ansible.Encrypt(vault.Text, "qwert1234")
	if err != nil {
		return nil, errors.NewInterlanServerError(
			fmt.Sprintf("Error when encrypting %s Error: %s", vault.Text, err))
	}
	vault.Secret = str
	return &vault, nil

}
