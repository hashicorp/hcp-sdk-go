package tokencache

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCache_ExpiredTokenRemoval(t *testing.T) {
	subject := cache{}

	// Prepare cache entries
	expired := cacheEntry{
		AccessToken:       "access-token",
		AccessTokenExpiry: time.Now().Add(-1 * time.Minute),
	}

	refreshable := cacheEntry{
		AccessToken:       "access-token",
		RefreshToken:      "refresh-token",
		AccessTokenExpiry: time.Now().Add(-1 * time.Minute),
	}

	valid := cacheEntry{
		AccessToken:       "access-token",
		AccessTokenExpiry: time.Now().Add(1 * time.Minute),
	}

	// Set cache entries for service principals
	subject.ServicePrincipals = map[string]cacheEntry{
		"expired":     expired,
		"refreshable": refreshable,
		"valid":       valid,
	}

	// Set cache entries for workloads
	subject.Workloads = map[string]cacheEntry{
		"expired":     expired,
		"refreshable": refreshable,
		"valid":       valid,
	}

	// Exercise
	subject.removeExpiredTokens()

	// Evaluate
	expectedEntries := map[string]cacheEntry{
		"refreshable": refreshable,
		"valid":       valid,
	}

	require.Equal(t, expectedEntries, subject.ServicePrincipals)
	require.Equal(t, expectedEntries, subject.Workloads)
}
