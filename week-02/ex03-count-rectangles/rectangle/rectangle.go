package rectangle

import "fmt"

type Point struct {
	X int
	Y int
}

func checkHorizontalPoint(x int, y int, rectangles [][]int, rowVisited []bool) []Point {
	firstPoint := Point{X: x, Y: y}
	var continuousPoints = []Point{}

	continuousPoints = append(continuousPoints, firstPoint)
	rowVisited[x] = true

	for x+1 < len(rectangles[0]) && rectangles[y][x+1] == 1 {
		newPoint := Point{X: x + 1, Y: y}
		continuousPoints = append(continuousPoints, newPoint)
		rowVisited[x+1] = true
		x = x + 1
	}

	return continuousPoints
}

func checkRectangle(horizontalPoints []Point, rectangles [][]int, visited [][]bool) bool {

	prevCount := 0
	for i := 0; i < len(horizontalPoints); i++ {
		col, row := horizontalPoints[i].X, horizontalPoints[i].Y
		count := 0

		for row+1 < len(rectangles) && rectangles[row+1][col] == 1 {
			count += 1
			visited[row+1][col] = true
			row = row + 1
		}

		prevCount = count

		if i != 0 && prevCount != count {
			return false
		}
	}

	return true
}

func CountRectangles(rectangles [][]int) (int, error) {
	// maxRow := len(rectangles)
	// maxColumn := len(rectangles[0])

	visited := [][]bool{}
	total := 0

	for rowIndex := range rectangles {
		row := []bool{}
		for columnIndex := range rectangles[rowIndex] {
			columnIndex += 0 // dummy line
			row = append(row, false)
		}
		visited = append(visited, row)
	}

	for rowIndex := range rectangles {
		for colIndex := range rectangles[rowIndex] {
			if rectangles[rowIndex][colIndex] == 1 { // [y][x]
				// handle check for rectangles here
				if !visited[rowIndex][colIndex] {
					continuousHorizontal := checkHorizontalPoint(colIndex, rowIndex, rectangles, visited[rowIndex]) // check duplication
					if checkRectangle(continuousHorizontal, rectangles, visited) {
						total += 1
					} else {
						return total, fmt.Errorf("adjacent rectangle somewhere around [%v, %v]", colIndex, rowIndex)
					}
				}
			}
		}
	}
	return total, nil
}
