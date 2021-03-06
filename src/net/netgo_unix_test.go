// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !cgo netgo
// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package net

import (
	"context"
	"testing"
)

func TestGoLookupIP(t *testing.T) {
	host := "localhost"
	_, err, ok := cgoLookupIP(host)
	if ok {
		t.Errorf("cgoLookupIP must be a placeholder")
	}
	if err != nil {
		t.Error(err)
	}
	if _, err := goLookupIP(context.Background(), host); err != nil {
		t.Error(err)
	}
}
