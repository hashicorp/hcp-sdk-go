// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tokencache

import (
	"os"
	"path"
	"testing"
	"time"

	requirepkg "github.com/stretchr/testify/require"
)

func TestCachingTokenSource_ServicePrincipals(t *testing.T) {
	require := requirepkg.New(t)

	// Create a temporary directory to hold the test files
	testDirectory, err := os.MkdirTemp("", "test-caching-token-source-")
	require.NoError(err)

	// Compile the credential cache path
	cacheFile := path.Join(testDirectory, ".config/hcp/creds-cache.json")

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
