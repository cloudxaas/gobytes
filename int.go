package cxbytes

//uint8 to byte
func Uint8ToBytes(n uint8) [1]byte {
	return [1]byte{n}
}

//uint16 to fixed 2 bytes, big endian
func Uint16ToBytes(n uint16) [2]byte {
	return [2]byte{
		byte(n >> 8),
		byte(n),
	}
}

//uint16 to fixed 2 bytes, little endian
func Uint16ToBytesR(n uint16) [2]byte {
	return [2]byte{
		byte(n),
		byte(n >> 8),
	}
}

//this works for variable length bytes
func BytesToUint16(b []byte) uint16 {
    var n uint16
    for i := range b {
        n |= uint16(b[i]) << uint(8*i)
    }
    return n
}



//uint24 to fixed 3 bytes
func Uint24ToBytes(n uint32) [3]byte {
	return [3]byte{
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}
}

//uint24 to fixed 3 bytes, little endian
func Uint24ToBytesR(n uint32) [3]byte {
	return [3]byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
	}
}


//bigendian
func VBytesToUint16(b []byte) uint16 {
    var j uint16
    for _, v := range b {
        j = (j << 8) | uint16(v)
    }
    return j
}


//bigendian
func VBytesToUint32(b []byte) uint32 {
    var j uint32
    for _, v := range b {
        j = (j << 8) | uint32(v)
    }
    return j
}



//bigendian
func VBytesToUint64(b []byte) uint64 {
    var j uint64
    for _, v := range b {
        j = (j << 8) | uint64(v)
    }
    return j
}


func BytesToUint64(array []byte) uint64 {
	var data uint64 = 0
	if len(array) == 0 {
		return 0
	}
	//for i:=0;i< len(array);i++  {
	for i := 0; i < len(array); i++ {
		data += uint64(uint(array[i]) << uint(8*(len(array)-i-1)))
	}
	return data
}


func Uint64ToBytes(n uint64) [8]byte {
	return [8]byte{
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

//bigendian
func BytesToInt56(element []byte) int64 {
    var val int64
    for i := 0; i < 7; i++ {
        if i < len(element) {
            val += int64(element[len(element)-1-i]) << uint(8*(7-i))
        }
    }
    if element[0]&0x80 != 0 {
        val -= 1 << 56
    }
    return val
}




func BytesToUint24(b []byte) uint32 {
	return VBytesToUint32(b)
}






