package cxbytes

func BytesReverse(element []byte) []byte {
    for i, j := 0, len(element)-1; i < j; i, j = i+1, j-1 {
        element[i], element[j] = element[j], element[i]
    }
    return element
}


