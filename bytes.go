package cxbytes

func BytesReverse(element []byte) []byte {
	for i := len(element)/2-1; i >= 0; i-- {
		opp := len(element)-1-i
		element[i], element[opp] = element[opp], element[i]
	}
	return element
}

/* replaced with PutUint64 encoding/binary
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
*/

func BytesToUint64(array []byte) uint64 {
	var data uint64 = 0
	if len(array) == 0 {
		return 0
	}
	for i := 0; i < len(array); i++ {
		data += uint64(uint(array[i]) << uint(8*(len(array)-i-1)))
	}
	return data
}


func BytesToUint16(array []byte) uint16 {
	//return binary.BigEndian.Uint16(array)
	if len(array) == 0 {
		return 0
	}
	var data uint16 = 0
	for i := 0; i < len(array); i++ {
		data += uint16(uint(array[i]) << uint(8*(len(array)-i-1)))
	}
	return data

}

func BytesToUint32(array []byte) uint32 {
	var data uint32 = 0
	if len(array) == 0 {
		return 0
	}
	for i := 0; i < len(array); i++ {
		data += uint32(uint(array[i]) << uint(8*(len(array)-i-1)))
	}
	return data
}
