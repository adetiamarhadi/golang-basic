package main

import "fmt"

func main() {
	// var color map[string]string

	color := make(map[string]string)

	color["red"] = "merah"

	// color := map[string]string{
	// 	"red":   "#ff0000",
	// 	"white": "#ffffff",
	// }

	delete(color, "red")

	fmt.Println(color)
}
