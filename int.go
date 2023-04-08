package cxbytes

func Int64ToBytes(n int) [8]byte {
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

func Int32ToBytes(n int32) [4]byte {
	uint32Value := uint32(n)
	return [4]byte{
		byte(uint32Value >> 24),
		byte(uint32Value >> 16),
		byte(uint32Value >> 8),
		byte(uint32Value),
	}
}

func Int16ToBytes(n int16) [2]byte {
	uint16Value := uint16(n)
	return [2]byte{
		byte(uint16Value >> 8),
		byte(uint16Value),
	}
}

func Int8ToBytes(n int8) [1]byte {
	uint8Value := uint8(n)
	return [1]byte{
		byte(uint8Value),
	}
}
