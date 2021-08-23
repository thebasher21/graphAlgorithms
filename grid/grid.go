package grid

import "math"

type Pair struct {
	X int
	Y int
}

func exploreIsland(grid [][]string, cell Pair, vis map[Pair]bool) {
	if cell.X < len(grid) && cell.Y < len(grid[0]) && cell.X >= 0 && cell.Y >= 0 {
		if grid[cell.X][cell.Y] == "L" {
			if _, ok := vis[cell]; !ok {
				vis[cell] = true
				up, down, left, right := Pair{cell.X - 1, cell.Y}, Pair{cell.X + 1, cell.Y}, Pair{cell.X, cell.Y - 1}, Pair{cell.X, cell.Y + 1}
				exploreIsland(grid, up, vis)
				exploreIsland(grid, down, vis)
				exploreIsland(grid, left, vis)
				exploreIsland(grid, right, vis)
			}
		}
	}
}

func islandSize(grid [][]string, cell Pair, vis map[Pair]bool) int {
	if cell.X < len(grid) && cell.Y < len(grid[0]) && cell.X >= 0 && cell.Y >= 0 {
		if grid[cell.X][cell.Y] == "L" {
			if _, ok := vis[cell]; !ok {
				vis[cell] = true
				up, down, left, right := Pair{cell.X - 1, cell.Y}, Pair{cell.X + 1, cell.Y}, Pair{cell.X, cell.Y - 1}, Pair{cell.X, cell.Y + 1}
				return 1 + islandSize(grid, up, vis) + islandSize(grid, down, vis) + islandSize(grid, left, vis) + islandSize(grid, right, vis)
			}
		}
	}
	return 0
}

func IslandCount(grid [][]string) int {
	vis := map[Pair]bool{}
	count := 0
	for i := range grid {
		for j := range grid[i] {
			cell := Pair{i, j}
			if _, ok := vis[cell]; grid[i][j] == "L" && !ok {
				count += 1
				exploreIsland(grid, cell, vis)
			}
		}
	}
	return count
}

func SmallestIsland(grid [][]string) int {
	vis := map[Pair]bool{}
	smallest := math.MaxInt16
	for i := range grid {
		for j := range grid[i] {
			cell := Pair{i, j}
			if _, ok := vis[cell]; grid[i][j] == "L" && !ok {
				newSize := islandSize(grid, cell, vis)
				if newSize < smallest {
					smallest = newSize
				}
			}
		}
	}
	return smallest
}
