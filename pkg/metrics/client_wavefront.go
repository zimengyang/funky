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
}

// Report reports metrics to metrics server
func (wc *WavefrontClient) Report() error {
	fmt.Println("wavefront client")

	hostTags := map[string]string{
		"source": "TODO",
	}
	wavefront.WavefrontDirect(
		gometrics.DefaultRegistry,
		5*time.Second,
		hostTags, "",
		wc.URL, wc.Token)

	return nil
}
