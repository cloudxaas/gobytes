package cxbytes

func IntToBytes(n int) [8]byte {
	uint64Value := uint64(n)
	return [8]byte{
		byte(uint64Value >> 56),
		byte(uint64Value >> 48),
		byte(uint64Value >> 40),
		byte(uint64Value >> 32),
		byte(uint64Value >> 24),
		byte(uint64Value >> 16),
		byte(uint64Value >> 8),
		byte(uint64Value),
	}
}
