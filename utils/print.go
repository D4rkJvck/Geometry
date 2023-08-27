package utils

import (
	"fmt"
	"geometry/models"
)

func PrintShapeStats(fig models.Shape) {
	fmt.Println("Area: ", fig.Area())
	fmt.Println("Perimeter: ", fig.Perimeter())
}

func PrintSolidStats(fig models.Solid) {
	fmt.Println("Area: ", fig.Area())
	fmt.Println("Volume: ", fig.Volume())
}
