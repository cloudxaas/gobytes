package cxbytes

import (
	"bytes"
)

func BytesReverse(element []byte) []byte {
    for i, j := 0, len(element)-1; i < j; i, j = i+1, j-1 {
        element[i], element[j] = element[j], element[i]
    }
    return element
}

func ListContainsBytes(list [][]byte, value []byte) bool {
    for _, v := range list {
        if bytes.Equal(v, value) {
            return true
        }
    }
    return false
}


func BytesIncr(data []byte) []byte {
	if len(data) == 0 {
		return []byte{1}
	}

	// Find the index of the first non-zero byte from the right
	idx := len(data) - 1
	for idx >= 0 && data[idx] == 0 {
		idx--
	}

	// If all bytes are 0, create a new byte slice with an extra leading 1
	if idx < 0 {
		b := make([]byte, len(data)+1)
		b[0] = 1
		return b
	}

	// Increment the last non-zero byte
	element := make([]byte, len(data))
	copy(element, data)
	element[idx]++

	return element
}


