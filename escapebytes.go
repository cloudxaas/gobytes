package cxbytes

import "bytes"

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


func AppendSortedKV(sorted *[][]byte, new [][]byte) {
	newKey := new[0]
	newValue := new[1]
	if len(*sorted) == 0 {
		*sorted = append(*sorted, newKey, newValue)
		return
	}
	if bytes.Compare((*sorted)[len(*sorted)-2], newKey) < 0 {
		*sorted = append(*sorted, newKey, newValue)
		return
	}
	if bytes.Compare((*sorted)[0], newKey) > 0 {
		*sorted = append([][]byte{newKey, newValue}, *sorted...)
		return
	}

	left := 0
	right := (len(*sorted) / 2) - 1

	for left <= right {
		mid := left + (right-left)/2
		cmp := bytes.Compare((*sorted)[2*mid], newKey)
		if cmp == 0 {
			copy((*sorted)[2*(mid+1):], (*sorted)[2*mid:])
			(*sorted)[2*mid] = newKey
			(*sorted)[2*mid+1] = newValue
			return
		}
		if cmp < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	insertIndex := 2 * left
	*sorted = append((*sorted)[:insertIndex], append([][]byte{newKey, newValue}, (*sorted)[insertIndex:]...)...)
}


func AppendSorted(sorted *[][]byte, new []byte) {
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


func AppendSortedUnique(sorted *[][]byte, new []byte) {
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
            return
        }
        if cmp < 0 {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    if left < len(*sorted) && bytes.Equal((*sorted)[left], new) {
        return
    }

    *sorted = append(*sorted, nil)
    copy((*sorted)[left+1:], (*sorted)[left:])
    (*sorted)[left] = new
}

