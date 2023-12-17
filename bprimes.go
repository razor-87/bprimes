package bprimes

import "unsafe"

var (
	sb  = unsafe.Pointer(&blocks)
	o2b = [m]uint32{
		1: 1 << 0, 3: 1 << 1, 7: 1 << 2, 9: 1 << 3, 11: 1 << 4, 13: 1 << 5, 17: 1 << 6, 19: 1 << 7,
		21: 1 << 8, 23: 1 << 9, 27: 1 << 10, 29: 1 << 11, 31: 1 << 12, 33: 1 << 13, 37: 1 << 14, 39: 1 << 15,
		41: 1 << 16, 43: 1 << 17, 47: 1 << 18, 49: 1 << 19, 51: 1 << 20, 53: 1 << 21, 57: 1 << 22, 59: 1 << 23,
		61: 1 << 24, 63: 1 << 25, 67: 1 << 26, 69: 1 << 27, 71: 1 << 28, 73: 1 << 29, 77: 1 << 30, 79: 1 << 31,
	}
)

const (
	two   = 2
	five  = 5
	m     = 80
	k     = 1280
	s     = 4
	b     = 64
	limit = 15485864
)

func IsPrime(n uint32) bool {
	switch {
	case n&1 == 0:
		return n == two
	case n%five == 0:
		return n == five
	}

	return primeBit(n/k, n%k) != 0
}

func Limit() int {
	return limit
}

func primeBit(q, r uint32) uint32 {
	return *(*uint32)(unsafe.Add(unsafe.Pointer((*block)(unsafe.Add(sb, q*b))), (r/m)*s)) & o2b[r%m]
}
