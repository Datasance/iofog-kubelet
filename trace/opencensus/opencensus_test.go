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

package opencensus

import (
	"testing"

	"github.com/eclipse-iofog/iofog-kubelet/v2/trace"
)

func TestTracerImplementsTracer(t *testing.T) {
	// ensure that Adapter implements trace.Tracer
	if tt := trace.Tracer(Adapter{}); tt == nil {
	}
}
