package tokencache

import (
	"os"
	"path"
	"testing"

	requirepkg "github.com/stretchr/testify/require"
)

func TestTokenCache_ClearLoginCache(t *testing.T) {
	require := requirepkg.New(t)

	// Create a temporary directory to hold the test files
	testDirectory, err := os.MkdirTemp("", "test-caching-token-source-")
	require.NoError(err)

	// Compile the credential cache path
	cacheFile := path.Join(testDirectory, ".config/hcp/creds-cache.json")

	// Ensure the tests files are cleaned up
	defer os.RemoveAll(cacheFile)

	// Create a test token source
	tokenSource := NewTestTokenSource("")

	// Create the caching token source for logins
	subject := NewLoginTokenSource(
		cacheFile,
		false,
		tokenSource,
		nil,
	)

	// Fetch the token once. It should get cached.
	token, err := subject.Token()
	require.NoError(err)
	require.Equal("access-token-1", token.AccessToken)

	// Fetch the token a second time to verify that it was cached
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-1", token.AccessToken)

	// Clear the login cache
	err = ClearLoginCache(cacheFile)
	require.NoError(err)

	// Fetch the token a third time. A new token should be fetched as the cache was cleared.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-2", token.AccessToken)
}
