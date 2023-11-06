package tokencache

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	requirepkg "github.com/stretchr/testify/require"
)

func TestCachingTokenSource_Login_WithoutOauthConfig(t *testing.T) {
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

	// Fetch the token a second time. It should be returned from cache.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-1", token.AccessToken)

	// Wait for token to expire
	time.Sleep(2 * time.Second)

	// Fetch the token a third time. It should be read from the token source and get cached.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-2", token.AccessToken)

	// Fetch the token a fourth time. It should be returned from cache.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-2", token.AccessToken)

	// Create a new token source with the same file and expect it to return the cached value
	subject = NewLoginTokenSource(
		cacheFile,
		false,
		tokenSource,
		nil,
	)

	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-2", token.AccessToken)

	// Create a new token source with force login enabled and expect it to return a new token
	subject = NewLoginTokenSource(
		cacheFile,
		true,
		tokenSource,
		nil,
	)

	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-3", token.AccessToken)

	// Fetch the token another time to verify that force login is only executed once.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-3", token.AccessToken)
}

func TestCachingTokenSource_Login_WithOauthConfig(t *testing.T) {
	require := requirepkg.New(t)

	// Create a temporary directory to hold the test files
	testDirectory, err := os.MkdirTemp("", "test-caching-token-source-")
	require.NoError(err)

	// Compile the credential cache path
	cacheFile := path.Join(testDirectory, ".config/hcp/creds-cache.json")

	// Ensure the tests files are cleaned up
	defer os.RemoveAll(cacheFile)

	// Create a test token source
	tokenSource := NewTestTokenSource("a")

	// Create a test config
	configTokenSource := NewTestTokenSource("b")
	config := NewTestOauth2Config(configTokenSource)

	// Create the caching token source for interactive logins
	subject := NewLoginTokenSource(
		cacheFile,
		false,
		tokenSource,
		config,
	)

	// Fetch the token once. It should get cached.
	token, err := subject.Token()
	require.NoError(err)
	require.Equal("access-token-a1", token.AccessToken)

	// Fetch the token a second time. It should be returned from cache.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-a1", token.AccessToken)

	// Wait for token to expire
	time.Sleep(2 * time.Second)

	// Fetch the token a third time. It should be refreshed using the configured oauth config and get cached.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-b1", token.AccessToken)

	// Fetch the token a fourth time. It should be returned from cache.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-b1", token.AccessToken)

	// Wait for token to expire again
	time.Sleep(2 * time.Second)

	// Expect refresh to fail
	configTokenSource.FailNextWith(fmt.Errorf("no refresh token available"))

	// Fetch the token again and expect it to be fetched from the configured token source.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-a2", token.AccessToken)
}
