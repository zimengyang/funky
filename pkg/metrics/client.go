///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package metrics

// Client interface for pluggable monitoring tool to plug in funky
// TODO: Formats the gometrics to Client specific metrics
type Client interface {
	Report() error
	// Format() error
}
