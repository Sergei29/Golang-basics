package main

import . "fmt"

type Profile struct {
	id      string
	name    string
	credits int
}

var data []Profile = []Profile{
	{id: "a1", name: "johnDoe", credits: 12}, {id: "d2", name: "annieLennox", credits: 100}, {id: "z8", name: "jennaJamesson", credits: 2020},
}

func structsAndMaps() {
	divider()

	list := getProfileList(data)
	Printf("%+v\n", list)
	divider()

	for id, profile := range list {
		profile.updateCredits(200)
		// important! in map: each list[index] item must be mutated as well, as below:
		list[id] = profile
	}

	Printf("%+v\n", list)
}

func getProfileList(data []Profile) map[string]Profile {
	profileList := make(map[string]Profile)

	for _, p := range data {
		profileList[p.id] = p
	}

	return profileList
}

func (p *Profile) updateCredits(newCredit int) {
	(*p).credits = newCredit

}

func divider() {
	Println("**********************************")
}
