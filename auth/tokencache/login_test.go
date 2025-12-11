// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

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
	subject, _ := NewLoginTokenSource(cacheFile, tokenSource, nil, "us")

	// Fetch the token once. It should get cached.
	token, err := subject.Token()
	require.NoError(err)
	require.Equal("access-token-1", token.AccessToken)

	// Verify geography is stored in the cache file
	readCache, err := readCache(cacheFile)
	require.NoError(err)
	require.Equal("us", readCache.Login.Geography)

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
	subject, _ = NewLoginTokenSource(cacheFile, tokenSource, nil, "us")

	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-2", token.AccessToken)
}

// TestCachingTokenSource_Login_CompatibleWithNoGeography ensures that tokens
// cached without a geography are still compatible with the current implementation
// that requires a geography. This is needed to ensure backwards compatibility.
func TestCachingTokenSource_Login_CompatibleWithNoGeography(t *testing.T) {
	require := requirepkg.New(t)

	// Create a temporary directory to hold the test files
	testDirectory, err := os.MkdirTemp("", "test-caching-token-source-")
	require.NoError(err)

	// Compile the credential cache path
	cacheFile := path.Join(testDirectory, ".config/hcp/creds-cache.json")

	// Ensure the tests files are cleaned up
	defer os.RemoveAll(cacheFile)

	// Create a cached credentials without geography
	cachedCredentials := &cache{
		Login: &cacheEntry{
			AccessToken:       "access-token-1",
			RefreshToken:      "refresh-token-1",
			AccessTokenExpiry: time.Now().Add(5 * time.Minute),
			Geography:         "", // No geography set to simulate old cache file
		},
	}
	err = cachedCredentials.write(cacheFile)
	require.NoError(err)

	// Create a test token source
	tokenSource := NewTestTokenSource("")

	// Create the caching token source for logins
	subject, err := NewLoginTokenSource(cacheFile, tokenSource, nil, "us")
	require.NoError(err)

	// Fetch the token once. It should get cached.
	token, err := subject.Token()
	require.NoError(err)
	require.Equal("access-token-1", token.AccessToken)

	// Verify geography is stored in the cache file
	cachedCredentials, err = readCache(cacheFile)
	require.NoError(err)
	require.Equal("access-token-1", cachedCredentials.Login.AccessToken)
	require.Equal("us", cachedCredentials.Login.Geography)
}

func TestCachingTokenSource_Login_SwitchGeography(t *testing.T) {
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

	// Create the caching token source for logins in the "us" geography
	subject, err := NewLoginTokenSource(cacheFile, tokenSource, nil, "us")
	require.NoError(err)

	// Fetch the token once. It should get cached.
	token, err := subject.Token()
	require.NoError(err)
	require.Equal("access-token-1", token.AccessToken)

	usCachedCredentials, err := readCache(cacheFile)
	require.NoError(err)
	require.Equal("access-token-1", usCachedCredentials.Login.AccessToken)
	require.Equal("us", usCachedCredentials.Login.Geography)

	// Switch to a different geography - "eu"
	subject, err = NewLoginTokenSource(cacheFile, tokenSource, nil, "eu")
	require.NoError(err)

	// Fetch the token once. It should create a new cache entry.
	token, err = subject.Token()
	require.NoError(err)
	require.Equal("access-token-2", token.AccessToken)

	// Verify geography is stored in the cache file
	newCachedCredentials, err := readCache(cacheFile)
	require.NoError(err)
	require.Equal("access-token-2", newCachedCredentials.Login.AccessToken)
	require.Equal("eu", newCachedCredentials.Login.Geography)
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
	subject, _ := NewLoginTokenSource(cacheFile, tokenSource, config, "us")

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
