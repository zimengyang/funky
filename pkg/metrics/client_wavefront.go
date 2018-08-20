///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package metrics

import (
	"fmt"
	"time"

	gometrics "github.com/rcrowley/go-metrics"
	wavefront "github.com/wavefrontHQ/go-metrics-wavefront"
)

// WavefrontClient client implementation
type WavefrontClient struct {
	URL   string
	Token string

	Metrics map[string]GenericMetric
}

// Report reports metrics to metrics server
func (wc *WavefrontClient) Report() error {
	fmt.Println("wavefront client")

	hostTags := map[string]string{
		// "source": "TODO",
	}

	return wavefront.WavefrontDirect(
		gometrics.DefaultRegistry,
		5*time.Second,
		hostTags, "",
		wc.URL, wc.Token)
}

// Register register metric to client
func (wc *WavefrontClient) Register(metrics map[string]GenericMetric) {
	for name, m := range metrics {
		if _, ok := wc.Metrics[name]; !ok {
			wc.Metrics[name] = m
		}
		fmt.Printf("[INFO] metrics %s already registered\n", name)
	}
}

// Update updates metric
func (wc *WavefrontClient) Update(name string, i interface{}) error {

	if m, ok := wc.Metrics[name]; ok {
		return m.Update(i)
	}
	return fmt.Errorf("metrics %s not registered by client", name)
}
