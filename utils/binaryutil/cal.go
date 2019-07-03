package binaryutil

func AppendBytes(origin []byte, data []byte) []byte {
	m := len(origin)
	n := m + len(data)
	if n > cap(origin) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, origin)
		origin = newSlice
	}
	origin = origin[0:n]
	copy(origin[m:n], data)
	return origin
}
