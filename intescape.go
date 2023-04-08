package cxbytes


func Uint16ToVBytes(n uint16) []byte {
    if n > 255 {
        return []byte{byte(n >> 8), byte(n)}
    }
    return []byte{byte(n)}
}



func Uint32ToBytes(n uint32) [4]byte {
	return [4]byte{
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

