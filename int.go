package cxbytes

func Int8ToBytesR(n int8) []byte {
	uint8Value := uint8(n)
	result := make([]byte, 1)
	result[0] = byte(uint8Value)
	return result
}

func Int16ToBytesR(n int16) []byte {
	uint16Value := uint16(n)
	result := make([]byte, 2)
	result[0] = byte(uint16Value)
	result[1] = byte(uint16Value >> 8)
	return result
}

func Int32ToBytesR(n int32) []byte {
	uint32Value := uint32(n)
	result := make([]byte, 4)
	result[0] = byte(uint32Value)
	result[1] = byte(uint32Value >> 8)
	result[2] = byte(uint32Value >> 16)
	result[3] = byte(uint32Value >> 24)
	return result
}

func Int64ToBytesR(n int64) []byte {
	uint64Value := uint64(n)
	result := make([]byte, 8)
	result[0] = byte(uint64Value)
	result[1] = byte(uint64Value >> 8)
	result[2] = byte(uint64Value >> 16)
	result[3] = byte(uint64Value >> 24)
	result[4] = byte(uint64Value >> 32)
	result[5] = byte(uint64Value >> 40)
	result[6] = byte(uint64Value >> 48)
	result[7] = byte(uint64Value >> 56)
	return result
}

func Int64ToBytes(n int) []byte {
	uint64Value := uint64(n)
	result := make([]byte, 8)
	result[0] = byte(uint64Value >> 56)
	result[1] = byte(uint64Value >> 48)
	result[2] = byte(uint64Value >> 40)
	result[3] = byte(uint64Value >> 32)
	result[4] = byte(uint64Value >> 24)
	result[5] = byte(uint64Value >> 16)
	result[6] = byte(uint64Value >> 8)
	result[7] = byte(uint64Value)
	return result
}

func Int32ToBytes(n int32) []byte {
	uint32Value := uint32(n)
	result := make([]byte, 4)
	result[0] = byte(uint32Value >> 24)
	result[1] = byte(uint32Value >> 16)
	result[2] = byte(uint32Value >> 8)
	result[3] = byte(uint32Value)
	return result
}

func Int16ToBytes(n int16) []byte {
	uint16Value := uint16(n)
	result := make([]byte, 2)
	result[0] = byte(uint16Value >> 8)
	result[1] = byte(uint16Value)
	return result
}

func Int8ToBytes(n int8) []byte {
	uint8Value := uint8(n)
	result := make([]byte, 1)
	result[0] = byte(uint8Value)
	return result
}

