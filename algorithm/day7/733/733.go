package _33

// simple

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	srcColor := image[sr][sc]
	floodFillRecursion(image, sr, sc, srcColor, newColor)
	return image
}

func floodFillRecursion(image [][]int, r int, c int, srcColor int, newColor int) {
	if r < 0 || r > len(image)-1 || c < 0 || c > len(image[0])-1 {
		return
	}

	if image[r][c] == newColor {
		return
	} else if image[r][c] == srcColor {
		image[r][c] = newColor
	} else {
		return
	}

	floodFillRecursion(image, r, c+1, srcColor, newColor)
	floodFillRecursion(image, r, c-1, srcColor, newColor)
	floodFillRecursion(image, r+1, c, srcColor, newColor)
	floodFillRecursion(image, r-1, c, srcColor, newColor)
}
