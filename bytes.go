package cxbytes

import (
	"bytes"
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

func AppendSortedByteSlices(sorted *[][]byte, new []byte) {
	if len(*sorted) == 0 {
		*sorted = append(*sorted, new)
		return
	}
	if bytes.Compare((*sorted)[len(*sorted)-1], new) < 0 {
		*sorted = append(*sorted, new)
		return
	}
	if bytes.Compare((*sorted)[0], new) > 0 {
		*sorted = append([][]byte{new}, *sorted...)
		return
	}

	left := 0
	right := len(*sorted) - 1

	for left <= right {
		mid := left + (right-left)/2
		cmp := bytes.Compare((*sorted)[mid], new)
		if cmp == 0 {
			copy((*sorted)[mid+1:], (*sorted)[mid:])
			(*sorted)[mid] = new
			return
		}
		if cmp < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	*sorted = append((*sorted)[:left+1], append([][]byte{new}, (*sorted)[left+1:]...)...)
}

func SortByteSlices(s *[][]byte) {
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

    SortByteSlices(&(*s)[:left])
    SortByteSlices(&(*s)[left+1:])
}


func ListContainsBytes(list [][]byte, value []byte) uint8 {
    for _, v := range list {
        if bytes.Equal(v, value) {
            return 1
        }
    }
    return 0
}


func SortedListContainsBytes(sorted [][]byte, target []byte) uint8 {
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

func Incr(data *[]byte) {
	if len(*data) == 0 {
		*data = []byte{1}
	}

	for i := len(*data); i > 0; i-- {
		(*data)[i-1]++
		if (*data)[i-1] != 0 {
			break
		} else if len(*data) == i {
			*data = append([]byte{1}, (*data)...)
		}
	}
}

