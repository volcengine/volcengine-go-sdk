package sdkrand

import "math/rand"

func Read(r *rand.Rand, p []byte) (int, error) {
	return r.Read(p)
}
