///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package metrics

// Client interface for pluggable monitoring tool to plug in funky
type Client interface {
	Report() error

	// Convert gometrics including generic metrics to Client spefict metrics
	// Format() error
}
