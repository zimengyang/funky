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

// GenericMetricInvocations function invocation counter
type GenericMetricInvocations struct {
	Name string
}

// GenericMetricExecutionDuration function execution duration with unit
type GenericMetricExecutionDuration struct {
	Name string
	Unit string
}

// GenericMetricTimeout time out error counter
type GenericMetricTimeout struct {
	Name string
}

// GenericMetricConnectionRefused connection refused counter
type GenericMetricConnectionRefused struct {
	Name string
}

// GenericMetricUnknownSystemError unknown system error counter
type GenericMetricUnknownSystemError struct {
	Name string
}

var (
	// InvocationsCounter var
	InvocationsCounter = GenericMetricInvocations{
		Name: "dispatch.function.invocations",
	}
	// ExecutionDuration var
	ExecutionDuration = GenericMetricExecutionDuration{
		Name: "dispatch.function.executionduration",
		Unit: "ms",
	}
	// TimeoutCounter var
	TimeoutCounter = GenericMetricTimeout{
		Name: "dispatch.function.timeout",
	}
	// ConnectionRefusedCounter var
	ConnectionRefusedCounter = GenericMetricConnectionRefused{
		Name: "dispatch.function.connectionrefused",
	}
	// UnknownSystemErrorCounter var
	UnknownSystemErrorCounter = GenericMetricUnknownSystemError{
		Name: "dispatch.function.unknownsystemerror",
	}
)

// Update will update the metrics w/o input (i as interface)
func (m *GenericMetricInvocations) Update(i interface{}) error {
	if val, ok := i.(int64); ok {
		counter := gometrics.GetOrRegisterCounter(m.Name, nil)
		counter.Inc(val)
		return nil
	}
	return fmt.Errorf("Update %s metric failed: cannot convert to float64 value", m.Name)
}

// Update will update the metrics w/o input (i as interface)
func (m *GenericMetricExecutionDuration) Update(i interface{}) error {
	if val, ok := i.(float64); ok {
		duration := gometrics.GetOrRegisterGaugeFloat64(m.Name, nil)
		duration.Update(val)
		return nil
	}
	return fmt.Errorf("Update %s metric failed: cannot convert to float64 value", m.Name)
}

// Update will update the metrics w/o input (i as interface)
func (m *GenericMetricTimeout) Update(i interface{}) error {
	if val, ok := i.(int64); ok {
		counter := gometrics.GetOrRegisterCounter(m.Name, nil)
		counter.Inc(val)
		return nil
	}
	return fmt.Errorf("Update %s metric failed: cannot convert to float64 value", m.Name)
}

// Update will update the metrics w/o input (i as interface)
func (m *GenericMetricConnectionRefused) Update(i interface{}) error {
	if val, ok := i.(int64); ok {
		counter := gometrics.GetOrRegisterCounter(m.Name, nil)
		counter.Inc(val)
		return nil
	}
	return fmt.Errorf("Update %s metric failed: cannot convert to float64 value", m.Name)
}

// Update will update the metrics w/o input (i as interface)
func (m *GenericMetricUnknownSystemError) Update(i interface{}) error {
	if val, ok := i.(int64); ok {
		counter := gometrics.GetOrRegisterCounter(m.Name, nil)
		counter.Inc(val)
		return nil
	}
	return fmt.Errorf("Update %s metric failed: cannot convert to float64 value", m.Name)
}
