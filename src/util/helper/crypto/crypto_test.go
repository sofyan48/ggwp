package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCryptoSure(t *testing.T) {
	assert := assert.New(t)
	cipher, _ := HashPassword("12345678")
	decryptPass := CheckPasswordHash("12345678", cipher)
	assert.Equal(true, decryptPass, "No")
}

func TestCryptoNotSure(t *testing.T) {
	assert := assert.New(t)
	cipher, _ := HashPassword("1234567")
	decryptPass := CheckPasswordHash("12345678", cipher)
	assert.NotEqual(true, decryptPass, "No")
}
