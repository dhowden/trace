// Copyright 2015, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package trace provides convenience functions for interfacting with Trace instances within
// context.Contexts.
package trace

import (
	"fmt"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"
)

// LazyLog is a light wrapper around golang.org/x/net/trace/Trace.LazyLog which extracts the
// Trace instance from ctx and calls LazyLog on it.  If no such instance exists, this method
// returns.
func LazyLog(ctx context.Context, x fmt.Stringer, sensitive bool) {
	tr, ok := trace.FromContext(ctx)
	if !ok {
		return
	}
	tr.LazyLog(x, sensitive)
}

// LazyPrintf is a light wrapper around golang.org/x/net/trace/Trace.LazyPrintf which extracts
// a Trace instance from ctx and calls LazyPrintf on it.  If no such instance exists, this method
// returns.
func LazyPrintf(ctx context.Context, format string, a ...interface{}) {
	tr, ok := trace.FromContext(ctx)
	if !ok {
		return
	}
	tr.LazyPrintf(format, a...)
}
