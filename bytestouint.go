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
