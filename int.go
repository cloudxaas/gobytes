package cxbytes

//uint8 to byte
func Uint8ToBytes(n uint8) []byte {
	return []byte{n}
}




//uint16 to fixed 2 bytes, big endian
func Uint16ToBytes(n uint16) []byte {
	return []byte{
		byte(n >> 8),
		byte(n),
	}
}

//uint16 to fixed 2 bytes, little endian
func Uint16ToBytesR(n uint16) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
	}
}

func Uint16ToVBytes(n uint16) []byte {
    if n > 255 {
        return []byte{byte(n >> 8), byte(n)}
    }
    return []byte{byte(n)}
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
func Uint24ToBytes(n uint32) []byte {
	return []byte{
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}
}

//uint24 to fixed 3 bytes, little endian
func Uint24ToBytesR(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
	}
}



func BytesToUint24(b []byte) uint32 {
	return VBytesToUint32(b)
}


//bigendian
func VBytesToUint16(b []byte) uint16 {
    var j uint16
    for _, v := range b {
        j = (j << 8) | uint16(v)
    }
    return j
}









func Int32ToBytes(n int32) []byte {
	return Uint32ToBytes(uint32(n))
}

func Uint32ToBytes(n uint32) [4]byte {
	return []byte{
		byte(n >> 24),
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}
}


func BytesToUint32(b []byte) uint32 {
    var data uint32
    for _, v := range b {
        data = (data << 8) | uint32(v)
    }
    return data
}


func BytesToInt32(b []byte) int32 {
	var j int32
	for i := 0; i < len(b); i++ {
		if i == 0 {
			j = (int32(b[i]-0x80) << 24)
		} else {
			j = j + (int32(b[i]) << (24 - (i * 8)))
		}
	}
	return j
}



func Uint32ToVBytes(n uint32) []byte {
	if n > 16777215 {
		return []byte{
			byte(n >> 24),
			byte(n >> 16),
			byte(n >> 8),
			byte(n),
		}
	} else if n > 65535 {
		return []byte{
			byte(n >> 16),
			byte(n >> 8),
			byte(n),
		}
	} else if n > 255 {
		return []byte{
			byte(n >> 8),
			byte(n),
		}

	} else {
		return []byte{byte(n)}
	}
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


/*
//this is big endian
func Uint64ToBytes(n uint64) []byte {
	var b [8]byte
	var i int
	for ; i < 8 && n > 0; i++ {
		b[i] = byte(n >> ((7 - i) * 8))
	}
	return b[:i]

}
*/

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

/*
//this is big endian, this is variable byte length returned
func Uint64ToVBytes(n uint64) (b []byte) {
	for i := 7; i >= 0; i-- {
		b[i] = byte(n & 0xff)
		n >>= 8
	}
	return
}
*/

//this is big endian
func Uint64ToVBytes(n uint64) []byte {
	if n > 281474976710655 {
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
	} else if n > 4294967295 {
		return []byte{
			byte(n >> 48),
			byte(n >> 40),
			byte(n >> 32),
			byte(n >> 24),
			byte(n >> 16),
			byte(n >> 8),
			byte(n),
		}
	} else if n > 65535 {
		return []byte{
			byte(n >> 32),
			byte(n >> 24),
			byte(n >> 16),
			byte(n >> 8),
			byte(n),
		}
	} else if n > 255 {
		return []byte{
			byte(n >> 16),
			byte(n >> 8),
			byte(n),
		}

	}

	return []byte{byte(n)}
	
}

