package main

import "fmt"

func main(){
	
	/*
	colors := map[string]string{
		"red": "#FF0000",
		"green": "#00FF00",
	}

	fmt.Println(colors)
	*/

	/*
	var colors map[string]string
	fmt.Println(colors)
	*/

	/*
	colors := make(map[string]string)
	colors["red"] = "#FF0000"
	delete(colors, "red")
	fmt.Println(colors)
	*/

	colors := map[string]string{
		"red": "#FF0000",
		"green": "#00FF00",
		"white": "FFFFFF",
	}

	printMap(colors)

}

func printMap(c map[string]string){
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}