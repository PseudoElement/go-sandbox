package utils

type SliceInfo struct {
	slice    []int
	capacity int
	length   int
}

func CustomAppend(sl []int, nums ...int) SliceInfo {
	var res []int
	totalResLen := len(sl) + len(nums)
	// not create new array in memory, used slice of old base array with high capacity
	if totalResLen <= cap(sl) {
		res = sl[:totalResLen]
		for idx, num := range nums {
			res[len(sl)+idx] = num
		}

		return SliceInfo{
			slice:    res,
			capacity: cap(res),
			length:   len(res),
		}
	}

	capacity := (len(sl) + len(nums)) * 2
	newSlice := make([]int, len(sl)+len(nums), capacity)
	copy(newSlice, sl)
	for idx, num := range nums {
		newSlice[len(sl)+idx] = num
	}

	return SliceInfo{
		slice:    newSlice,
		capacity: capacity,
		length:   len(newSlice),
	}
}
