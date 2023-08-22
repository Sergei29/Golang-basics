package main

import . "fmt"

func mapsDemo() {
	colors := ColorMap{
		"red":   "#ff0000",
		"green": "#00b300",
		"blue":  "#0040ff",
	}

	colors.print()
	colors.addColor("white", "#fff")
	Printf("%+v\n", colors)
	colors.addColor("white", "#ffffff")
	Printf("%+v\n", colors)

}

func (m ColorMap) print() {
	for key, val := range m {
		Println("Hex code for", key, "is", val)
	}
}

func (m ColorMap) addColor(color string, hex string) {
	m[color] = hex
}
