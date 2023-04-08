package cxbytes

// bytes to uint8
func BytesToUint8(b []byte) uint8 {
    return b[0]
}

// bytes to uint16, big endian
func BytesToUint16(b []byte) uint16 {
    return uint16(b[0])<<8 | uint16(b[1])
}

// bytes to uint16, little endian
func BytesToUint16R(b []byte) uint16 {
    return uint16(b[0]) | uint16(b[1])<<8
}

// bytes to uint32, big endian
func BytesToUint32(b []byte) uint32 {
    return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

// bytes to uint32, little endian
func BytesToUint32R(b []byte) uint32 {
    return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

// bytes to uint64, big endian
func BytesToUint64(b []byte) uint64 {
    return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 |
           uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
}

// bytes to uint64, little endian
func BytesToUint64R(b []byte) uint64 {
    return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
           uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

//Variable bytes


// Big-endian
func VBytesToUint8(b []byte) uint8 {
	var j uint8
	for i := 0; i < len(b); i++ {
		j += (uint8(b[i]) << ((len(b) - i - 1) * 8))
	}
	return j
}

// Big-endian
func VBytesToUint16(b []byte) uint16 {
	var j uint16
	for i := 0; i < len(b); i++ {
		j += (uint16(b[i]) << ((len(b) - i - 1) * 8))
	}
	return j
}

// Big-endian
func VBytesToUint32(b []byte) uint32 {
	var j uint32
	for i := 0; i < len(b); i++ {
		j = j + (uint32(b[i]) << ((len(b) - i - 1) * 8))
	}
	return j
}

// Big-endian
func VBytesToUint64(b []byte) uint64 {
	var j uint64
	for i := 0; i < len(b); i++ {
		j = j + (uint64(b[i]) << ((len(b) - i - 1) * 8))
	}
	return j
}



// Little-endian
func VBytesRToUint8(b []byte) uint8 {
	var j uint8
	for i := 0; i < len(b); i++ {
		j += (uint8(b[i]) << (i * 8))
	}
	return j
}

// Little-endian
func VBytesRToUint16(b []byte) uint16 {
	var j uint16
	for i := 0; i < len(b); i++ {
		j += (uint16(b[i]) << (i * 8))
	}
	return j
}

// Little-endian
func VBytesRToUint32(b []byte) uint32 {
	var j uint32
	for i := 0; i < len(b); i++ {
		j = j + (uint32(b[i]) << (i * 8))
	}
	return j
}

// Little-endian
func VBytesRToUint64(b []byte) uint64 {
	var j uint64
	for i := 0; i < len(b); i++ {
		j = j + (uint64(b[i]) << (i * 8))
	}
	return j
}
