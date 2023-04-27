package cxbytes


// uint8 to byte
func Uint8ToBytes(n uint8) []byte {
    b := make([]byte, 1)
    b[0] = n
    return b
}

// uint16 to bytes, big endian
func Uint16ToBytes(n uint16) []byte {
    b := make([]byte, 2)
    b[0] = byte(n >> 8)
    b[1] = byte(n)
    return b
}

// uint16 to bytes, little endian
func Uint16ToBytesR(n uint16) []byte {
    b := make([]byte, 2)
    b[0] = byte(n)
    b[1] = byte(n >> 8)
    return b
}


// uint32 to bytes, big endian
func Uint32ToBytes(n uint32) []byte {
    b := make([]byte, 4)
    b[0] = byte(n >> 24)
    b[1] = byte(n >> 16)
    b[2] = byte(n >> 8)
    b[3] = byte(n)
    return b
}

// uint32 to bytes, little endian
func Uint32ToBytesR(n uint32) []byte {
    b := make([]byte, 4)
    b[0] = byte(n)
    b[1] = byte(n >> 8)
    b[2] = byte(n >> 16)
    b[3] = byte(n >> 24)
    return b
}

// uint64 to bytes, big endian
func Uint64ToBytes(n uint64) []byte {
    b := make([]byte, 8)
    b[0] = byte(n >> 56)
    b[1] = byte(n >> 48)
    b[2] = byte(n >> 40)
    b[3] = byte(n >> 32)
    b[4] = byte(n >> 24)
    b[5] = byte(n >> 16)
    b[6] = byte(n >> 8)
    b[7] = byte(n)
    return b
}

// uint64 to bytes, little endian
func Uint64ToBytesR(n uint64) []byte {
    b := make([]byte, 8)
    b[0] = byte(n)
    b[1] = byte(n >> 8)
    b[2] = byte(n >> 16)
    b[3] = byte(n >> 24)
    b[4] = byte(n >> 32)
    b[5] = byte(n >> 40)
    b[6] = byte(n >> 48)
    b[7] = byte(n >> 56)
    return b
}






