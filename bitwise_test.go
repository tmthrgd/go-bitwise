// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitwise

import (
	"bytes"
	"math/rand"
	"testing"
	"testing/quick"
)

type testVector struct {
	dst []byte
	a   []byte
	b   []byte
}

var xorTestVectors = []testVector{
	{
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
	},
}

var andTestVectors = []testVector{
	{
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
	},
	{
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
}

var andNotTestVectors = []testVector{
	{
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
	},
	{
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
}

var nandTestVectors = []testVector{
	{
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
	},
	{
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
	},
}

var orTestVectors = []testVector{
	{
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
	},
}

var norTestVectors = []testVector{
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
	},
	{
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
	},
	{
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
}

var notTestVectors = []testVector{
	{
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		nil,
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		nil,
	},
	{
		[]byte{0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A, 0x55, 0xAA, 0xA5, 0x5A},
		[]byte{0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5, 0xAA, 0x55, 0x5A, 0xA5},
		nil,
	},
}

func testXORBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}

	return n
}

func testAndBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	for i := 0; i < n; i++ {
		dst[i] = a[i] & b[i]
	}

	return n
}

func testAndNotBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	for i := 0; i < n; i++ {
		dst[i] = a[i] &^ b[i]
	}

	return n
}

func testNandBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	for i := 0; i < n; i++ {
		dst[i] = ^(a[i] & b[i])
	}

	return n
}

func testOrBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	for i := 0; i < n; i++ {
		dst[i] = a[i] | b[i]
	}

	return n
}

func testNorBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	for i := 0; i < n; i++ {
		dst[i] = ^(a[i] | b[i])
	}

	return n
}

func testNotBytes(dst, src, _ []byte) int {
	n := len(src)
	if len(dst) < n {
		n = len(dst)
	}

	for i := 0; i < n; i++ {
		dst[i] = ^src[i]
	}

	return n
}

func testNotThree(dst, src, _ []byte) int {
	return Not(dst, src)
}

func testThree(t *testing.T, fn, testFn func(dst, a, b []byte) int, testVectors []testVector) {
	for i, vector := range testVectors {
		dst := make([]byte, len(vector.dst))
		fn(dst, vector.a, vector.b)

		if !bytes.Equal(vector.dst, dst) {
			t.Errorf("test case #%d failed, expected %x, got %x", i, vector.dst, dst)
		}
	}

	for alignP := 0; alignP < 2; alignP++ {
		for alignQ := 0; alignQ < 2; alignQ++ {
			for alignD := 0; alignD < 2; alignD++ {
				p := make([]byte, 1024)[alignP:]
				rand.Read(p)

				q := make([]byte, 1024)[alignQ:]
				rand.Read(q)

				d1 := make([]byte, 1024+alignD)[alignD:]
				fn(d1, p, q)

				d2 := make([]byte, 1024+alignD)[alignD:]
				testFn(d2, p, q)

				if !bytes.Equal(d1, d2) {
					t.Error("not equal")
				}
			}
		}
	}

	if err := quick.CheckEqual(func(dst, a, b []byte) []byte {
		d1 := append([]byte{}, dst...)
		testFn(d1, a, b)
		return d1
	}, func(dst, a, b []byte) []byte {
		fn(dst, a, b)
		return dst
	}, &quick.Config{
		MaxCountScale: 500,
	}); err != nil {
		t.Error(err)
	}
}

func TestXOR(t *testing.T) {
	testThree(t, XOR, testXORBytes, xorTestVectors)
}

func TestAnd(t *testing.T) {
	testThree(t, And, testAndBytes, andTestVectors)
}

func TestAndNot(t *testing.T) {
	testThree(t, AndNot, testAndNotBytes, andNotTestVectors)
}

func TestNand(t *testing.T) {
	testThree(t, Nand, testNandBytes, nandTestVectors)
}

func TestOr(t *testing.T) {
	testThree(t, Or, testOrBytes, orTestVectors)
}

func TestNor(t *testing.T) {
	testThree(t, Nor, testNorBytes, norTestVectors)
}

func TestNot(t *testing.T) {
	testThree(t, testNotThree, testNotBytes, notTestVectors)
}

var benchSizes = []struct {
	name string
	l    int
}{
	{"15", 15},
	{"32", 32},
	{"128", 128},
	{"1K", 1 * 1024},
	{"16K", 16 * 1024},
	{"128K", 128 * 1024},
	{"1M", 1024 * 1024},
	{"16M", 16 * 1024 * 1024},
	{"128M", 128 * 1024 * 1024},
}

func benchmarkThree(b *testing.B, testFn func(dst, a, b []byte) int) {
	maxSize := benchSizes[len(benchSizes)-1]

	dst := make([]byte, maxSize.l)

	p, q := make([]byte, maxSize.l), make([]byte, maxSize.l)
	rand.Read(p)
	rand.Read(q)

	for _, size := range benchSizes {
		b.Run(size.name, func(b *testing.B) {
			b.SetBytes(int64(size.l))

			dst, p, q := dst[:size.l], p[:size.l], q[:size.l]

			for i := 0; i < b.N; i++ {
				testFn(dst, p, q)
			}
		})
	}
}

func BenchmarkXOR(b *testing.B) {
	benchmarkThree(b, XOR)
}

func BenchmarkXORGo(b *testing.B) {
	benchmarkThree(b, testXORBytes)
}

func BenchmarkAnd(b *testing.B) {
	benchmarkThree(b, And)
}

func BenchmarkAndGo(b *testing.B) {
	benchmarkThree(b, testAndBytes)
}

func BenchmarkAndNot(b *testing.B) {
	benchmarkThree(b, AndNot)
}

func BenchmarkAndNotGo(b *testing.B) {
	benchmarkThree(b, testAndNotBytes)
}

func BenchmarkNand(b *testing.B) {
	benchmarkThree(b, Nand)
}

func BenchmarkNandGo(b *testing.B) {
	benchmarkThree(b, testNandBytes)
}

func BenchmarkOr(b *testing.B) {
	benchmarkThree(b, Or)
}

func BenchmarkOrGo(b *testing.B) {
	benchmarkThree(b, testOrBytes)
}

func BenchmarkNor(b *testing.B) {
	benchmarkThree(b, Nor)
}

func BenchmarkNorGo(b *testing.B) {
	benchmarkThree(b, testNorBytes)
}

func BenchmarkNot(b *testing.B) {
	benchmarkThree(b, testNotThree)
}

func BenchmarkNotGo(b *testing.B) {
	benchmarkThree(b, testNotBytes)
}
