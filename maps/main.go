package main

import "fmt"

func main()  {
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	//different ways to define map
	//var  colors map[string]string
	//colors := make(map[string]string)
	//colors["white"] = "#ffffff"
	//fmt.Println(colors)
	//delete(colors, "white")
	printMap(colors)
}
func printMap(c map[string]string)  {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
