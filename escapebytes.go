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

func AppendKV(list *[][]byte, new [][]byte) {
	for i := 0; i < len(new); i += 2 {
		newKey := new[i]
		newValue := new[i+1]
		*list = append(*list, newKey, newValue)
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
                insertIndex := -1

                for left <= right {
                        mid := left + (right-left)/2
                        cmp := bytes.Compare((*sorted)[2*mid], newKey)
                        if cmp == 0 {
                                (*sorted)[2*mid] = newKey
                                (*sorted)[2*mid+1] = newValue
                                insertIndex = -1
                                break
                        }
                        if cmp < 0 {
                                left = mid + 1
                        } else {
                                right = mid - 1
                        }
                }

                if insertIndex == -1 {
                        insertIndex = 2 * left
                }

                *sorted = append((*sorted)[:insertIndex], append([][]byte{newKey, newValue}, (*sorted)[insertIndex:]...)...)
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
		found := false

		for left <= right {
			mid := left + (right-left)/2
			cmp := bytes.Compare((*sorted)[2*mid], newKey)
			if cmp == 0 {
				(*sorted)[2*mid+1] = newValue
				found = true
				break
			}
			if cmp < 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if found {
			continue
		}
		insertIndex := 2 * left
		*sorted = append((*sorted)[:insertIndex], append([][]byte{newKey, newValue}, (*sorted)[insertIndex:]...)...)
	}
}

func AppendUniqueKV(sorted *[][]byte, new [][]byte) {
	for i := 0; i < len(new); i += 2 {
		newKey := new[i]
		newValue := new[i+1]

		existingIndex := -1
		for j := 0; j < len(*sorted); j += 2 {
			if bytes.Equal((*sorted)[j], newKey) {
				existingIndex = j
				break
			}
		}

		if existingIndex != -1 {
			(*sorted)[existingIndex+1] = newValue
		} else {
			*sorted = append(*sorted, newKey, newValue)
		}
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



func RemoveAllFromSortedKVList(sorted *[][]byte, key []byte, caseSensitive uint8) {
	for {
		_, found := SortedKVListContains(sorted, key)
		if !found {
			break
		}

		RemoveFromSortedKVList(sorted, key, caseSensitive)
	}
}

func RemoveFromSortedKVList(sorted *[][]byte, key []byte, caseSensitive uint8) {
	left := 0
	right := (len(*sorted) / 2) - 1
	var keyLower []byte
	if caseSensitive == 1 {
		keyLower = bytes.ToLower(key)
	}else{
		keyLower = key
	}

	for left <= right {
		mid := left + (right-left)/2
		midKey := (*sorted)[2*mid]
		if caseSensitive == 1 {
			midKey = bytes.ToLower(midKey)
		}
		cmp := bytes.Compare(midKey, keyLower)
		if cmp == 0 {
			copy((*sorted)[2*mid:], (*sorted)[2*(mid+1):])
			(*sorted)[len(*sorted)-2] = nil
			(*sorted)[len(*sorted)-1] = nil
			*sorted = (*sorted)[:len(*sorted)-2]
			left = 0
			right = (len(*sorted) / 2) - 1
			continue
		}
		if cmp < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
