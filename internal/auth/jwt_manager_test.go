package auth

import (
	"testing"
	"time"

	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testUser() model.User {
	return model.User{ID: 1, Email: "aisha@example.com", Role: "student"}
}

func TestJWTManager_AccessToken_RoundTrip(t *testing.T) {
	jm := NewJWTManager("test-secret", time.Minute, time.Hour)
	user := testUser()

	token, expiresAt, err := jm.NewAccessToken(user)
	require.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Greater(t, expiresAt, time.Now().Unix())

	parsedUser, err := jm.ParseAccessToken(token)
	require.NoError(t, err)
	assert.Equal(t, user.ID, parsedUser.ID)
	assert.Equal(t, user.Email, parsedUser.Email)
	assert.Equal(t, user.Role, parsedUser.Role)
}

func TestJWTManager_ParseAccessToken_ExpiredToken(t *testing.T) {
	// A negative TTL means the token is already expired the instant it's created.
	jm := NewJWTManager("test-secret", -time.Minute, time.Hour)

	token, _, err := jm.NewAccessToken(testUser())
	require.NoError(t, err)

	_, err = jm.ParseAccessToken(token)
	assert.Error(t, err)
}

func TestJWTManager_ParseAccessToken_WrongTokenType(t *testing.T) {
	jm := NewJWTManager("test-secret", time.Minute, time.Hour)

	// Create a refresh token, then try to parse it as an access token.
	refreshToken, _, err := jm.NewRefreshToken(testUser())
	require.NoError(t, err)

	_, err = jm.ParseAccessToken(refreshToken)
	assert.Error(t, err)
}

func TestJWTManager_ParseAccessToken_WrongSecret(t *testing.T) {
	issuer := NewJWTManager("correct-secret", time.Minute, time.Hour)
	verifier := NewJWTManager("different-secret", time.Minute, time.Hour)

	token, _, err := issuer.NewAccessToken(testUser())
	require.NoError(t, err)

	_, err = verifier.ParseAccessToken(token)
	assert.Error(t, err)
}

func TestJWTManager_ParseAccessToken_EmptyToken(t *testing.T) {
	jm := NewJWTManager("test-secret", time.Minute, time.Hour)

	_, err := jm.ParseAccessToken("")
	assert.Error(t, err)
}
