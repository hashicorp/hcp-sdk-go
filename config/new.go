// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package config

import (
	"fmt"
	"io"
	"log"

	"github.com/hashicorp/hcp-sdk-go/config/geography"
)

// NewHCPConfig will return a HCPConfig. The configuration will be constructed
// from default values and the passed options.
//
// Based on the provided options the configuration values can be provided
// directly or will be read from other places (e.g. the environment).
//
// The configuration will default to values for the HCP production environment
// in US, but can be overwritten for development purposes and additional
// geographies.
//
// In addition to the default values the configuration requires client
// credentials to be provided through one of the passed options (e.g. by using
// WithCredentials or by using FromEnv and providing the client credentials via
// environment variables).
func NewHCPConfig(opts ...HCPConfigOption) (HCPConfig, error) {
	// Prepare basic config with default values for us
	config, err := configFromGeography(newDefaultConfig(), geography.Default)
	if err != nil {
		return nil, fmt.Errorf("failed to apply geography configuration: %w", err)
	}

	if config.suppressLogging {
		log.SetOutput(io.Discard)
	}

	// Apply all options
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, fmt.Errorf("failed to apply configuration option: %w", err)
		}
	}

	// Configure the token source
	if err := config.setTokenSource(); err != nil {
		return nil, err
	}

	// Validate the HCP config
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("the configuration is not valid: %w", err)
	}

	return config, nil
}
