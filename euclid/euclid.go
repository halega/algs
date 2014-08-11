// Copyright 2014 Stas Kalashnikov. All rights reserved.

/*
Package euclid implements Euclid's algorithm.

	Algorithm E (Euclid's algorithm). Given two positive integers m and n, find their
	greatest common divisor, that is, the largest positive integer that evenly divides both m and n.
	E0. [Ensure m ≥ n.] If m < n, exchange m ↔ n.
	E1. [Find remainder.] Divide m by n and let r be the remainder. (We will have 0 ≤ r < n.)
	E2. [Is it zero?] If r = 0, the algorithm terminates; n is the answer.
	E3. [Reduce.] Set m ← n, n ← r, and go back to step E1. ∎

	D. Knuth, The Art of Computer Programming, Volume 1.

Default implementation uses the following modified Euclid's algorithm:

	Algorithm F (Euclid's algorithm). Given two positive integers m and n, find their greatest common divisor.
	F0. [Ensure m ≥ n.] If m < n, exchange m ↔ n.
	F1. [Remainder m/n.] Divide m by n and let m be the remainder.
	F2. [Is it zero?] If m = 0, the algorithm terminates with answer n.
	F3. [Remainder n/m.] Divide n by m and let n be the remainder.
	F4. [Is it zero?] If n = 0, the algorithm terminates with answer m; otherwise go back to step F1. ∎

	D. Knuth, The Art of Computer Programming, Volume 1, Exercise 1.1.3.
*/
package euclid

import "errors"

// A function signature for implementation of Euclid's algorithm
// to find the greatest common divisor of two positive integers.
type gcd_func func(m, n int) int

func GCD(m, n int) (int, error) {
	return gcd(m, n, gcd_fs)
}

// Main entry function with input checks
func gcd(m, n int, f gcd_func) (int, error) {
	if m <= 0 || n <= 0 {
		return 0, errors.New("gcd: integers are not positive")
	}
	// E0. [Ensure m ≥ n.]
	if m < n {
		m, n = n, m
	}
	return f(m, n), nil
}

// Classic implementation.
func gcd_e(m, n int) int {
	for r := m % n; r != 0; r = m % n {
		m, n = n, r
	}
	return n
}

// Implementation with goto statements.
func gcd_goto(m, n int) int {
E1:
	r := m % n

	// E2:
	if r == 0 {
		goto TERMINATE
	}

	// E3:
	m = n
	n = r
	goto E1

TERMINATE:
	return n
}

// Recursive implementation.
func gcd_recursion(m, n int) int {
	if r := m % n; r == 0 {
		return n
	} else {
		return gcd_recursion(n, r)
	}
}

// Implementation avoiding trivial replacements.
func gcd_f(m, n int) int {
	for {
		m = m % n
		if m == 0 {
			return n
		}
		n = n % m
		if n == 0 {
			return m
		}
	}
}

// Implementation avoiding trivial replacement and
// with only one loop's exit condition.
func gcd_fs(m, n int) int {
	for m != 0 && n != 0 {
		m = m % n
		if m != 0 {
			n = n % m
		}
	}
	if m != 0 {
		return m
	} else {
		return n
	}
}
