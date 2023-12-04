package bprimes

import "unsafe"

var (
	sb  = unsafe.Pointer(&blocks)
	o2b = [m]uint32{
		1: 1 << 0, 3: 1 << 1, 5: 1 << 2, 7: 1 << 3, 9: 1 << 4, 11: 1 << 5, 13: 1 << 6, 15: 1 << 7,
		17: 1 << 8, 19: 1 << 9, 21: 1 << 10, 23: 1 << 11, 25: 1 << 12, 27: 1 << 13, 29: 1 << 14, 31: 1 << 15,
		33: 1 << 16, 35: 1 << 17, 37: 1 << 18, 39: 1 << 19, 41: 1 << 20, 43: 1 << 21, 45: 1 << 22, 47: 1 << 23,
		49: 1 << 24, 51: 1 << 25, 53: 1 << 26, 55: 1 << 27, 57: 1 << 28, 59: 1 << 29, 61: 1 << 30, 63: 1 << 31,
	}
)

const (
	two   = 2
	m     = 64
	k     = 1024
	s     = 4
	b     = 64
	limit = 15485864
)

func IsPrime(n uint32) bool {
	if n&1 == 0 {
		return n == two
	}
	return primeBit(n/k, n%k) != 0
}

func Limit() int {
	return limit
}

func primeBit(q, r uint32) uint32 {
	return *(*uint32)(unsafe.Add(unsafe.Pointer((*block)(unsafe.Add(sb, q*b))), (r/m)*s)) & o2b[r%m]
}
