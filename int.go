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
		return []byte{
			byte(n >> 8),
			byte(n),
		}

	} else {
		return []byte{byte(n)}
	}
}



//this works for variable length bytes
func BytesToUint16(b []byte) uint16 {
    if len(b) == 0 {
        return 0
    }
	if len(b) == 1 {
		return uint16(b[0])
	}
    return uint16(b[1]) | uint16(b[0])<<8
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
	for i := 0; i < len(b); i++ {
		j += (uint16(b[i]) << ((len(b) - i - 1) * 8))
	}
	return j
}










func Int32ToBytes(n int32) []byte {
	return Uint32ToBytes(uint32(n))
}

func Uint32ToBytes(n uint32) []byte {
	return []byte{
		byte(n >> 24),
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}
}


func BytesToUint32(array []byte) uint32 {
	var data uint32 = 0
	if len(array) == 0 {
		return 0
	}
	//for i:=0;i< len(array);i++  {
	for i := 0; i < len(array); i++ {
		data += uint32(uint(array[i]) << uint(8*(len(array)-i-1)))
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
	for i := 0; i < len(b); i++ {
		j = j + (uint32(b[i]) << ((len(b) - i - 1) * 8))
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
	for i := 0; i < len(b); i++ {
		j = j + (uint64(b[i]) << ((len(b) - i - 1) * 8))
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


//this is big endian
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
