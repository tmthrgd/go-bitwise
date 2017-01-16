// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

// +build amd64,!gccgo,!appengine

package bitwise

func XOR(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	if n == 0 {
		return 0
	}

	xorASM(&dst[0], &a[0], &b[0], uint64(n))
	return n
}

func And(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	if n == 0 {
		return 0
	}

	andASM(&dst[0], &a[0], &b[0], uint64(n))
	return n
}

func AndNot(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	if n == 0 {
		return 0
	}

	andNotASM(&dst[0], &a[0], &b[0], uint64(n))
	return n
}

func Or(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	if n == 0 {
		return 0
	}

	orASM(&dst[0], &a[0], &b[0], uint64(n))
	return n
}

func Not(dst, src []byte) int {
	n := len(src)
	if len(dst) < n {
		n = len(dst)
	}

	if n == 0 {
		return 0
	}

	notASM(&dst[0], &src[0], uint64(n))
	return n
}

//go:generate go run asm_gen.go

// This function is implemented in bitwise_xor_amd64.s
//go:noescape
func xorASM(dst, a, b *byte, len uint64)

// This function is implemented in bitwise_and_amd64.s
//go:noescape
func andASM(dst, a, b *byte, len uint64)

// This function is implemented in bitwise_andnot_amd64.s
//go:noescape
func andNotASM(dst, a, b *byte, len uint64)

// This function is implemented in bitwise_or_amd64.s
//go:noescape
func orASM(dst, a, b *byte, len uint64)

// This function is implemented in bitwise_not_amd64.s
//go:noescape
func notASM(dst, src *byte, len uint64)
