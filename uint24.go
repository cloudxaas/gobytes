package cxbytes

import (
	"errors"
)

// ToBytes converts a uint32 to a 3-byte slice (representing a 24-bit unsigned integer) using big-endian byte order.
func Uint24ToBytes(value uint32) ([]byte, error) {
	if value > 0xFFFFFF {
		return nil, errors.New("value exceeds the maximum for uint24")
	}

	bytes := make([]byte, 3)
	bytes[0] = byte((value >> 16) & 0xFF)
	bytes[1] = byte((value >> 8) & 0xFF)
	bytes[2] = byte(value & 0xFF)
	return bytes, nil
}

// FromBytes converts a 3-byte slice (representing a 24-bit unsigned integer) to a uint32 using big-endian byte order.
func BytesToUint24(bytes []byte) (uint32, error) {
	if len(bytes) != 3 {
		return 0, errors.New("input byte slice must have a length of 3")
	}

	value := uint32(0)
	value |= uint32(bytes[0]) << 16
	value |= uint32(bytes[1]) << 8
	value |= uint32(bytes[2])
	return value, nil
}
