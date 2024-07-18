// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc struct-markdown

package ipsw

import (
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

// StepRunConfig contains the configuration for how to start a VM
type StepRunConfig struct {
	// Set to start macos vm in recovery mode
	BootToRecovery bool `mapstructure:"boot_to_recovery"`
}

// Prepare sets the default value of "Prlctl" property.
func (c *StepRunConfig) Prepare(ctx *interpolate.Context) []error {
	if !c.BootToRecovery {
		c.BootToRecovery = false
	}

	return nil
}
