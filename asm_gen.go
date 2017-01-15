// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

// +build ignore

package main

import "github.com/tmthrgd/asm"

const header = `// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// This file is auto-generated - do not modify

// +build amd64,!gccgo,!appengine
`

func threeArgumentASM(a *asm.Asm, name string, pop, opb func(ops ...asm.Operand)) {
	a.NewFunction(name)
	a.NoSplit()

	dst := a.Argument("dst", 8)
	srcA := a.Argument("a", 8)
	srcB := a.Argument("b", 8)
	length := a.Argument("len", 8)

	a.Start()

	hugeloop := a.NewLabel("hugeloop")
	bigloop := a.NewLabel("bigloop")
	loop := a.NewLabel("loop")
	ret := a.NewLabel("ret")

	di, sA, sB, cx := asm.DI, asm.SI, asm.DX, asm.BX

	a.Movq(di, dst)
	a.Movq(sA, srcA)
	a.Movq(sB, srcB)
	a.Movq(cx, length)

	a.Cmpq(asm.Constant(16), cx)
	a.Jb(loop)

	a.Cmpq(asm.Constant(64), cx)
	a.Jb(bigloop)

	a.Label(hugeloop)

	a.Movou(asm.X0, asm.Address(sA, cx, asm.SX1, -16))
	a.Movou(asm.X2, asm.Address(sA, cx, asm.SX1, -32))
	a.Movou(asm.X4, asm.Address(sA, cx, asm.SX1, -48))
	a.Movou(asm.X6, asm.Address(sA, cx, asm.SX1, -64))

	a.Movou(asm.X1, asm.Address(sB, cx, asm.SX1, -16))
	a.Movou(asm.X3, asm.Address(sB, cx, asm.SX1, -32))
	a.Movou(asm.X5, asm.Address(sB, cx, asm.SX1, -48))
	a.Movou(asm.X7, asm.Address(sB, cx, asm.SX1, -64))

	pop(asm.X1, asm.X0)
	pop(asm.X3, asm.X2)
	pop(asm.X5, asm.X4)
	pop(asm.X7, asm.X6)

	a.Movou(asm.Address(di, cx, asm.SX1, -16), asm.X1)
	a.Movou(asm.Address(di, cx, asm.SX1, -32), asm.X3)
	a.Movou(asm.Address(di, cx, asm.SX1, -48), asm.X5)
	a.Movou(asm.Address(di, cx, asm.SX1, -64), asm.X7)

	a.Subq(cx, asm.Constant(64))
	a.Jz(ret)

	a.Cmpq(asm.Constant(64), cx)
	a.Jae(hugeloop)

	a.Cmpq(asm.Constant(16), cx)
	a.Jb(loop)

	a.Label(bigloop)

	a.Movou(asm.X0, asm.Address(sA, cx, asm.SX1, -16))
	a.Movou(asm.X1, asm.Address(sB, cx, asm.SX1, -16))

	pop(asm.X1, asm.X0)

	a.Movou(asm.Address(di, cx, asm.SX1, -16), asm.X1)

	a.Subq(cx, asm.Constant(16))
	a.Jz(ret)

	a.Cmpq(asm.Constant(16), cx)
	a.Jae(bigloop)

	a.Label(loop)

	a.Movb(asm.AX, asm.Address(sA, cx, asm.SX1, -1))
	opb(asm.AX, asm.Address(sB, cx, asm.SX1, -1))
	a.Movb(asm.Address(di, cx, asm.SX1, -1), asm.AX)

	a.Subq(cx, asm.Constant(1))
	a.Jnz(loop)

	a.Label(ret)

	a.Ret()
}

func xorASM(a *asm.Asm) {
	threeArgumentASM(a, "xorASM", a.Pxor, a.Xorb)
}

func andASM(a *asm.Asm) {
	threeArgumentASM(a, "andASM", a.Pand, a.Andb)
}

func andNotASM(a *asm.Asm) {
	threeArgumentASM(a, "andNotASM", a.Pandn, func(ops ...asm.Operand) {
		if len(ops) != 2 {
			panic("wrong number of operands")
		}

		a.Movb(asm.R15, ops[1])
		a.Notb(asm.R15)
		a.Andb(ops[0], asm.R15)
	})
}

func orASM(a *asm.Asm) {
	threeArgumentASM(a, "orASM", a.Por, a.Orb)
}

func main() {
	if err := asm.Do("bitwise_xor_amd64.s", header, xorASM); err != nil {
		panic(err)
	}

	if err := asm.Do("bitwise_and_amd64.s", header, andASM); err != nil {
		panic(err)
	}

	if err := asm.Do("bitwise_andnot_amd64.s", header, andNotASM); err != nil {
		panic(err)
	}

	if err := asm.Do("bitwise_or_amd64.s", header, orASM); err != nil {
		panic(err)
	}
}