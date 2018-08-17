///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package metrics

// Client interface for pluggable monitoring tool to plug in funky
type Client interface {
	Register(metrics map[string]GenericMetric)
	Update(name string, i interface{}) error
	Report() error
}
