/*
 *  *******************************************************************************
 *  * Copyright (c) 2023 Datasance Teknoloji A.S.
 *  *
 *  * This program and the accompanying materials are made available under the
 *  * terms of the Eclipse Public License v. 2.0 which is available at
 *  * http://www.eclipse.org/legal/epl-2.0
 *  *
 *  * SPDX-License-Identifier: EPL-2.0
 *  *******************************************************************************
 *
 */

package cmd

import (
	"testing"

	"github.com/cpuguy83/strongerrors"

	"go.opencensus.io/trace"
)

func TestGetTracingExporter(t *testing.T) {
	defer delete(tracingExporters, "mock")

	mockExporterFn := func(_ TracingExporterOptions) (trace.Exporter, error) {
		return nil, nil
	}

	_, err := GetTracingExporter("notexist", TracingExporterOptions{})
	if !strongerrors.IsNotFound(err) {
		t.Fatalf("expected not found error, got: %v", err)
	}

	RegisterTracingExporter("mock", mockExporterFn)

	if _, err := GetTracingExporter("mock", TracingExporterOptions{}); err != nil {
		t.Fatal(err)
	}
}

func TestAvailableExporters(t *testing.T) {
	defer delete(tracingExporters, "mock")

	mockExporterFn := func(_ TracingExporterOptions) (trace.Exporter, error) {
		return nil, nil
	}
	RegisterTracingExporter("mock", mockExporterFn)

	for _, e := range AvailableTraceExporters() {
		if e == "mock" {
			return
		}
	}

	t.Fatal("could not find mock exporter in list of registered exporters")
}
