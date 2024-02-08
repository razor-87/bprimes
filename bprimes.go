package bprimes

import "unsafe"

var (
	sb  = unsafe.Pointer(&blocks)
	o2b = [m]uint32{
		1: 1 << 0, 5: 1 << 1, 7: 1 << 2, 11: 1 << 3, 13: 1 << 4, 17: 1 << 5, 19: 1 << 6, 23: 1 << 7,
		25: 1 << 8, 29: 1 << 9, 31: 1 << 10, 35: 1 << 11, 37: 1 << 12, 41: 1 << 13, 43: 1 << 14, 47: 1 << 15,
		49: 1 << 16, 53: 1 << 17, 55: 1 << 18, 59: 1 << 19, 61: 1 << 20, 65: 1 << 21, 67: 1 << 22, 71: 1 << 23,
		73: 1 << 24, 77: 1 << 25, 79: 1 << 26, 83: 1 << 27, 85: 1 << 28, 89: 1 << 29, 91: 1 << 30, 95: 1 << 31,
	}
)

const (
	fp    = 1<<2 | 1<<3 | 1<<5 | 1<<7
	m     = 96
	k     = 1536
	s     = 4
	b     = 64
	limit = 15485864
)

func IsPrime(n uint32) bool {
	switch {
	case n < 8:
		return fp&(1<<n) != 0
	case n&1 == 0, n%3 == 0, n%5 == 0, n%7 == 0:
		return false
	}

	return primeBit(n/k, n%k) != 0
}

func Limit() int {
	return limit
}

func primeBit(q, r uint32) uint32 {
	return *(*uint32)(unsafe.Add(unsafe.Pointer((*block)(unsafe.Add(sb, q*b))), (r/m)*s)) & o2b[r%m]
}
