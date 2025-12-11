// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package tokencache

import (
	"os"
	"path"
	"testing"
	"time"

	requirepkg "github.com/stretchr/testify/require"
)

func TestCachingTokenSource_Workloads(t *testing.T) {
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

	// Create the caching token sources for workload identity
	subjectA, err := NewWorkloadTokenSource(
		cacheFile,
		"workload-a",
		tokenSourceA,
		"us",
	)
	require.NoError(err)

	subjectB, err := NewWorkloadTokenSource(
		cacheFile,
		"workload-b",
		tokenSourceB,
		"eu",
	)
	require.NoError(err)

	// Fetch both tokens. They should get cached.
	token, err := subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a1", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b1", token.AccessToken)

	// Verify geography is stored in the cache file
	cachedTokens, err := readCache(cacheFile)
	require.NoError(err)
	require.NotNil(cachedTokens)
	require.Equal("us", cachedTokens.Workloads["workload-a"].Geography)
	require.Equal("eu", cachedTokens.Workloads["workload-b"].Geography)

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

// TestCachingTokenSource_Workloads_CompatibleWithNoGeography ensures that tokens
// cached without a geography are still compatible with the current implementation
// that requires a geography. This is needed to ensure backwards compatibility.
func TestCachingTokenSource_Workloads_CompatibleWithNoGeography(t *testing.T) {
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
		Workloads: map[string]cacheEntry{
			"workload-a": {
				AccessToken:       "access-token-a1",
				AccessTokenExpiry: time.Now().Add(5 * time.Minute),
				Geography:         "", // No geography set to simulate old cache file
			},
			"workload-b": {
				AccessToken:       "access-token-b1",
				AccessTokenExpiry: time.Now().Add(5 * time.Minute),
				Geography:         "", // No geography set to simulate old cache file
			},
		},
	}

	err = cachedCredentials.write(cacheFile)
	require.NoError(err)
	// Verify geography is empty in the cache file
	readCacheOld, err := readCache(cacheFile)
	require.NoError(err)
	require.Equal("", readCacheOld.Workloads["workload-a"].Geography)
	require.Equal("", readCacheOld.Workloads["workload-b"].Geography)

	// Create a test token sources
	tokenSourceA := NewTestTokenSource("a")
	tokenSourceB := NewTestTokenSource("b")

	// Create the caching token sources for workload identity
	subjectA, err := NewWorkloadTokenSource(
		cacheFile,
		"workload-a",
		tokenSourceA,
		"us",
	)
	require.NoError(err)

	subjectB, err := NewWorkloadTokenSource(
		cacheFile,
		"workload-b",
		tokenSourceB,
		"eu",
	)
	require.NoError(err)

	// Fetch both tokens. They should get refreshed and cached.
	token, err := subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a1", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b1", token.AccessToken)

	// Verify geography is stored in the cache file
	readCache, err := readCache(cacheFile)
	require.NoError(err)
	require.Equal("us", readCache.Workloads["workload-a"].Geography)
	require.Equal("eu", readCache.Workloads["workload-b"].Geography)
}

// TestCachingTokenSource_Workload_SwitchGeography ensures that
//   - switching the geography of a workload token source creates a new cache entry
//     if the geography is changed.
//   - existing cached token is not overwritten if the geography is the same.
func TestCachingTokenSource_Workload_SwitchGeography(t *testing.T) {
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

	// Create the caching token sources for workloads
	subjectA, err := NewWorkloadTokenSource(
		cacheFile,
		"workload-a",
		tokenSourceA,
		"us",
	)
	require.NoError(err)

	subjectB, err := NewWorkloadTokenSource(
		cacheFile,
		"workload-b",
		tokenSourceB,
		"us",
	)
	require.NoError(err)

	// Fetch both tokens. They should get cached.
	token, err := subjectA.Token()
	require.NoError(err)
	require.Equal("access-token-a1", token.AccessToken)

	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b1", token.AccessToken)

	// Switch tokenSourceA to a different geography - "eu"
	subjectC, err := NewWorkloadTokenSource(cacheFile, "workload-a", tokenSourceA, "eu")
	require.NoError(err)

	// Fetch the token once. It should create a new cache entry due to switching geography.
	token, err = subjectC.Token()
	require.NoError(err)
	require.Equal("access-token-a2", token.AccessToken)

	// Fetch the token of US a second time. They should be returned from cache.
	token, err = subjectB.Token()
	require.NoError(err)
	require.Equal("access-token-b1", token.AccessToken)

	// Verify geography is stored in the cache file
	newCachedCredentials, err := readCache(cacheFile)
	require.NoError(err)
	require.Equal("access-token-a2", newCachedCredentials.Workloads["workload-a"].AccessToken)
	require.Equal("eu", newCachedCredentials.Workloads["workload-a"].Geography)
	// The cached token for workload-b should not be changed because the geography is the same
	require.Equal("access-token-b1", newCachedCredentials.Workloads["workload-b"].AccessToken)
	require.Equal("us", newCachedCredentials.Workloads["workload-b"].Geography)
}
