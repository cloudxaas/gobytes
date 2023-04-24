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
	for i := 0; i < len(new); i += 2 {
		newKey := new[i]
		newValue := new[i+1]
		sortedLen := len(*sorted)

		if sortedLen == 0 {
			*sorted = [][]byte{newKey, newValue}
			continue
		}

		left := 0
		right := (sortedLen / 2) - 1

		for left <= right {
			mid := left + (right-left)/2
			cmp := bytes.Compare((*sorted)[2*mid], newKey)
			if cmp == 0 {
				*sorted = append((*sorted)[:2*mid+2], (*sorted)[2*mid:]...)
				(*sorted)[2*mid] = newKey
				(*sorted)[2*mid+1] = newValue
				break
			}
			if cmp < 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left <= right {
			continue
		}
		insertIndex := 2 * left
		*sorted = append(append((*sorted)[:insertIndex], newKey, newValue), (*sorted)[insertIndex:]...)
	}
}


func AppendSortedUniqueKV(sorted *[][]byte, new [][]byte) {
	for i := 0; i < len(new); i += 2 {
		newKey := new[i]
		newValue := new[i+1]
		sortedLen := len(*sorted)

		if sortedLen == 0 {
			*sorted = [][]byte{newKey, newValue}
			continue
		}

		left := 0
		right := (sortedLen / 2) - 1

		for left <= right {
			mid := left + (right-left)/2
			cmp := bytes.Compare((*sorted)[2*mid], newKey)
			if cmp == 0 {
				(*sorted)[2*mid+1] = newValue
				break
			}
			if cmp < 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left <= right {
			continue
		}
		insertIndex := 2 * left
		*sorted = append(append((*sorted)[:insertIndex], newKey, newValue), (*sorted)[insertIndex:]...)
	}
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

