package rk

const prime = 16777619

func Hash(s string, x uint32) (h uint32) {
	for i := 0; i < len(s); i++ {
		h = h*x + uint32(s[i])
	}
	return h
}

func Search(s, sep string) int {
	n := len(sep)
	if n > len(s) {
		return -1
	}

	hsep := hash(sep, prime)
	p := pow(len(sep)-1, prime)
	h := hash(s[:n], prime)

	if hsep == h {
		return 0
	}
	for i := n; i < len(s); i++ {
		h -= p * uint32(s[i-n])
		h *= prime
		h += uint32(s[i])
		if hsep == h {
			return i
		}
	}
	return -1
}

func pow(n int, x uint32) (pow uint32) {
	pow = 1
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= x
		}
		x *= x
	}
	return
}
