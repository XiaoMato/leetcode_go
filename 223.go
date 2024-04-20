package main

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {

	area := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)

	top := min(ay2, by2)
	bottom := maxInt(ay1, by1)
	right := min(ax2, bx2)
	left := maxInt(ax1, bx1)

	if top <= bottom || right <= left {
		return area
	}

	return area - (top-bottom)*(right-left)
}
