package cxbytes

func BytesToInt8(b []byte) int8 {
	return int8(b[0])
}

func BytesToInt16R(b []byte) int16 {
	return int16(uint16(b[0]) | (uint16(b[1]) << 8))
}

func BytesToInt32R(b []byte) int32 {
	return int32(uint32(b[0]) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24))
}

func BytesToInt64R(b []byte) int64 {
	return int64(uint64(b[0]) | (uint64(b[1]) << 8) | (uint64(b[2]) << 16) | (uint64(b[3]) << 24) | (uint64(b[4]) << 32) | (uint64(b[5]) << 40) | (uint64(b[6]) << 48) | (uint64(b[7]) << 56))
}

func BytesToInt64(b []byte) int64 {
	return int64((uint64(b[0]) << 56) | (uint64(b[1]) << 48) | (uint64(b[2]) << 40) | (uint64(b[3]) << 32) | (uint64(b[4]) << 24) | (uint64(b[5]) << 16) | (uint64(b[6]) << 8) | uint64(b[7]))
}

func BytesToInt32(b []byte) int32 {
	return int32((uint32(b[0]) << 24) | (uint32(b[1]) << 16) | (uint32(b[2]) << 8) | uint32(b[3]))
}

func BytesToInt16(b []byte) int16 {
	return int16((uint16(b[0]) << 8) | uint16(b[1]))
}
