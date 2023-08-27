package controllers

import (
	"fmt"
	"geometry/error"
	"geometry/models"
	"geometry/utils"
	"strconv"
)

func ShapeSelector(shp []*models.Shape) {
	fmt.Println("Choose your shape's index...")
	fmt.Println("1- Circle")
	fmt.Println("2- Square")
	fmt.Println("3- Triangle")
	fmt.Println("4- Rectangle")
	fmt.Println("5- Pentagon")
	fmt.Println("6- Hexagon")
	fmt.Println("7- Trapezoid")
	var choice string
	fmt.Scanln(&choice)
	i, err := strconv.Atoi(choice)
	error.Handle(err)
	switch i < 8 {
	case true:
		utils.PrintShapeStats(*shp[i-1])
	default:
		fmt.Println("Please choose a valid index...")
	}
}
