package tokencache

import (
	"context"
	"fmt"
	requirepkg "github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"os"
	"path"
	"testing"
	"time"
)

type testTokenSource struct {
	label         string
	returnedError error
	tokenCounter  int
	expired       bool
}

func NewTestTokenSource(label string) *testTokenSource {
	return &testTokenSource{
		label: label,
	}
}

func (source *testTokenSource) Token() (*oauth2.Token, error) {
	source.tokenCounter++

	err := source.returnedError

	// Reset expiry and error for next request
	source.expired = false
	source.returnedError = nil

	return &oauth2.Token{
		AccessToken:  fmt.Sprintf("access-token-%s%d", source.label, source.tokenCounter),
		RefreshToken: fmt.Sprintf("refresh-token-%s%d", source.label, source.tokenCounter),
		Expiry:       time.Now().Add(2 * time.Second),
	}, err
}

func (source *testTokenSource) FailNextWith(err error) {
	source.returnedError = err
}

func (source *testTokenSource) Expire() {
	source.expired = true
}

type testOauth2Config struct {
	tokenSource oauth2.TokenSource
}

func NewTestOauth2Config(tokenSource oauth2.TokenSource) *testOauth2Config {
	return &testOauth2Config{
		tokenSource: tokenSource,
	}
}

func (config *testOauth2Config) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	return config.tokenSource
}

func TestCachingTokenSource_Login_WithoutOauthConfig(t *testing.T) {
	require := requirepkg.New(t)

	// Create a temporary directory to hold the test files
	testDirectory, err := os.MkdirTemp("", "test-caching-token-source-")
	require.NoError(err)

	// Compile the credential cache path
	cacheFile := path.Join(testDirectory, ".config/hcp/credentials.json")

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
	cacheFile := path.Join(testDirectory, ".config/hcp/credentials.json")

	// Ensure the tests files are cleaned up
	defer os.RemoveAll(cacheFile)

	// Create a test token source
	tokenSource := NewTestTokenSource("a")

	// Create a test config
	configTokenSource := NewTestTokenSource("b")
	configTokenSource.FailNextWith(fmt.Errorf("no refresh token available"))
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

	// Fetch the token a third time. It should be read from the configured token source and get cached.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-b2", token.AccessToken)

	// Fetch the token a fourth time. It should be returned from cache.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-b2", token.AccessToken)
}

func TestCachingTokenSource_ServicePrincipals(t *testing.T) {
	require := requirepkg.New(t)

	// Create a temporary directory to hold the test files
	testDirectory, err := os.MkdirTemp("", "test-caching-token-source-")
	require.NoError(err)

	// Compile the credential cache path
	cacheFile := path.Join(testDirectory, ".config/hcp/credentials.json")

	// Ensure the tests files are cleaned up
	defer os.RemoveAll(cacheFile)

	// Create a test token sources
	tokenSourceA := NewTestTokenSource("a")
	tokenSourceB := NewTestTokenSource("b")

	// Create the caching token sources for service principals
	subjectA := NewServicePrincipalTokenSource(
		cacheFile,
		"client-id-a",
		tokenSourceA,
	)

	subjectB := NewServicePrincipalTokenSource(
		cacheFile,
		"client-id-b",
		tokenSourceB,
	)

	// Fetch both tokens. They should get cached.
	token, err := subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a1", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b1", token.AccessToken)

	// Fetch the tokens a second time. They should be returned from cache.
	token, err = subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a1", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b1", token.AccessToken)

	// Wait for tokens to expire
	time.Sleep(2 * time.Second)

	// Fetch the tokens a third time. They should be read from the token source and get cached.
	token, err = subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a2", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b2", token.AccessToken)

	// Fetch the token a fourth time. They should be returned from cache.
	token, err = subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a2", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b2", token.AccessToken)
}

func TestCachingTokenSource_Workloads(t *testing.T) {
	require := requirepkg.New(t)

	// Create a temporary directory to hold the test files
	testDirectory, err := os.MkdirTemp("", "test-caching-token-source-")
	require.NoError(err)

	// Compile the credential cache path
	cacheFile := path.Join(testDirectory, ".config/hcp/credentials.json")

	// Ensure the tests files are cleaned up
	defer os.RemoveAll(cacheFile)

	// Create a test token sources
	tokenSourceA := NewTestTokenSource("a")
	tokenSourceB := NewTestTokenSource("b")

	// Create the caching token sources for service principals
	subjectA := NewWorkloadTokenSource(
		cacheFile,
		"workload-a",
		tokenSourceA,
	)

	subjectB := NewWorkloadTokenSource(
		cacheFile,
		"workload-b",
		tokenSourceB,
	)

	// Fetch both tokens. They should get cached.
	token, err := subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a1", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b1", token.AccessToken)

	// Fetch the tokens a second time. They should be returned from cache.
	token, err = subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a1", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b1", token.AccessToken)

	// Wait for tokens to expire
	time.Sleep(2 * time.Second)

	// Fetch the tokens a third time. They should be read from the token source and get cached.
	token, err = subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a2", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b2", token.AccessToken)

	// Fetch the token a fourth time. They should be returned from cache.
	token, err = subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a2", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b2", token.AccessToken)
}
