package cxbytes

func Uint8ToVBytes(n uint8) []byte {
    return []byte{n}
}

func Uint16ToVBytes(n uint16) []byte {
    if n > 255 {
        return []byte{byte(n >> 8), byte(n)}
    }
    return []byte{byte(n)}
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
	}
	return []byte{byte(n)}
}

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
