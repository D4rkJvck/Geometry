package main

import (
	"fmt"
	"geometry/models"
)

func main() {
	var (
		shp = make([]*models.Shape, 7)
		sld = make([]*models.Solid, 6)
	)
	fmt.Println("Choose your figure type index:")
	fmt.Println("1- Shape")
	fmt.Println("2- Solid")
	var choice string
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		src.ShapeSelector(shp)
	case "2":
		src.SolidSelector(sld)
	default:
		fmt.Println("")
	}
}
