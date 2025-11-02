package tokens_tests

import (
	"github.com/Azat201003/summorist-users/internal/tokens"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidationOk(t *testing.T) {
	token, err := tokens.GenerateToken(2, "../../")
	assert.NoError(t, err)
	id, err := tokens.ValidateToken(token, "../../")
	assert.NoError(t, err)
	assert.Equal(t, id, uint64(2))
}

func TestValidationError(t *testing.T) {
	_, err := tokens.ValidateToken("abrakadabra", "../../")
	assert.Error(t, err)
}
