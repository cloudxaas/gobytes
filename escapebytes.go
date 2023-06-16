package cxbytes

import "bytes"

func Incr(data []byte) {
	carry := true
	for i := len(data) - 1; i >= 0 && carry; i-- {
		data[i]++
		carry = data[i] == 0
	}
	if carry {
		data = append(data, 0)
		copy(data[1:], data)
		data[0] = 1
	}
}

func AppendKV(list *[][]byte, new [][]byte) {
	for i := 0; i < len(new); i += 2 {
		*list = append(*list, new[i], new[i+1])
	}
}

func AppendSortedKV(sorted *[][]byte, new [][]byte) {
        sortedLen := len(*sorted)
        for i := 0; i < len(new); i += 2 {
                newKey := new[i]
                newValue := new[i+1]

                left := 0
                right := sortedLen/2 - 1
                insertIndex := sortedLen // set insertIndex to sortedLen as default

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

                if insertIndex != -1 {
                        insertIndex = 2 * left
                        copy((*sorted)[insertIndex+2:], (*sorted)[insertIndex:])
                        (*sorted)[insertIndex] = newKey
                        (*sorted)[insertIndex+1] = newValue
                        sortedLen += 2
                }
        }
}



func AppendSortedUniqueKV(sorted *[][]byte, new [][]byte) {
	sortedLen := len(*sorted)
	for i := 0; i < len(new); i += 2 {
		newKey := new[i]
		newValue := new[i+1]

		left := 0
		right := sortedLen/2 - 1
		insertIndex := sortedLen // set insertIndex to sortedLen as default

		for left <= right {
			mid := left + (right-left)/2
			cmp := bytes.Compare((*sorted)[2*mid], newKey)
			if cmp == 0 {
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

		if insertIndex != -1 {
			insertIndex = 2 * left
			copy((*sorted)[insertIndex+2:], (*sorted)[insertIndex:])
			(*sorted)[insertIndex] = newKey
			(*sorted)[insertIndex+1] = newValue
			sortedLen += 2
		}
	}
}


func AppendUniqueKV(sorted *[][]byte, new [][]byte, overwrite uint8) {
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
			if overwrite == 1 {
				(*sorted)[existingIndex+1] = newValue
			}
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

func RemoveAllFromKVList(list *[][]byte, key []byte, caseSensitive uint8) {
	for i := 0; i < len(*list); i += 2 {
		var equal bool
		if caseSensitive == 1 {
			equal = bytes.Equal((*list)[i], key)
		} else {
			equal = bytes.EqualFold((*list)[i], key)
		}

		if equal {
			*list = append((*list)[:i], (*list)[i+2:]...)
			i -= 2
		}
	}
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
