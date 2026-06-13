package main

import "fmt"

func main() {
	const (
		LandArea = 13 * 16.0
		TileArea = 2 * 3.0
	)

	numberOfNeededTiles := LandArea / TileArea

	fmt.Print(numberOfNeededTiles)
}
