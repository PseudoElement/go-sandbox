package funcs

import (
	"slices"
)

func MaxArea(height []int) int {
	if len(height) < 3 {
		return slices.Min(height)
	}

	startPtr := 0
	endPtr := len(height) - 1
	res := (endPtr - startPtr) * min(height[startPtr], height[endPtr])

	for startPtr < endPtr {
		newRes := (endPtr - startPtr) * min(height[startPtr], height[endPtr])

		if newRes > res {
			res = newRes
		}

		if height[endPtr] > height[startPtr] {
			startPtr++
		} else {
			endPtr--
		}
	}

	return res
}
