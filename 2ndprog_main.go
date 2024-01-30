package main

import "fmt"

type NSPEvent struct {
	EventDate      string
	EventName      string
	PrimaryContact string // PrimaryContact for NSPEvent
	OrganizingTeam OrganizingTeam
}

type OrganizingTeam struct {
	TeamMembers    []string
	PrimaryContact string // PrimaryContact for OrganizingTeam
}

func main() {
	fmt.Println("EXERCISE 2")

	o := OrganizingTeam{TeamMembers: []string{"Shruti", "Giri", "Sahana"},
		PrimaryContact: ""}

	name := "Sowjanya"
	if checkPrimaryContact(o, name) {
		o.PrimaryContact = name
	} else {
		o.PrimaryContact = ""
	}

	n := NSPEvent{EventDate: "23rd Jan 2024",
		EventName:      "Sankranthi Celebrations",
		PrimaryContact: "Sowjanya",
		OrganizingTeam: o,
	}

	fmt.Println("Details are as follows-")
	fmt.Println("Event Date:", n.EventDate)
	fmt.Println("Event Name:", n.EventName)
	fmt.Println("Primary Contact (for NSPEvent):", n.PrimaryContact)
	fmt.Println("Primary Contact (for OrganizingTeam):", n.OrganizingTeam.PrimaryContact)
	fmt.Println("Organizing Team Members:", n.OrganizingTeam.TeamMembers)
}

func checkPrimaryContact(o OrganizingTeam, name string) bool {

	for _, member := range o.TeamMembers {
		if member == name {
			return true
		} else {
			return false
		}
	}
	return false
}
