package cxbytesslice

// If the length of the source slice is greater than the target slice, it creates a new allocation.
func ReplaceByteSlice(target *[]byte, source []byte) {
    if len(*target) < len(source) {
        // Create a new slice with the length of source
        newSlice := make([]byte, len(source))
        copy(newSlice, source)
        *target = newSlice
    } else if len(*target) > len(source) {
        // Resize target to match the length of source
        *target = (*target)[:len(source)]
        copy(*target, source)
    } else {
        // If lengths are equal, just copy source into target
        copy(*target, source)
    }
}
