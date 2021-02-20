package vault

import (
	"github.com/GrolimundSolutions/2vault/domain/vault"
	"github.com/GrolimundSolutions/2vault/services"
	"github.com/GrolimundSolutions/2vault/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var vault vault.Vault
	if err := c.ShouldBindJSON(&vault); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateVault(vault)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	/*
		// Decrypt secret data encrypted in the previos step
		str, errs := ansible.Decrypt(result.Secret, "qwert1234")
		if errs != nil {
			panic(errs)
		}
		fmt.Println("decrypted data:", str)

	*/

	c.JSON(http.StatusCreated, result)
}
