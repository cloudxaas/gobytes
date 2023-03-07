package cxbytes

func BytesReverse(element []byte) []byte {
	for i := len(element)/2-1; i >= 0; i-- {
		opp := len(element)-1-i
		element[i], element[opp] = element[opp], element[i]
	}
	return element
}
