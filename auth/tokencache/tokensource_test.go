// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tokencache

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/oauth2"
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
		Expiry:       time.Now().Add(minTTL + 2*time.Second),
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
