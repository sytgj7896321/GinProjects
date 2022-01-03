package x25519

import (
	"filippo.io/age"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Secure struct{}

func (s *Secure) GenerateKeyPair(c *gin.Context) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"public_key": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"public_key": identity.Recipient().String(),
	})
}
