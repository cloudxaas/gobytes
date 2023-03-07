package cxbytes

func BytesReverse(element []byte) []byte {
	for i := len(element)/2-1; i >= 0; i-- {
		opp := len(element)-1-i
		element[i], element[opp] = element[opp], element[i]
	}
	return element
}

//this is big endian
func Uint64ToBytes(n uint64) []byte {
	return []byte{
		byte(n >> 56),
		byte(n >> 48),
		byte(n >> 40),
		byte(n >> 32),
		byte(n >> 24),
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}
}
