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
	sortedLen := len(*sorted)
	for i := 0; i < len(new); i += 2 {
		newKey := new[i]
		newValue := new[i+1]

		existingIndex := -1
		for j := 0; j < sortedLen; j += 2 {
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
			if sortedLen < cap(*sorted) - 1 {
				(*sorted)[sortedLen] = newKey
				(*sorted)[sortedLen+1] = newValue
				sortedLen += 2
			}
		}
	}
}


func AppendSorted(sorted *[][]byte, new []byte) {
	sortedLen := len(*sorted)

	if sortedLen == 0 || bytes.Compare((*sorted)[sortedLen-1], new) < 0 {
		(*sorted)[sortedLen] = new
		return
	}

	if bytes.Compare((*sorted)[0], new) > 0 {
		copy((*sorted)[1:], (*sorted)[0:])
		(*sorted)[0] = new
		return
	}

	left := 0
	right := sortedLen - 1

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

	copy((*sorted)[left+2:], (*sorted)[left+1:])
	(*sorted)[left+1] = new
}

func AppendSortedUnique(sorted *[][]byte, new []byte) {
	sortedLen := len(*sorted)

	if sortedLen == 0 || bytes.Compare((*sorted)[sortedLen-1], new) < 0 {
		(*sorted)[sortedLen] = new
		return
	}

	if bytes.Compare((*sorted)[0], new) > 0 {
		copy((*sorted)[1:], (*sorted)[0:])
		(*sorted)[0] = new
		return
	}

	left := 0
	right := sortedLen - 1

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

	if left < sortedLen && bytes.Equal((*sorted)[left], new) {
		return
	}

	copy((*sorted)[left+2:], (*sorted)[left+1:])
	(*sorted)[left+1] = new
}

func RemoveAllFromKVList(list *[][]byte, key []byte, caseSensitive uint8) {
	i := 0
	for i < len(*list) {
		var equal bool
		if caseSensitive == 1 {
			equal = bytes.Equal((*list)[i], key)
		} else {
			equal = bytes.EqualFold((*list)[i], key)
		}

		if equal {
			copy((*list)[i:], (*list)[i+2:])
			(*list)[len(*list)-2] = nil
			(*list)[len(*list)-1] = nil
			*list = (*list)[:len(*list)-2]
			continue
		}
		i += 2
	}
}

func RemoveAllFromSortedKVList(sorted *[][]byte, key []byte, caseSensitive uint8) {
	_, found := SortedKVListContains(sorted, key)
	for found {
		RemoveFromSortedKVList(sorted, key, caseSensitive)
		_, found = SortedKVListContains(sorted, key)
	}
}

func RemoveFromSortedKVList(sorted *[][]byte, key []byte, caseSensitive uint8) {
	left := 0
	right := (len(*sorted) / 2) - 1
	var keyLower []byte
	if caseSensitive == 1 {
		keyLower = bytes.ToLower(key)
	} else {
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
			return
		}
		if cmp < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
