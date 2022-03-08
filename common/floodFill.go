package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	if image[sr][sc] == newColor { return image }
	helper(&image, sr, sc, image[sr][sc], newColor)
	return image

}

func helper(image *[][]int, i, j, oldColor, newColor int) {
	if i < 0 || i >= len(*image) || j < 0 || j >= len((*image)[0]) || (*image)[i][j] != oldColor { return }
	(*image)[i][j] = newColor
	helper(image, i+1, j, oldColor, newColor)
	helper(image, i-1, j, oldColor, newColor)
	helper(image, i, j+1, oldColor, newColor)
	helper(image, i, j-1, oldColor, newColor)
}