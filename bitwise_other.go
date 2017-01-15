// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !amd64 gccgo appengine

package bitwise

import (
	"runtime"
	"unsafe"
)

const wordSize = int(unsafe.Sizeof(uintptr(0)))
const supportsUnaligned = runtime.GOARCH == "386" || runtime.GOARCH == "amd64"

func fastXORBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	w := n / wordSize
	if w > 0 {
		dw := *(*[]uintptr)(unsafe.Pointer(&dst))
		aw := *(*[]uintptr)(unsafe.Pointer(&a))
		bw := *(*[]uintptr)(unsafe.Pointer(&b))

		for i := 0; i < w; i++ {
			dw[i] = aw[i] ^ bw[i]
		}
	}

	for i := (n - n%wordSize); i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}

	return n
}

func safeXORBytes(dst, a, b []byte) int {
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

func XOR(dst, a, b []byte) int {
	if supportsUnaligned {
		return fastXORBytes(dst, a, b)
	}

	// TODO: if (dst, a, b) have common alignment
	// we could still try fastXORBytes.
	return safeXORBytes(dst, a, b)
}

func fastAndBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	w := n / wordSize
	if w > 0 {
		dw := *(*[]uintptr)(unsafe.Pointer(&dst))
		aw := *(*[]uintptr)(unsafe.Pointer(&a))
		bw := *(*[]uintptr)(unsafe.Pointer(&b))

		for i := 0; i < w; i++ {
			dw[i] = aw[i] & bw[i]
		}
	}

	for i := (n - n%wordSize); i < n; i++ {
		dst[i] = a[i] & b[i]
	}

	return n
}

func safeAndBytes(dst, a, b []byte) int {
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

func And(dst, a, b []byte) int {
	if supportsUnaligned {
		return fastAndBytes(dst, a, b)
	}

	// TODO: if (dst, a, b) have common alignment
	// we could still try fastAndBytes.
	return safeAndBytes(dst, a, b)
}

func fastAndNotBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	w := n / wordSize
	if w > 0 {
		dw := *(*[]uintptr)(unsafe.Pointer(&dst))
		aw := *(*[]uintptr)(unsafe.Pointer(&a))
		bw := *(*[]uintptr)(unsafe.Pointer(&b))

		for i := 0; i < w; i++ {
			dw[i] = aw[i] &^ bw[i]
		}
	}

	for i := (n - n%wordSize); i < n; i++ {
		dst[i] = a[i] &^ b[i]
	}

	return n
}

func safeAndNotBytes(dst, a, b []byte) int {
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

func AndNot(dst, a, b []byte) int {
	if supportsUnaligned {
		return fastAndNotBytes(dst, a, b)
	}

	// TODO: if (dst, a, b) have common alignment
	// we could still try fastAndNotBytes.
	return safeAndNotBytes(dst, a, b)
}

func fastOrBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if len(dst) < n {
		n = len(dst)
	}

	w := n / wordSize
	if w > 0 {
		dw := *(*[]uintptr)(unsafe.Pointer(&dst))
		aw := *(*[]uintptr)(unsafe.Pointer(&a))
		bw := *(*[]uintptr)(unsafe.Pointer(&b))

		for i := 0; i < w; i++ {
			dw[i] = aw[i] | bw[i]
		}
	}

	for i := (n - n%wordSize); i < n; i++ {
		dst[i] = a[i] | b[i]
	}

	return n
}

func safeOrBytes(dst, a, b []byte) int {
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

func Or(dst, a, b []byte) int {
	if supportsUnaligned {
		return fastOrBytes(dst, a, b)
	}

	// TODO: if (dst, a, b) have common alignment
	// we could still try fastORBytes.
	return safeOrBytes(dst, a, b)
}