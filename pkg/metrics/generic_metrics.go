///////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package metrics

import (
	"fmt"

	gometrics "github.com/rcrowley/go-metrics"
)

// GenericMetric is generic metric interface
type GenericMetric interface {
	Update(interface{}) error
}

// Counter is int64 counter interface
type Counter struct {
	Name string
}

// Update value
func (c *Counter) Update(i interface{}) error {
	if val, ok := i.(int64); ok {
		counter := gometrics.GetOrRegisterCounter(c.Name, nil)
		counter.Inc(val)
		return nil
	}
	return fmt.Errorf("Update %s metric failed: cannot convert to int64 value", c.Name)
}

// NewCounter create new counter
func NewCounter(name string) GenericMetric {
	return &Counter{
		Name: name,
	}
}

// Gauge float64 value metrics
type Gauge struct {
	Name string
}

// Update value
func (g *Gauge) Update(i interface{}) error {
	if val, ok := i.(float64); ok {
		duration := gometrics.GetOrRegisterGaugeFloat64(g.Name, nil)
		duration.Update(val)
		return nil
	}
	return fmt.Errorf("Update %s metric failed: cannot convert to float64 value", g.Name)
}

// NewGauge creates new Gauge
func NewGauge(name string) GenericMetric {
	return &Gauge{
		Name: name,
	}
}
