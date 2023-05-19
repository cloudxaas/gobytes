package cxbytes

import (
	"bytes"
	"sort"
)



func SplitByNewline(data []byte) [][]byte {
    var lines [][]byte
    start := 0
    for i, b := range data {
        if b == '\n' {
            lines = append(lines, data[start:i])
            start = i + 1
        }
    }
    if start < len(data) {
        lines = append(lines, data[start:])
    }
    return lines
}


func Reverse(element *[]byte) {
    for i, j := 0, len(*element)-1; i < j; i, j = i+1, j-1 {
	    (*element)[i], (*element)[j] = (*element)[j], (*element)[i]
    }    
}

func KVListContains(list *[][]byte, key []byte) ([]byte, uint8) {
	if len(*list) % 2 != 0 {
    return nil, 0

	}
	for i := 0; i < len(*list); i += 2 {
	    if bytes.Equal((*list)[i], key) {
		    return (*list)[i+1], 1
        }
    }
    return nil, 0
}


func KVListContainsCI(list *[][]byte, key []byte) ([]byte, uint8) {
	if len(*list) % 2 != 0 {
    return nil, 0

	}
	for i := 0; i < len(*list); i += 2 {
	    if bytes.EqualFold((*list)[i], key) {
		    return (*list)[i+1], 1
        }
    }
    return nil, 0
}

func ReverseBytesList(sortedBytes [][]byte, appendChars []byte) [][]byte {
	reversedBytes := make([][]byte, len(sortedBytes))

	for i, bytes := range sortedBytes {
		reversed := make([]byte, len(bytes))
		for j := 0; j < len(bytes); j++ {
			reversed[j] = bytes[len(bytes)-j-1]
		}
		if appendChars != nil {
			reversed = append(reversed, appendChars...)
		}
		reversedBytes[i] = reversed
	}

	return reversedBytes
}


func RemoveFromSortCaseInsensitive(sorted *[][]byte, toRemove []byte) {
	left := 0
	right := len(*sorted) - 1
	toRemoveLower := bytes.ToLower(toRemove)

	for left <= right {
		mid := left + (right-left)/2
		midLower := bytes.ToLower((*sorted)[mid])
		cmp := bytes.Compare(midLower, toRemoveLower)
		if cmp == 0 {
			copy((*sorted)[mid:], (*sorted)[mid+1:])
			(*sorted)[len(*sorted)-1] = nil // release the last element
			*sorted = (*sorted)[:len(*sorted)-1] // shrink the slice
			return
		}
		if cmp < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func SortedListContainsPrefix(sortedBytes [][]byte, target []byte) bool {
	index := sort.Search(len(sortedBytes), func(i int) bool {
		return bytes.HasPrefix(target, sortedBytes[i])
	})

	return index < len(sortedBytes) && bytes.HasPrefix(target, sortedBytes[index])
}


func ListContainsSuffix(listBytes [][]byte, target []byte) uint8 {
	for _, b := range listBytes {
		if bytes.HasSuffix(target, b) {
			return 1
		}
	}
	return 0
}



func RemoveFromSort(sorted *[][]byte, toRemove []byte) {
    left := 0
    right := len(*sorted) - 1

    for left <= right {
        mid := left + (right-left)/2
        cmp := bytes.Compare((*sorted)[mid], toRemove)
        if cmp == 0 {
            copy((*sorted)[mid:], (*sorted)[mid+1:])
            (*sorted)[len(*sorted)-1] = nil // release the last element
            *sorted = (*sorted)[:len(*sorted)-1] // shrink the slice
            return
        }
        if cmp < 0 {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
}

func ListContains(list [][]byte, value []byte) uint8 {
    for _, v := range list {
        if bytes.Equal(v, value) {
            return 1
        }
    }
    return 0
}


func SortedKVListContains(sorted *[][]byte, key []byte) ([][]byte, bool) {
	left := 0
	right := (len(*sorted) / 2) - 1

	// Find the first occurrence of the key using binary search
	for left <= right {
		mid := left + (right-left)/2
		cmp := bytes.Compare((*sorted)[2*mid], key)

		if cmp == 0 && (mid == 0 || bytes.Compare((*sorted)[2*(mid-1)], key) != 0) {
			left = mid
			break
		}

		if cmp < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// If the key is not found, return false
	if left > right {
		return nil, false
	}

	// Collect all values for the key
	values := [][]byte{}
	for i := left; i < len(*sorted)/2 && bytes.Compare((*sorted)[2*i], key) == 0; i++ {
		values = append(values, (*sorted)[2*i+1])
	}

	return values, true
}

func SortedListContains(sorted [][]byte, target []byte) uint8 {
    n := len(sorted)
    if n == 0 {
        return 0
    }
    i, j := 0, n-1
    for i <= j {
        h := int(uint(i+j) >> 1) // avoid overflow
        cmp := bytes.Compare(sorted[h], target)
        if cmp == 0 {
            return 1
        } else if cmp < 0 {
            i = h + 1
        } else {
            j = h - 1
        }
    }
    return 0
}




func Sort(s *[][]byte) {
	if len(*s) < 2 {
		return
	}

	pivot := (*s)[len(*s)-1]
	left, right := 0, len(*s)-2

	for left <= right {
		if bytes.Compare((*s)[left], pivot) < 0 {
			left++
			continue
		}

		if bytes.Compare((*s)[right], pivot) >= 0 {
			right--
			continue
		}

		(*s)[left], (*s)[right] = (*s)[right], (*s)[left]
		left++
		right--
	}

	(*s)[len(*s)-1], (*s)[left] = (*s)[left], (*s)[len(*s)-1]

	leftS := (*s)[:left]
	rightS := (*s)[left+1:]

	Sort(&leftS)
	Sort(&rightS)
}


