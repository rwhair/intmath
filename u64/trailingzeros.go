package u64

import (
	"code.google.com/p/intmath/constants/c64"
)

// TrailingZeros returns the number of trailing 0 bits in v. If v is 0, it returns 0.
func TrailingZeros(v uint64) uint64 {
	// Adapted from Sonia Keys' xmath library
	return uint64(c64.DeBruijnBits[v&-v*c64.DeBruijnMultiple>>c64.DeBruijnShift])

}
