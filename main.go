package main

import . "fmt"

type Profile struct {
	id   int
	name string
}

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00b300",
	// 	"blue":  "#0040ff",
	// }

	// printMap(colors)

	profiles := []Profile{{id: 1, name: "John"}, {id: 2, name: "Mary"}, {id: 3, name: "Louise"}}
	profilesMap := make(map[Profile]int)

	for index, profile := range profiles {
		profilesMap[profile] = (index + 5) * 1000 * profile.id
	}

	Printf("%+v\n", profilesMap)

}

func printMap(m map[string]string) {
	for key, val := range m {
		Println("Hex code for", key, "is", val)
	}
}
