// Copyright 2014 Stas Kalashnikov. All rights reserved.

package euclid

import "testing"

// A function signature for implementation of Euclid's algorithm
// to find the greatest common divisor of two positive integers.
type gcd_func func(m, n int) int

// Functions with slightly different implementations to find
// the greatest commond divisor using Euclid's algorithm.
var gcd_funcs = [...]gcd_func{
	gcd_e,
	gcd_goto,
	gcd_recursion,
	gcd_f,
	gcd_fs,
}

var gcd_func_names = [...]string{
	"gcd_e",
	"gcd_goto",
	"gcd_recursion",
	"gcd_f",
	"gcd_fs",
}

func TestAllImplementations(t *testing.T) {
	for i := range gcd_funcs {
		if r, err := gcd(2166, 6099, gcd_funcs[i]); err != nil {
			t.Error(gcd_func_names[i], "returned error")
		} else if r != 57 {
			t.Error(gcd_func_names[i], "returned", r, "but expected 57")
		}
	}
}

func TestGCD(t *testing.T) {
	r, err := GCD(2166, 6099)
	if err != nil {
		t.Error("error returned")
	} else if r != 57 {
		t.Error("expected 57, actual", r)
	}
}

func TestGCDInput(t *testing.T) {
	if _, err := GCD(-1, 2); err == nil {
		t.Error("there's no error when m < 0")
	}
	if _, err := GCD(2, -1); err == nil {
		t.Error("there's no error when n < 0")
	}
	if _, err := GCD(0, 2); err == nil {
		t.Error("there's no error when m = 0")
	}
	if _, err := GCD(2, 0); err == nil {
		t.Error("there's no error when n = 0")
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(2166, 6099)
	}
}

func BenchmarkE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcd_e(2166, 6099)
	}
}

func BenchmarkGoto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcd_goto(2166, 6099)
	}
}

func BenchmarkRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcd_recursion(2166, 6099)
	}
}

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcd_f(2166, 6099)
	}
}

func BenchmarkFs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcd_fs(2166, 6099)
	}
}
